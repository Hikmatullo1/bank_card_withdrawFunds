package card

import (
	"bank/pkg/bank/types"
)

func Withdraw(card types.Card, amount types.Money) types.Card {

	if card.Active == false {
		return card
	}

	if card.Balance < amount {
		return card
	}

	if amount <= 0 {
		return card
	}
	const limit = 20_000_00
	if amount > limit {
		return card
	}

	card.Balance -= amount
	return card

}
