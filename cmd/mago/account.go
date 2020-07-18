package main

import (
	"fmt"
	"os"
	"path/filepath"

	money "github.com/craftamap/mago"
	"github.com/prometheus/common/log"
	"github.com/spf13/cobra"
)

func init() {
	accountCmd.AddCommand(accountListCmd)
	accountCmd.AddCommand(accountCreateCmd)
	rootCmd.AddCommand(accountCmd)
}

var accountCmd = &cobra.Command{
	Use: "account",
}

var accountListCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		rootDir, err := cmd.Flags().GetString("target")
		if err != nil {
			log.Error(err)
		}
		manager, err := money.FromRoot(rootDir)
		if err != nil {
			log.Error(err)
		}

		for _, account := range manager.Accounts {
			fmt.Printf("%s \n", account.Name)
		}
	},
}

var accountCreateCmd = &cobra.Command{
	Use: "create",
	Run: func(cmd *cobra.Command, args []string) {
		//TODO: Move this
		if len(args) != 1 {
			os.Exit(1)
		}
		rootDir, err := cmd.Flags().GetString("target")
		if err != nil {
			log.Error(err)
		}
		os.Mkdir(filepath.Join(rootDir, args[0]), 0755)
	},
}
