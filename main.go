package main

import (
	"fmt"

	web3 "github.com/regcostajr/go-web3"
	"github.com/regcostajr/go-web3/providers"
)

func main() {

	client := web3.NewWeb3(providers.NewHTTPProvider("10.162.106.209:8545", 10, false))
	account := "0xf637c9c6f886f3edf046eb9b657c30c00fb602be"
	pass := "teste123"

	_, err := client.Personal.UnlockAccount(account, pass, 10)

	if err != nil {
		fmt.Println(err.Error(), " - error unlocking account")
		return
	}

	accounts, err := client.Eth.ListAccounts()
	if err != nil {
		fmt.Println(err.Error(), " - error listing accounts")
		return
	}

	for _, acc := range accounts {
		fmt.Println("Account: ", acc)
	}

}
