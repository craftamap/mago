package cmd

import (
	"fmt"

	"github.com/craftamap/mago/money"
	"github.com/prometheus/common/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(balanceCmd)
}

func runBalanceCmd(cmd *cobra.Command, args []string) {
	rootDir, err := cmd.Flags().GetString("target")	
	if err != nil {
		log.Error(err)
	}
	manager, err := money.FromRoot(rootDir)
	if err != nil {
		log.Error(err)
	}

	for _, account := range manager.Accounts {
		fmt.Printf("%s: %s\n", account.Name, account.Balance())	
	}
}

var balanceCmd = &cobra.Command{
	Use: "balance",
	Run: runBalanceCmd,
}
