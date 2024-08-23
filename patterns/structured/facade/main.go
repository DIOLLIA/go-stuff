package main

import (
	"fmt"
	"log"
)

func main() {

	walletFacade := newWalletFacade("eur", 500)

	if err := walletFacade.addMoneyToWallet("eur", 500, 10); err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	if err := walletFacade.deductMoneyFromWallet("eur", 500, 7); err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Printf("\n total balance is:%v", walletFacade.wallet.balance)
}
