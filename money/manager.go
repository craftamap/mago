package money

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

type Manager struct {
	RootDir  string
	Accounts map[string]Account
}

func FromRoot(path string) (*Manager, error) {
	manager := Manager{
		RootDir:  path,
		Accounts: map[string]Account{},
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			account, err := FromDirectory(filepath.Join(path, file.Name()))
			if err != nil {
				log.Println(err)
				continue
			}
			manager.Accounts[account.Name] = *account
		}
	}

	return &manager, nil
}
