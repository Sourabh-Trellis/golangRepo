package main

import (
	"bankingSystem/accounts"
	"fmt"
)

var accountNumber int32

func main() {

	// accounts.Abc()
	var accNum int32 = 174832
	bank := make(map[int32]accounts.Account)

	fmt.Println("Welcome to bank")
	fmt.Println("Choose from following options:")
	for {
		fmt.Println("------------------------------------------------------------")
		var choice int
		fmt.Println("1)Open new Account\n2)Withdraw Amount\n3)Deposit amount\n4)Check balance\n5)Transfer Funds\n6)Show all accounts\n7)Exit")
		fmt.Scan(&choice)
		switch choice {
		case 1:

			acc := accounts.OpenAccount(accNum)
			bank[acc.GetAccNumber()] = acc
			accNum++
			fmt.Println("Account crated successfully.\nDefault pin to access your account :", acc.GetPin())

		case 2:
			var acc int32
			fmt.Println("Enter your accout number")
			fmt.Scan(&acc)
			if _, ok := bank[acc]; ok {
				check := accounts.CheckValidPin(&bank, acc)
				if check {
					accounts.CheckFlag(&bank, acc)
					fmt.Println("Enter amount to withdraw:")
					var amount float64
					fmt.Scan(&amount)
					accounts.Withdraw(&bank, acc, amount)
				} else {
					fmt.Println("Invalid pin.")
				}
			} else {
				fmt.Print("Invalid Account Number.")
			}

		case 3:
			// var accountNumber int32
			fmt.Println("Ente your account number:")
			fmt.Scan(&accountNumber)
			if _, ok := bank[accountNumber]; ok {
				check := accounts.CheckValidPin(&bank, accountNumber)
				if check {
					accounts.CheckFlag(&bank, accountNumber)
					fmt.Println("Enter amount you want to deposit")
					var amount float64
					fmt.Scan(&amount)
					accounts.Deposit(&bank, accountNumber, amount)
				} else {
					fmt.Println("Invalid pin.")
				}
			} else {
				fmt.Println("Invalid Account Number.")
			}
		case 4:
			// var accountNumber int32
			fmt.Println("Enter your account number:")
			fmt.Scan(&accountNumber)
			if _, ok := bank[accountNumber]; ok {
				check := accounts.CheckValidPin(&bank, accountNumber)
				if check {
					accounts.CheckFlag(&bank, accountNumber)
					accounts.CheckBalance(&bank, accountNumber)
				} else {
					fmt.Println("Invalid Pin.")
				}
			} else {
				fmt.Println("Invalid Account Number.")
			}
		case 5:
			var sender int32
			var reciever int32
			fmt.Println("Enter your account number:")
			fmt.Scan(&sender)
			fmt.Println("Enter recipient's account number:")
			fmt.Scan(&reciever)
			if _, ok := bank[sender]; ok {
				check := accounts.CheckValidPin(&bank, sender)
				if check {
					accounts.CheckFlag(&bank, sender)
					if _, ok := bank[reciever]; ok {
						fmt.Println("Enter amount to transfer:")
						var amount float64
						fmt.Scan(&amount)
						accounts.Withdraw(&bank, sender, amount)
						accounts.Deposit(&bank, reciever, amount)
					} else {
						fmt.Println("Invalid recipient account number.")
					}
				}
			} else {
				fmt.Println("Invalid sender account number.")
			}

		case 6:
			fmt.Printf("ACCOUNT\t\t\tHOLDER\t\t\tACCOUNT\t\t\tCURRENT\t\t\tBIRTH\n")
			fmt.Printf("NUMBER\t\t\tNAME\t\t\tTYPE\t\t\tBALANCE\t\t\tDATE\n\n")

			for k, v := range bank {
				fmt.Printf("%v\t\t\t%v\t\t\t%v\t\t\t%v\t\t\t%v\n", k, v.GetName(), v.GetAccType(), v.GetBalance(), v.GetDateOfBirth())
			}
		case 7:
			return

		}
	}
}

// ////////////////////////////////////////////////////////
// type account struct {
// 	accNumber int32
// 	name      string
// 	acctype   string
// 	balance   float64
// }

// func (account *account) setAccNumber(n int32) {
// 	account.accNumber = n
// }

// func (account *account) setName(name string) {
// 	account.name = name
// }

// func (account *account) setAccType(Type string) {
// 	account.acctype = Type
// }

// func (account *account) setBalance(balance float64) {
// 	account.balance = balance
// }

// func (account *account) getAccNumber() int32 {
// 	return account.accNumber
// }

// func openAccount(accNum int32) account {
// 	reader := bufio.NewReader(os.Stdin)
// 	acc := account{}

// 	acc.setAccNumber(accNum)

// 	fmt.Println("Enter your name:")
// 	var name string
// 	name, _ = reader.ReadString('\n')
// 	name = strings.TrimSpace(name)
// 	acc.setName(name)

// 	fmt.Println("enter account type from the following ( Current / Savings ):")
// 	Type, _ := reader.ReadString('\n')
// 	Type = strings.TrimSpace(Type)
// 	acc.setAccType(Type)

// 	acc.setBalance(1000)
// 	return acc
// }

// func withdraw(bank *map[int32]account, accnumber int32) {
// 	var amount float64
// 	fmt.Println("Enter ammount to withdraw :")
// 	fmt.Scan(&amount)
// 	holder := (*bank)[accnumber]
// 	if holder.balance-amount > 1000 {
// 		holder.balance = holder.balance - amount
// 		(*bank)[accnumber] = holder

// 	} else {
// 		fmt.Println("Insufficient Balance.Account should maintain minimum balance Rs.1000")
// 	}
// 	fmt.Println("Current balance is", holder.balance)
// }

// func deposit(bank *map[int32]account, accnumber int32) {
// 	var amount float64
// 	fmt.Println("Enter ammount to deposit :")
// 	fmt.Scan(&amount)
// 	holder := (*bank)[accnumber]
// 	holder.balance = holder.balance + amount
// 	(*bank)[accnumber] = holder
// 	fmt.Println("Amount deposited successfully.")
// 	fmt.Println("Current balance is", holder.balance)
// }

// func checkBalance(bank *map[int32]account, accaccount int32) {
// 	holder := (*bank)[accaccount]
// 	fmt.Printf("Current balance is %v.", holder.balance)
// }

// func (num) isAccountValid(bank *map[int32]account) bool {
// 	if _, ok := (*bank)[num]; ok {
// 		return true
// 	} else {
// 		return false
// 	}
// }
