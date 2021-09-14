package cmd

import (
	"context"
	"fmt"
	"os"

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
		// fmt.Println("Welcome! to Hello World")
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
		} else {
			fmt.Println("not caught")
		}
		return
	}
	fmt.Println(query.Repository.DefaultBranchRef.Target.Commit.History.Nodes[0].File.Object.Blob.Text)
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

////////////////////////////////// test query

// query {
//   repository(name: "Hello-World", owner: "Helowld") {
//     defaultBranchRef {
//       target {
//         ... on Commit {
//           history(first: 1) {
//             nodes {
//               file(path: "README.md") {
//                 object {
//                   ... on Blob {
//                     text
//                   }
//                 }
//               }
//             }
//           }
//         }
//       }
//     }
//   }
// }

// func rootTask(extension string) {
// 	filename := "hello_world." + extension
// 	query := fmt.Sprintf(`query { repository(name: "Hello-World", owner: "Helowld") { defaultBranchRef { target { ... on Commit { history(first: 1) { nodes { file(path: "%s") { object { ... on Blob { text } } } } } } } } } }
// `, filename)
// 	fmt.Printf("query: %v\n", query)
// 	jsonData := map[string]string{
// 		"query": query,
// 	}
// 	jsonValue, _ := json.Marshal(jsonData)
// 	fmt.Printf("jsonValue: %v\n", jsonValue)
// 	request, err := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(jsonValue))
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	request.Header = http.Header{
// 		"Authorization": {"bearer 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"},
// 	}
// 	client := &http.Client{Timeout: time.Second * 10}
// 	response, err := client.Do(request)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	defer response.Body.Close()
// 	if err != nil {
// 		fmt.Printf("The HTTP request failed with error %s\n", err)
// 	}
// 	data, _ := ioutil.ReadAll(response.Body)
// 	fmt.Println(string(data))

// }
