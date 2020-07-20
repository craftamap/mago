package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	money "github.com/craftamap/mago"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

func runAddCmd(cmd *cobra.Command, args []string) {
	var amount money.Amount
	var accountName string
	var description string
	var categories []string
	if len(args) > 0 {
		amountF, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			log.Fatal(err)
		}
		amount = money.Amount(amountF * 100)
	}
	if len(args) > 1 {
		accountName = args[1]
	}
	if len(args) > 2 {
		descriptionAndTags := args[2:]
		for _, word := range descriptionAndTags {
			//TODO: We propably do not want to append to the description after the first tag ist
			if strings.HasPrefix(word, "+") {
				categories = append(categories, word)
			} else {
				description = description + " " + word
			}
		}
	}
	//TODO: Make this beautiful, lol
	rootDir, err := cmd.Flags().GetString("target")
	if err != nil {
		log.Fatal(err)
	}
	manager, err := money.FromRoot(rootDir)
	if err != nil {
		log.Fatal(err)
	}
	var account money.Account
	if accountName == "" {
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
				log.Fatal(err)
			}
			text = strings.TrimSpace(text)
			integer, err = strconv.Atoi(text)
			if err != nil {
				log.Fatal(err)
			}
			if integer >= 1 && integer <= len(manager.Accounts) {
				break
			}
		}
		account = manager.Accounts[accountMap[integer-1]]
	} else {
		var ok bool
		account, ok = manager.Accounts[accountName]
		if !ok {
			log.Fatal("Account not found")
			return
		}
	}

	if amount == 0 {
		fmt.Println("How much do you want to add?")
		for {
			lineReader := bufio.NewReader(os.Stdin)
			text, err := lineReader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			text = strings.TrimSpace(text)
			amountF, err := strconv.ParseFloat(text, 64)
			if err != nil {
				log.Fatal(err)
				continue
			}
			amount = money.Amount(amountF / 100)
			break
		}
	}

	if description == "" {
		fmt.Println("Whats the description?")
		{
			lineReader := bufio.NewReader(os.Stdin)
			description, err = lineReader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	description = strings.TrimSpace(description)
	_, err = account.CreateTransaction(amount, description, strings.Join(categories, ","))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Created Transaction over %s in Account %s with the description \"%s\" and the following categories: %s", amount, account.Name, description, categories)

}

var addCmd = &cobra.Command{
	Use: "add [amount [account [description [+tags]]]]",
	Run: runAddCmd,
}
