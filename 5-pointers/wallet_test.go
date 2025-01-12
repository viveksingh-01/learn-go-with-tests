package pointers

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{balance: 0}
		wallet.Deposit(10)
		got := wallet.Balance()
		want := 10
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

// NEED FOR POINTERS
// When we run the test, it fails with the following error:
// got 0 want 10
// This is because the Deposit method does not actually modify the balance field of the Wallet struct.
// This is because Go copies the Wallet when we call Deposit, so the method is actually working on a 
// copy of the original Wallet.
// To fix this, we need to change the Deposit method to use a pointer receiver.
// This means that the method will be working with a pointer to the original Wallet, rather than a copy.
// We can do this by changing the receiver for the Deposit method from (w Wallet) to (w *Wallet).
}