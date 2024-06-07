package accounts

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"golang.org/x/text/message"
)

type customError struct {
	message string
	code int
}

type Account struct {
	accNumber    int32
	name         string
	acctype      string
	balance      float64
	dateOfBirth  time.Time
	pin          int
	pinResetFlag bool
}

func (account *Account) SetAccNumber(n int32) {
	account.accNumber = n
}

func (account *Account) SetName(name string) {
	account.name = name
}

func (account *Account) SetAccType(Type string) {
	account.acctype = Type
}

func (account *Account) SetBalance(balance float64) {
	account.balance = balance
}

func (account *Account) SetDateOfBirth(date string) error {
	birthdate, err := time.Parse("2006-01-02", date)
	if err != nil {
		account.dateOfBirth = birthdate
	}
	return err
}

func (account *Account) setPin() {
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1000 and 9999 (inclusive)
	num := rand.Intn(9000) + 1000
	account.pin = num
}

func (account *Account) SetResetPinFlag() {
	account.pinResetFlag = false
}

func (account *Account) GetName() string {
	return account.name
}

func (account *Account) GetAccNumber() int32 {
	return account.accNumber
}

func (account *Account) GetAccType() string {
	return account.acctype
}

func (account *Account) GetBalance() float64 {
	return account.balance
}

func (account *Account) GetPin() int {
	return account.pin
}

func (account *Account) GetDateOfBirth() time.Time {
	return account.dateOfBirth
}

func OpenAccount(accNum int32) Account {
	reader := bufio.NewReader(os.Stdin)
	acc := Account{}

	acc.SetAccNumber(accNum)
	fmt.Println("Enter your name:")
	var name string
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	acc.SetName(name)

	fmt.Println("enter account type from the following ( Current / Savings ):")
	Type, _ := reader.ReadString('\n')
	Type = strings.TrimSpace(Type)
	acc.SetAccType(Type)

	acc.SetBalance(5000)

DOB:
	var dob string
	fmt.Println("Enter date of Birth (2000-12-01):")
	fmt.Scan(&dob)
	err := acc.SetDateOfBirth(dob)
	if err != nil {
		fmt.Println("Enter date in valid date format(2000-12-01).")
		goto DOB
	}

	acc.setPin()
	acc.SetResetPinFlag()

	return acc

}

func Withdraw(bank *map[int32]Account, accnumber int32, amount float64) {

	holder := (*bank)[accnumber]
	if holder.balance-amount > 0 {
		holder.balance = holder.balance - amount
		(*bank)[accnumber] = holder
	} else {
		fmt.Println("Insufficient Funds.")
	}
	fmt.Println("Current balance is", holder.balance)
}

func Deposit(bank *map[int32]Account, accnumber int32, amount float64) {
	// var amount float64
	// fmt.Println("Enter ammount to deposit :")
	// fmt.Scan(&amount)
	holder := (*bank)[accnumber]
	holder.balance = holder.balance + amount
	(*bank)[accnumber] = holder
	fmt.Println("Amount deposited successfully.")
	fmt.Println("Current balance is", holder.balance)
}

func CheckBalance(bank *map[int32]Account, accnumber int32) {
	holder := (*bank)[accnumber]
	fmt.Printf("Current balance is %v.", holder.balance)
}

func CheckFlag(bank *map[int32]Account, accnumber int32) {
	holder := (*bank)[accnumber]
	// if !holder.pinResetFlag {
	// 	var newPin int
	// 	fmt.Println("Enter new 4 digit Pin:")
	// start:
	// 	fmt.Scan(newPin)
	// 	if newPin != holder.pin {
	// 		holder.pin = newPin
	// 		(*bank)[accnumber] = holder
	// 	} else {
	// 		fmt.Println("Enter valid pin.")
	// 		goto start
	// 	}
	// }

	if holder.pinResetFlag == false {
		for {
			var newPin int
			fmt.Println("Enter new pin:")
			fmt.Scan(&newPin)

			if holder.GetPin() != newPin {
				holder.pin = newPin
				holder.pinResetFlag = true
				(*bank)[accnumber] = holder
				break
			} else {
				continue
			}
		}
	}

}

func CheckValidPin(bank *map[int32]Account, accnumber int32) bool {
	holder := (*bank)[accnumber]
	var pin int
	fmt.Println("Enter your pin:")
	fmt.Scan(&pin)
	if holder.GetPin() == pin {
		return true
	} else {
		return false
	}
}
