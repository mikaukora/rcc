package cmd

import (
	"os"
)

const (
	environmentAccount = `RCC_CREDENTIALS_ID`
)

func AccountName() string {
	if len(accountName) > 0 {
		return accountName
	}
	return os.Getenv(environmentAccount)
}
