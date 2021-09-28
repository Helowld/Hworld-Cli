package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/alecthomas/chroma/quick"
	"github.com/shurcooL/githubv4"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hw",
	Short: "just mention extension of language, you'll get rest",
	Long: `
	 `,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if present := resolveArgs("." + args[0]); present {
			getFileContents(args[0])
		} else {
			fmt.Println("Not a valid extension")
		}
	},
}

func getFileContents(extension string) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)
	var query struct {
		Repository struct {
			DefaultBranchRef struct {
				Target struct {
					Commit struct {
						History struct {
							Nodes []struct {
								File struct {
									Object struct {
										Blob struct {
											Text githubv4.String
										} `graphql:"... on Blob"`
									}
								} `graphql:"file(path: $name)"`
							}
						} `graphql:"history(first: 1)"`
					} `graphql:"... on Commit"`
				}
			}
		} `graphql:"repository(owner: \"Helowld\", name: \"Hello-World\")"`
	}

	filename := "hello_world." + extension
	variables := map[string]interface{}{
		"name": githubv4.String(filename),
	}
	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		if err.Error() == fmt.Sprintf("Could not resolve file for path '%s'.", filename) {
			fmt.Println("Looks like this language is not yet implemented")
		}
	}
	err = quick.Highlight(os.Stdout, string(query.Repository.DefaultBranchRef.Target.Commit.History.Nodes[0].File.Object.Blob.Text), extension, "terminal16m", "vs")
	if err != nil {
		fmt.Println(err)
		return
	}

}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
