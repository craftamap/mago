package money

import (
	"encoding/json"
	"errors"
	"io"
)

var ErrNoTransactionID = errors.New("no ID found in Transaction")

// Transaction represents an financial transaction
type Transaction struct {
	ID			int		`json:"id"`
	Amount      float64 `json:"amount,omitempty"`
	Description string  `json:"description,omitempty"`
	Category    string  `json:"category,omitempty"`
	Account     *Account `json:"-"`
}

func FromReader(account *Account, reader io.Reader) (*Transaction, error) {
	transaction := Transaction{}
	err := json.NewDecoder(reader).Decode(&transaction)
	if (err != nil) {
		return nil, err
	}
	if transaction.ID == 0 {
		return nil, ErrNoTransactionID
	}
	transaction.Account = account
	return &transaction, nil
}
