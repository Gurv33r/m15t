package deck

import (
	"math/rand"
	"strings"
	"time"
)

type Deck struct {
	Cards []*Card `json: "cards"`
}

func NewDeck() *Deck {
	suits, values := [4]string{"♠", "♥", "♦", "♣"}, strings.Split("2,3,4,5,6,7,8,9,10,J,Q,K,A", ",")
	deck := []*Card{}
	for _, suit := range suits {
		for n, value := range values {
			deck = append(deck, NewCard(suit, value, n+1))
		}
	}
	return &Deck{deck}
}
func EmptyDeck() *Deck {
	return &Deck{[]*Card{}}
}

func (d *Deck) Size() int {
	return len(d.Cards)
}
func (d *Deck) Top() *Card {
	return d.Cards[0]
}
func (d *Deck) Collect(cards ...*Card) {
	if len(cards) == 1 {
		d.Cards = append(cards, d.Cards...)
	} else {
		d.Cards = append(d.Cards, cards...)
	}
}
func (d *Deck) Shuffle() {
	//new seed every time Shuffle() is called to prevent the same deck from being used over and over agagin
	rand.Seed(time.Now().UnixNano())
	for i := d.Size() - 1; i > 0; i-- {
		target := rand.Intn(i)
		oldValue := d.Cards[target]
		d.Cards[target] = d.Cards[i]
		d.Cards[i] = oldValue
	}
}
func (d *Deck) Remove(c *Card) *Card {
	//find the index
	i := d.indexOf(c)
	if i >= 0 {
		rv := d.Cards[i]
		//splice it out
		d.Cards = append(d.Cards[:i], d.Cards[i+1:]...)
		return rv
	} else {
		return nil
	}
}

func (d *Deck) indexOf(c *Card) int {
	for i, card := range d.Cards {
		if card.Suit == c.Suit && c.Value == card.Value {
			return i
		}
	}
	return -1
}
