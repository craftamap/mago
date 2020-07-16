package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "mago",
	Short: "mago is your git-based finance manager",
}

func init() {
	rootCmd.PersistentFlags().StringP("target", "t", ".", "Target directory on which to operate")
}

func Execute() error {
	return rootCmd.Execute()
}
