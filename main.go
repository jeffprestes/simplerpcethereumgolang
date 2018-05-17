package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	web3 "github.com/regcostajr/go-web3"
	"github.com/regcostajr/go-web3/dto"
	"github.com/regcostajr/go-web3/eth"
	"github.com/regcostajr/go-web3/providers"
)

var contract *eth.Contract

func main() {

	client := web3.NewWeb3(providers.NewHTTPProvider(os.Getenv("ETH_IP")+":"+os.Getenv("ETH_PORT"), 30, false))
	account := os.Getenv("ETH_ACCOUNT")
	pass := os.Getenv("ETH_SENHA")

	_, err := client.Personal.UnlockAccount(account, pass, 10)
	if err != nil {
		log.Fatalln(err.Error(), " - error unlocking account")
		return
	}

	accounts, err := client.Eth.ListAccounts()
	if err != nil {
		log.Fatalln(err.Error(), " - error listing accounts")
		return
	}

	for _, acc := range accounts {
		fmt.Println("Account: ", acc)
	}

	var abi string
	abi = `[{"constant":true,"inputs":[{"name":"visitor","type":"address"}],"name":"getMessageOfVisit","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"whoIsTheOwner","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[],"name":"kill","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"visitor","type":"address"},{"name":"message","type":"string"}],"name":"recordVisit","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"newOwner","type":"address"}],"name":"changeOwner","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"}]`
	contract, err = client.Eth.NewContract(abi)
	if err != nil {
		log.Fatalln(err.Error(), " - error initializing a contract")
		return
	}

	fmt.Println("Iniciando o servidor web...")
	http.HandleFunc("/", executeContractMethod)
	http.ListenAndServe(os.Getenv("WEB_IP")+":"+os.Getenv("WEB_PORT"), nil)
}

func executeContractMethod(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "<html><head><title>Calling smart contract via web</title><head><body>")

	transaction := new(dto.TransactionParameters)
	transaction.From = "0xa1a2bf87fd49e1d222353821b4335fc21a61880f"
	transaction.To = "0xfd244d32aD243C47C5122EA1EBE7b56e36e188d8"
	result, err := contract.Call(transaction, "getMessageOfVisit", "0xf50ec851faef19d723f63ee218a85622e9e810e8")
	if err != nil {
		fmt.Fprintln(w, "<h5>", err.Error(), " - error calling getMessageOfVisit contract method")
		return
	}
	if result != nil {
		msg, _ := result.ToComplexString()
		fmt.Fprintln(w, "<h1>Last visit message: ", msg.ToString(), "</h1>")
	}
	result, err = contract.Call(transaction, "whoIsTheOwner", nil)
	if err != nil {
		fmt.Fprintln(w, "<h5>", err.Error(), " - error calling whoIsTheOwner contract method</h5>")
		return
	}
	if result != nil {
		fmt.Fprintf(w, "<h1>Owner: %+v </h1>\n", result.Result)
	}
	fmt.Fprintln(w, "</body></html>")
}
