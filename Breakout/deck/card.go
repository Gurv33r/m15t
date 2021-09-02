package deck

import "strings"

type Card struct {
	Suit  string `json: "suit"`
	Value string `json: "value"`
	Num   int    `json: "number"`
}

func NewCard(suit, value string, n int) *Card {
	return &Card{suit, value, n}
}
func (c *Card) Color() string {
	if c.Value == "7" || c.Value == "2" || c.Value == "10" {
		return "golden"
	} else if c.Value == "♠" || c.Value == "♣" {
		return "black"
	} else {
		return "red"
	}
}
func (c *Card) String() string {
	return c.Value + " " + c.Suit
}
func (c *Card) IsSpecial() bool {
	return strings.Compare(c.Color(), "golden") == 0
}
