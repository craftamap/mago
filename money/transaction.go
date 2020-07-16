package money

import (
	"encoding/json"
	"io"
)

// Transaction represents an financial transaction
type Transaction struct {
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
	transaction.Account = account
	return &transaction, nil
}
