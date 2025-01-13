package pointers

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{balance: 0}
		wallet.Deposit(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(10)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}