/*
Copyright © 2023 HAIDER ALI <haider.lee23@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Haider8/aicommit/ai"
	"github.com/Haider8/aicommit/git_util"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aicommit",
	Short: "Let AI write your commit messages!",
	Long: `aicommit analyzes your code changes and generates concise and descriptive commit messages based on best practices.
Just stage the desired changes using git add and let aicommit handle the rest.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		git_diff := git_util.Diff()
		if git_diff == "" {
			fmt.Println("No changes found! Use git add to stage changes.")
			os.Exit(0)
		}
		diff_prompt := "Write a git commit message not more than 60 words for the diff: " + git_diff

		client := ai.Connect()

		is_completed := false
		var input string
		for !is_completed {
			res := ai.PromptResult(client, diff_prompt)
			fmt.Print("Commit '" + res + "'? (y/n/r [retry]) ")
			fmt.Scanln(&input)

			if input == "r" {
				is_completed = false
			} else if input == "y" {
				git_util.Commit(res)
				is_completed = true
			} else if input == "n" {
				is_completed = true
			} else {
				fmt.Println("Invalid selection " + input)
				os.Exit(1)
			}
		}

		os.Exit(0)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.aicommit.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
