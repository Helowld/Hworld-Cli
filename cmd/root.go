package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hw",
	Short: "hw will give you all hello world programs in all possible language",
	Long: `
	 `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome ! to world of Hello World")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
