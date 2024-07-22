package main

import "fmt"

type WalletFacade struct {
	account      *Account
	wallet       *Wallet
	securityCode *SecurityCode
	notification *Notification
	ledger       *Ledger
}

func newWalletFacade(accountId string, code int) *WalletFacade {
	fmt.Println("Starting create account")
	walletFacade := WalletFacade{
		account:      newAccount(accountId),
		ledger:       &Ledger{},
		notification: &Notification{},
		securityCode: newSecurityCode(code),
		wallet:       newWallet(),
	}
	fmt.Println("Account created")

	return &walletFacade
}

func (w *WalletFacade) addMoneyToWallet(accountID string, securityCode, amount int) error {
	fmt.Print("Starting add money to wallet")
	if err := w.account.checkAccount(accountID); err != nil {
		return err
	}
	if err := w.securityCode.checkSecurityCode(securityCode); err != nil {
		return err
	}
	w.wallet.creditBalance(amount)
	w.notification.sendCreditWalletNotification()
	w.ledger.makeEntry(accountID, "CREDIT", amount)
	return nil
}

func (w *WalletFacade) deductMoneyFromWallet(accountID string, securityCode, amount int) error {
	fmt.Print("Starting add money to wallet")
	if err := w.account.checkAccount(accountID); err != nil {
		return err
	}
	if err := w.securityCode.checkSecurityCode(securityCode); err != nil {
		return err
	}
	err := w.wallet.debitBalance(amount)
	if err != nil {
		return err
	}
	w.notification.sendDebitWalletNotification()
	w.ledger.makeEntry(accountID, "DEBIT", amount)
	return nil
}
