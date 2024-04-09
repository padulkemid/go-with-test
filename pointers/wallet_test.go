package pointers

import (
	"testing"
)

func assertNoError(t testing.TB, err error)  {
  t.Helper()

  if err != nil {
    t.Fatalf("got an error: %#v should not got one", err)
  }
}

func assertError(t testing.TB, err error, msg string)  {
  t.Helper()

  if err == nil {
    t.Fatal("should throw an error but no error thrown")
  }

  if err.Error() != msg {
    t.Errorf("got %q, want %q", err, msg)
  }
}
func assertBalance(t testing.TB, w Wallet, b Bitcoin) {
	t.Helper()

	got := w.Balance()

	if got != b {
		t.Errorf("got %s want %s", got, b)
	}
}

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(20))

		got := wallet.Balance()

		want := Bitcoin(20)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		w := Wallet{balance: Bitcoin(100)}
    err := w.Withdraw(Bitcoin(50))

		got := w.Balance()
		want := Bitcoin(50)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

    assertNoError(t, err)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		b := Bitcoin(100)
		w := Wallet{b}
		err := w.Withdraw(Bitcoin(200))

		assertBalance(t, w, b)
    assertError(t, err, "can't withdraw insufficient balance")
	})
}
