/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate stab",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		// name := args[0]
		// contents, err := os.ReadFile(name)
		// if err != nil {
		// 	fmt.Println("File reading error", err)
		// 	return
		// }

		ls, err := exec.Command("ls").Output()
		if err != nil {
			fmt.Errorf("%w", err)
		}
		fmt.Println(string(ls))
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
