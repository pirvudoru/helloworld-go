package crypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		assert.Equal(t, want, got)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}

		err := wallet.Withdraw(Bitcoin(5))

		got := wallet.Balance()
		want := Bitcoin(5)

		assert.Nil(t, err)
		assert.Equal(t, want, got)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}

		result := wallet.Withdraw(Bitcoin(142))

		assert.NotNil(t, result.Error())
	})
}
