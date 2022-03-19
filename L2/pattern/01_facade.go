package main

/*
	Фасад — это структурный паттерн проектирования, который предоставляет простой интерфейс к сложной системе классов, библиотеке или фреймворку.
*/

import (
	"fmt"
	"log"
)

type Account struct {
	ID string
}

func NewAccount(id string) *Account {
	return &Account{
		ID: id,
	}
}

func (a *Account) CheckAccount(id string) error {
	if a.ID != id {
		return fmt.Errorf("Account ID is incorrect!")
	}

	log.Println("Account verified!")
	return nil
}

type Wallet struct {
	Balance int
}

func NewWallet() *Wallet {
	return &Wallet{
		Balance: 0,
	}
}

func (w *Wallet) Deposit(amount int) {
	w.Balance += amount
	log.Println("Wallet balance successfully replenished!")
}

func (w *Wallet) Withdraw(amount int) error {
	if w.Balance < amount {
		return fmt.Errorf("Balance is not sufficient!")
	}

	w.Balance -= amount
	log.Println("Amount successfully withdrawn from wallet balance!")
	return nil
}

type Security struct {
	Code int
}

func NewSecurity(code int) *Security {
	return &Security{
		Code: code,
	}
}

func (sc *Security) CheckCode(code int) error {
	if sc.Code != code {
		return fmt.Errorf("Security code is incorrect!")
	}

	log.Println("Security code verified!")
	return nil
}

type Notification struct{}

func (n *Notification) SendWalletDepositNotification() {
	log.Println("Wallet deposit notification sent!")
}

func (n *Notification) SendWalletWithdrawNotification() {
	log.Println("Wallet withdraw notification sent!")
}

type WalletFacade struct {
	Account      *Account
	Wallet       *Wallet
	Security     *Security
	Notification *Notification
}

func NewWalletFacade(id string, code int) *WalletFacade {
	log.Println("Start creating an account")
	defer log.Println("Account created successfully!")
	return &WalletFacade{
		Account:      NewAccount(id),
		Wallet:       NewWallet(),
		Security:     NewSecurity(code),
		Notification: &Notification{},
	}
}

func (wf *WalletFacade) DepositMoneyToWallet(accountID string, securityCode int, amount int) error {
	log.Println("Starting deposit money to wallet")
	if err := wf.Account.CheckAccount(accountID); err != nil {
		return err
	}
	if err := wf.Security.CheckCode(securityCode); err != nil {
		return err
	}

	wf.Wallet.Deposit(amount)
	wf.Notification.SendWalletDepositNotification()
	return nil
}

func (wf *WalletFacade) WithdrawMoneyFromWallet(accountID string, securityCode int, amount int) error {
	log.Println("Starting withdraw money from wallet")
	if err := wf.Account.CheckAccount(accountID); err != nil {
		return err
	}
	if err := wf.Security.CheckCode(securityCode); err != nil {
		return err
	}

	if err := wf.Wallet.Withdraw(amount); err != nil {
		return err
	}

	wf.Notification.SendWalletWithdrawNotification()
	return nil
}

func main() {
	walletFacade := NewWalletFacade("JOE43P32", 8439)

	if err := walletFacade.DepositMoneyToWallet(walletFacade.Account.ID, walletFacade.Security.Code, 1000); err != nil {
		log.Fatalf("Error: %s\n", err)
	}

	if err := walletFacade.WithdrawMoneyFromWallet(walletFacade.Account.ID, walletFacade.Security.Code, 500); err != nil {
		log.Fatalf("Error: %s\n", err)
	}
}

// 2022/03/03 17:15:50 Start creating an account
// 2022/03/03 17:15:51 Account created successfully!
// 2022/03/03 17:15:51 Starting deposit money to wallet
// 2022/03/03 17:15:51 Account verified!
// 2022/03/03 17:15:51 Security code verified!
// 2022/03/03 17:15:51 Wallet balance successfully replenished!
// 2022/03/03 17:15:51 Wallet deposit notification sent!
// 2022/03/03 17:15:51 Starting withdraw money from wallet
// 2022/03/03 17:15:51 Account verified!
// 2022/03/03 17:15:51 Security code verified!
// 2022/03/03 17:15:51 Amount successfully withdrawn from wallet balance!
// 2022/03/03 17:15:51 Wallet withdraw notification sent!
