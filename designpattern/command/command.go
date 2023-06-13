package command

import "fmt"

/**
Ordinary statements are perishable
 Cannot undo field assignment
 Cannot directly serialize invocation

Want an object that represents an operation
 person should change its age to value 22
 car should do explode()

Uses: GUI commands, multi-level undo/redo, macro recording and more!

An object which represents and instruction to perform a particular action.
Contains all the information necessary for the action to be taken.
*/

// Bank Account
var overdraftLimit = -500

type BankAccount struct {
	balance int
}

func (b *BankAccount) Deposit(amount int) {
	b.balance += amount
	fmt.Println("Deposit ", amount, "\b, balance is now ", b.balance)
}

func (b *BankAccount) Withx(amount int) bool {
	if b.balance-amount >= overdraftLimit {
		b.balance -= amount
		fmt.Println("Withdraw ", amount, "\b, balance is now ", b.balance)
		return true
	}
	return false
}

type Command interface {
	Call()
	Undo()
	Succeeded() bool
	SetSucceeded(value bool)
}

type Action int

const (
	Deposit Action = iota
	Withdraw
)

type BankAccountCommand struct {
	account   *BankAccount
	action    Action
	amount    int
	succeeded bool
}

func (b *BankAccountCommand) Succeeded() bool {
	return true
}

func (b *BankAccountCommand) SetSucceeded(value bool) {
	b.succeeded = value
}

func (b *BankAccountCommand) Undo() {
	if !b.succeeded {
		return
	}
	switch b.action {
	case Deposit:
		b.account.Withx(b.amount)
	case Withdraw:
		b.account.Deposit(b.amount)
	}
}

func (b *BankAccountCommand) Call() {
	switch b.action {
	case Deposit:
		b.account.Deposit(b.amount)
		b.succeeded = true
	case Withdraw:
		b.succeeded = b.account.Withx(b.amount)
	default:
		panic("Unsupported action")
	}
}
func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
	return &BankAccountCommand{account: account, action: action, amount: amount}
}

var _ Command = (*BankAccountCommand)(nil)

type CompositeBankAccountCommand struct {
	commands []Command
}

func (c *CompositeBankAccountCommand) Call() {
	for _, v := range c.commands {
		v.Call()
	}
}

func (c *CompositeBankAccountCommand) Undo() {
	for idx := range c.commands {
		c.commands[len(c.commands)-idx-1].Undo()
	}
}

func (c *CompositeBankAccountCommand) Succeeded() bool {
	for _, cmd := range c.commands {
		if !cmd.Succeeded() {
			return false
		}
	}
	return true
}

func (c *CompositeBankAccountCommand) SetSucceeded(value bool) {
	for _, cmd := range c.commands {
		cmd.SetSucceeded(value)
	}
}

var _ Command = (*CompositeBankAccountCommand)(nil)

type MoneyTransferCommand struct {
	CompositeBankAccountCommand
	from, to *BankAccount
	amount   int
}

func NewMoneyTransferCommand(from, to *BankAccount, amount int) *MoneyTransferCommand {
	c := &MoneyTransferCommand{from: from, to: to, amount: amount}
	c.commands = append(c.commands, NewBankAccountCommand(from, Withdraw, amount))
	c.commands = append(c.commands, NewBankAccountCommand(to, Deposit, amount))
	return c
}

func (m *MoneyTransferCommand) Call() {
	ok := true
	for _, cmd := range m.commands {
		if ok {
			cmd.Call()
			ok = cmd.Succeeded()
		} else {
			cmd.SetSucceeded(false)
		}
	}
}

func Start01() {
	b := &BankAccount{}
	cmd := NewBankAccountCommand(b, Deposit, 10000)
	cmd2 := NewBankAccountCommand(b, Withdraw, 500)
	cmd.Call()
	fmt.Println()
	cmd2.Call()
}

func Start() {
	from := &BankAccount{100}
	to := &BankAccount{0}
	tx := NewMoneyTransferCommand(from, to, 100)
	tx.Call()
}
