package money

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"
)

var ErrNoTransactionID = errors.New("no ID found in Transaction")

type Amount int

func (a Amount) String() string {
	//TODO: This could add floating point errors
	inEuro := (float64(a) / 100.0)
	return fmt.Sprintf("%.2f", inEuro)
}

// Transaction represents an financial transaction
type Transaction struct {
	ID               int       `json:"id"`
	Amount           Amount    `json:"amount,omitempty"`
	Description      string    `json:"description,omitempty"`
	Category         string    `json:"category,omitempty"`
	Account          *Account  `json:"-"`
	CreationDate     time.Time `json:"creationDate"`
	ModificationDate time.Time `json:"modificationDate"`
}

func FromReader(account *Account, reader io.Reader) (*Transaction, error) {
	transaction := Transaction{}
	err := json.NewDecoder(reader).Decode(&transaction)
	if err != nil {
		return nil, err
	}
	if transaction.ID == 0 {
		return nil, ErrNoTransactionID
	}
	transaction.Account = account
	return &transaction, nil
}
