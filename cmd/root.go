package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "algo",
	Short: "A program to spin up a quick Go project using a predefined structure",
	Long: `Algo is a CLI tool that allows users to spin up a Go project with the corresponding structure seamlessly. 
It also gives the option to integrate with one of the more popular Go frameworks!`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}