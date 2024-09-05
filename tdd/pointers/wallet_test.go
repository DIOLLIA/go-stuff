package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit and balance", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		assertBalance(t, wallet.Balance(), Euro(10))
	})
	t.Run("withdrawal and balance", func(t *testing.T) {
		wallet := Wallet{Euro(10)}
		err := wallet.Withdraw(10)
		assertNoError(t, err)
		assertBalance(t, wallet.Balance(), Euro(0))
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Euro(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Euro(100))

		assertBalance(t, wallet.balance, startingBalance)
		assertError(t, err, "cannot withdraw, insufficient funds")
	})
}

func assertError(t testing.TB, result error, expected string) {
	t.Helper()
	if result == nil {
		t.Fatal("wanted an error but didn't get one")
	}
	if result.Error() != expected {
		t.Errorf("expected %q result %q", expected, result)
	}
}
func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("got an error but didn't want one")
	}
}

func assertBalance(t testing.TB, result, expected Euro) {
	t.Helper()
	if expected != result {
		t.Errorf("expected %q result %q", expected, result) //by changing %d to %s we can use the String() method
	}
}
