package memento

import "fmt"

/**
An object or system goes through changes
There are different ways of navigating those changes
One way is to record evey change and teach ta command to undo itself
Another is to simply save snapshots of the system

A token representing the system state.
Lets us roll back to the state when the token was generated.
May or may not directly expose state information.
*/

type Memento struct {
	Balance int
}

type BankAccount struct {
	balance int
}

func NewBankAccount(balance int) (*BankAccount, *Memento) {
	return &BankAccount{balance}, &Memento{balance}
}

func (b *BankAccount) Deposit(amount int) *Memento {
	b.balance += amount
	return &Memento{b.balance}
}

func (b *BankAccount) Restore(m *Memento) {
	b.balance = m.Balance
}

func Ex01() {
	ba, m0 := NewBankAccount(100)

	m1 := ba.Deposit(50)
	m2 := ba.Deposit(25)

	fmt.Println(ba)

	ba.Restore(m1)
	fmt.Println(ba)

	ba.Restore(m2)
	fmt.Println(ba)

	ba.Restore(m0)
	fmt.Println(ba)
}

type Memento2 struct {
	Balance int
}

type BankAccount2 struct {
	balance int
	changes []*Memento2
	current int
}

func (b *BankAccount2) String() string {
	return fmt.Sprint("Balance = $", b.balance, ", current = ", b.current)
}

func NewBankAccount2(balance int) *BankAccount2 {
	b := &BankAccount2{balance: balance}
	b.changes = append(b.changes, &Memento2{balance})
	return b
}

func (b *BankAccount2) Deposit(amount int) *Memento2 {
	b.balance += amount
	m := &Memento2{b.balance}
	b.changes = append(b.changes, m)
	b.current++
	fmt.Println("Deposited", amount, ", balance is now", b.balance)
	return m
}

func (b *BankAccount2) Restore(m *Memento2) {
	if m != nil {
		b.balance = m.Balance
		b.changes = append(b.changes, m)
		b.current = len(b.changes) - 1
	}
}

func (b *BankAccount2) Undo() *Memento2 {
	if b.current > 0 {
		b.current--
		m := b.changes[b.current]
		b.balance = m.Balance
		return m
	}
	return nil
}

func (b *BankAccount2) Redo() *Memento2 {
	if b.current+1 < len(b.changes) {
		b.current++
		m := b.changes[b.current]
		b.balance = m.Balance
		return m
	}
	return nil
}

func Start() {
	ba := NewBankAccount2(100)

	ba.Deposit(50)
	ba.Deposit(25)

	fmt.Println(ba)

	ba.Undo()
	fmt.Println("Undo 1:", ba)

	ba.Undo()
	fmt.Println("Undo 2:", ba)

	ba.Redo()
	fmt.Println("Redo :", ba)
}

/**
Memento vs Flyweight
- Both pattern provide a 'token' clients can hold on to
- Memento is used only to be fed back into the system
- A flyweight is similar to an ordinary reference to object

Mementos are used to roll back states arbitrarily
A memento is simply a token/handle with no methods of its own
A memento is not required to expose directly the state(s) to which it reverts the system
Can be used to implement undo/redo
*/
