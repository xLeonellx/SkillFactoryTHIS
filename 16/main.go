package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

type Client struct {
	cash int
}

var mu = sync.Mutex{}

//создание методов для структуры клиента
func (c *Client) Deposit(amount int) {
	mu.Lock()
	c.cash += amount
	mu.Unlock()
}
func (c *Client) Withdrawal(amount int) error {
	if amount > c.cash {
		return errors.New("Недостаточно средств на балансе, операция не выполнена")
	} else {
		mu.Lock()
		c.cash -= amount
		mu.Unlock()
		return nil
	}
}
func (c Client) Balance() int {
	return c.cash
}

//реализация интерфейса
type BankClient interface {
	// Deposit deposits given amount to clients account
	Deposit(amount int)

	// Withdrawal withdraws given amount from clients account.
	// return error if clients balance less the withdrawal amount
	Withdrawal(amount int) error

	// Balance returns clients balance
	Balance() int
}

func PrintBalance(b BankClient) {
	fmt.Println("Ваш текущий баланс:", b.Balance())
}
func BankDeposit(b BankClient) {
	var amount int
	fmt.Println("Сколько средств внести на счет?")
	_, err := fmt.Scanln(&amount)
	if err != nil {
		fmt.Println("Вы ввели не целое число, попробуйте использовать команду снова")
	} else {
		b.Deposit(amount)
	}
}
func BankWithdrawal(b BankClient) {
	var amount int
	fmt.Println("Сколько средств снять со счета?")
	_, scanErr := fmt.Scanln(&amount)
	if scanErr != nil {
		fmt.Println("Вы ввели не целое число, попробуйте использовать команду снова")
	}
	err := b.Withdrawal(amount)
	if err != nil {
		fmt.Println(err)
	}
}

//функция рандомного числа в выбранном диапозоне
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

//функция зачисления на счет случайной суммы до 10
func income(destination *int) {
	for {
		mu.Lock()
		*destination += randInt(1, 10)
		//fmt.Println(*destination)
		mu.Unlock()
		val := randInt(500, 1000)
		time.Sleep(time.Duration(val) * time.Millisecond)
	}
}

//функция снятия с счета случайной суммы до 5
func subtraction(destination *int) {
	for {
		num := randInt(1, 5)
		mu.Lock()
		if num > *destination {
			fmt.Println("Невозможно автоматическое снятие средств с счета, недостаточно средств на балансе")
			mu.Unlock()
		} else {
			*destination -= num
			mu.Unlock()
			//fmt.Println(*destination)
			val := randInt(500, 1000)
			time.Sleep(time.Duration(val) * time.Millisecond)
		}
	}
}
func main() {
	rand.Seed(time.Now().Unix())

	//инициализация юзера банка
	User := &Client{
		cash: 0,
	}

	//запуск 10 горутин, которые зачисляют на счет клиента сумму до 10
	for i := 0; i < 10; i++ {
		go income(&User.cash)
	}

	//запуск 5 горутин, которые снимают с счета клиента сумму до 5
	for i := 0; i < 5; i++ {
		go subtraction(&User.cash)
	}

	//реализация консольных команд
	HelloText := "Hello. You can use commands: balance, deposit, withdrawal, exit"
	fmt.Println(HelloText)
	var command string
	for {
		fmt.Scanln(&command)
		switch command {
		case "balance":
			PrintBalance(User)
		case "deposit":
			BankDeposit(User)
		case "withdrawal":
			BankWithdrawal(User)
		case "exit":
			os.Exit(1)
		default:
			fmt.Println("Unsupported command. You can use commands: balance, deposit, withdrawal, exit")

		}
	}
}
