package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/craftamap/mago/money"
	"github.com/prometheus/common/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

func runAddCmd(cmd *cobra.Command, args []string) {
	//TODO: Make this beautiful, lol
	rootDir, err := cmd.Flags().GetString("target")
	if err != nil {
		log.Error(err)
	}
	manager, err := money.FromRoot(rootDir)
	if err != nil {
		log.Error(err)
	}
	fmt.Println("Choose the account you want to add a transaction to:")
	i := 1
	accountMap := map[int]string{}
	for account := range manager.Accounts {
		fmt.Printf(" [%d] %s\n", i, account)
		accountMap[i-1] = account
		i++
	}
	var integer = 0
	for {
		lineReader := bufio.NewReader(os.Stdin)
		text, err := lineReader.ReadString('\n')
		if err != nil {
			log.Error(err)
		}
		text = strings.TrimSpace(text)
		integer, err = strconv.Atoi(text)
		if err != nil {
			log.Error(err)
		}
		if integer >= 1 && integer <= len(manager.Accounts) {
			break
		}
	}
	account := manager.Accounts[accountMap[integer-1]]

	fmt.Println("What do you want to add?")

	var amount money.Amount
	for {
		lineReader := bufio.NewReader(os.Stdin)
		text, err := lineReader.ReadString('\n')
		if err != nil {
			log.Error(err)
		}
		text = strings.TrimSpace(text)
		amountI, err := strconv.Atoi(text)
		if err != nil {
			log.Error(err)
			continue
		}
		amount = money.Amount(amountI)
		break
	}

	fmt.Println("Whats the description?")
	var description string
	{
		lineReader := bufio.NewReader(os.Stdin)
		text, err := lineReader.ReadString('\n')
		if err != nil {
			log.Error(err)
		}
		description = strings.TrimSpace(text)

	}

	account.CreateTransaction(amount, description, "")

}

var addCmd = &cobra.Command{
	Use: "add",
	Run: runAddCmd,
}
