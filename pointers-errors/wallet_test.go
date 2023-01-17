package wallet

import (
	"errors"
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, w Wallet, want Bitcoin) {
		t.Helper()
		got := w.Balance()
		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}

	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		fmt.Printf("the balance is now: %s \n", wallet.balance)
		assertBalance(t, wallet, 10)
	})
	t.Run("Withdrawal", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, 10)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)

		assertErrorContains(t, err, ErrInsufficientFunds)
	})
}

func assertErrorContains(t testing.TB, want error, got error) {
	t.Helper()
	if want == nil {
		t.Fatal("wanted an error but didn't get one")
	}
	if !errors.Is(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}
func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error %s", err)
	}
}
