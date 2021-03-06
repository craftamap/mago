package mago

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type Account struct {
	Name         string
	Path         string
	Transactions map[int]Transaction
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
	account.Path = path

	account.Transactions = map[int]Transaction{}

	//TODO: Refactor as go routine
	for _, file := range files {
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
		account.Transactions[transaction.ID] = *transaction
	}

	return &account, nil
}

func (a *Account) Balance() Amount {
	var balance Amount

	for _, transaction := range a.Transactions {
		balance += transaction.Amount
	}

	return balance
}

func (a *Account) RemoveTransaction(transaction *Transaction) error {
	delete(a.Transactions, transaction.ID)

	err := os.Remove(filepath.Join(a.Path, strconv.Itoa(transaction.ID)))

	if err != nil {
		return err
	}
	return nil
}

func (a *Account) AddTransaction(transaction *Transaction) error {
	file, err := os.Create(filepath.Join(a.Path, strconv.Itoa(transaction.ID)))
	if err != nil {
		return err
	}
	err = json.NewEncoder(file).Encode(transaction)
	if err != nil {
		return err
	}

	a.Transactions[transaction.ID] = *transaction

	return nil
}

func (a *Account) CreateTransaction(amount Amount, description string, category string) (*Transaction, error) {
	//TODO: Do we really want to generate the ids like this?
	newId := time.Now().Unix()
	doBreak := false
	if len(a.Transactions) > 0 {
		for {
			for id := range a.Transactions {
				if id == int(newId) {
					newId += 1
					break
				}
				doBreak = true
			}
			if doBreak {
				break
			}
		}
	}

	transaction := Transaction{
		ID:               int(newId),
		Amount:           amount,
		Description:      description,
		Category:         category,
		Account:          a,
		CreationDate:     time.Now(),
		ModificationDate: time.Now(),
	}

	err := a.AddTransaction(&transaction)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}
