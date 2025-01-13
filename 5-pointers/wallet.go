package pointers

import "fmt"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

// Stringer interface
// type Stringer interface {
// 		String() string
// }
// A Stringer is a type that can describe itself as a string. 
// The fmt package (and many others) look for this interface to print values.
// The String method will be used to print values passed to %s in format strings.
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}