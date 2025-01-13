package pointers

type Wallet struct {
	balance int
}

// Go is a "pass by value" language therefore, when you call a function or a method, the arguments are copied.
// If you pass a variable to a function, and the function modifies the variable,
// the changes will not be reflected in the original variable.
// To fix this, we need to change the Deposit method to use a pointer receiver.
// This means that the method will be working with a pointer to the original Wallet, rather than a copy.
// We can do this by changing the receiver for the Deposit method from (w Wallet) to (w *Wallet).

func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

// Now we might wonder, why did the test pass? We didn't dereference the pointer in the function, like so:
// func (w *Wallet) Deposit(amount int) {
// 		(*w).balance += amount
// }
// and seemingly addressed the object directly. In fact, the code above using (*w) is absolutely valid.
// However, the makers of Go deemed this notation cumbersome, so the language permits us to write w.balance,
// without an explicit dereference.
// These pointers to structs even have their own name: 'struct pointers' and they are automatically dereferenced.
// This is why we didn't need to change the calls to Deposit in our tests.

// Technically you do not need to change Balance to use a pointer receiver as
// taking a copy of the balance is fine. However, by convention you should
// keep your method receiver types the same for consistency.
func (w *Wallet) Balance() int {
	return w.balance
}