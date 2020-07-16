package money

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Account struct {
	Name string
	Transactions []Transaction
}


func FromDirectory(path string) (*Account, error) {
	account := Account{}	
	// Read directory 
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	name := filepath.Base(path)	
	account.Name = name

	//TODO: Refactor as go routine
	for _, file := range files  {
		if file.IsDir() {
			continue
		}

		reader, err := os.Open(filepath.Join(path, file.Name()))
		if err != nil {
			log.Print(err)
			continue
		}

		transaction, err := FromReader(&account, reader)
		if err != nil {
			log.Print(err)
			continue
		}
		account.Transactions	= append(account.Transactions, *transaction)
	}

	return &account, nil
}

func (a *Account) Balance() float64 {
	var balance float64

	for _, transaction := range a.Transactions {
		balance += transaction.Amount
	}
		
	return balance
}
