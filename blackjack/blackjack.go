package blackjack

const (
	other = 0
	two   = 2
	three = 3
	four  = 4
	five  = 5
	six   = 6
	seven = 7
	eight = 8
	nine  = 9
	ten   = 10
	jack  = 10
	queen = 10
	king  = 10
	ace   = 11
)

const (
	stand            = "S"
	hit              = "H"
	split            = "P"
	automaticallyWin = "W"
)

const blackjack = 21

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	default:
		return other
	case "two":
		return two
	case "three":
		return three
	case "four":
		return four
	case "five":
		return five
	case "six":
		return six
	case "seven":
		return seven
	case "eight":
		return eight
	case "nine":
		return nine
	case "ten":
		return ten
	case "jack":
		return jack
	case "queen":
		return queen
	case "king":
		return king
	case "ace":
		return ace
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	cv1 := ParseCard(card1)
	cv2 := ParseCard(card2)
	dcv := ParseCard(dealerCard)

	cvSum := cv1 + cv2

	if cv1 == ace && cv2 == ace {
		return split
	}

	if cvSum == blackjack {
		if dcv >= ten {
			return stand
		}
		return automaticallyWin
	}

	if cvSum >= 17 && cvSum <= 20 {
		return stand
	}

	if cvSum >= 12 && cvSum <= 16 {
		if dcv >= seven {
			return hit
		}
		return stand
	}

	return hit
}
