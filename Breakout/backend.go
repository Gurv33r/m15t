package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/Gurv33r/Spider/deck"
)

var draw, center *deck.Deck = deck.NewDeck(), deck.EmptyDeck()
var opp, player *deck.Player = deck.NewPlayer(), deck.NewPlayer()
var seven bool = false

func main() {

	// http.HandleFunc("/", index)
	// http.HandleFunc("/game", game)
	// fmt.Println("Server starting...")
	// http.ListenAndServe(":8080", nil)
	//trash := deck.EmptyDeck()
	draw.Shuffle()
	// fmt.Println(draw.Cards, center.Cards, trash.Cards)
	dealCards()
	fmt.Println("opp", opp, "\nplayer", player, "\ncenter", center)
}

func index(display http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(display, "<h1>Hello World</h1>")
}
func game(display http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(display, "<h1>Starting game<h1>")
	// draw, center := deck.NewDeck(), deck.EmptyDeck()
	// draw.Shuffle()
	// opp, player := deck.NewPlayer(), deck.NewPlayer()
	// //fmt.Println(draw.Cards, center.Cards, trash.Cards)
	// dealCards()
}

func dealCards() {
	for i := 0; i < 18; i++ {
		dealt := draw.Remove(draw.Top())
		if i < 9 {
			if i <= 2 {
				player.Hand = append(player.Hand, dealt)
			} else if i < 6 && i >= 3 {
				player.Pen[i-3] = dealt
			} else {
				player.Final[i-6] = dealt
			}
		} else {
			if i < 12 {
				opp.Hand = append(opp.Hand, dealt)
			} else if i < 15 && i >= 12 {
				opp.Pen[i-12] = dealt
			} else {
				opp.Final[i-15] = dealt
			}
		}
	}
	center.Collect(draw.Remove(draw.Top()))
}

func autoPlay(p *deck.Player, d *deck.Deck) {
	rand.Seed(time.Now().UnixNano())
	hand, playable := determineHand(p), []*deck.Card{}
	if hand == 1 { // hand case
		//find playable cards
		for _, card := range p.Hand {
			if judge(card) > 0 {
				playable = append(playable, card)
			}
		}
		if len(playable) == 0 {
			//collect center
			p.Hand = append(p.Hand, center.Cards...)
			center = deck.EmptyDeck()
		} else if len(playable) == 1 {
			center.Collect(playable[0])
			extract(p.Hand, playable[0])
		} else {
			willPlay := lowestCardOf(playable)
			center.Collect(willPlay)
			extract(p.Hand, willPlay)
		}
	} else if hand == 2 { // pen case
		//find playable cards
		for _, card := range p.Hand {
			if judge(card) > 0 {
				playable = append(playable, card)
			}
		}
		if len(playable) == 0 {
			//collect center
			p.Hand = append(p.Hand, center.Cards...)
			center = deck.EmptyDeck()
		} else if len(playable) == 1 {
			center.Collect(playable[0])
			extract(p.Pen, playable[0])
		} else {
			willPlay := lowestCardOf(playable)
			center.Collect(willPlay)
			extract(p.Pen, willPlay)
		}
	} else if hand == 3 { //final case
		// blind flip of the card
		chosenCardIndex := rand.Intn(len(p.Final))
		center.Collect()
		extract(p.Final, p.Final[chosenCardIndex])
	}
}
func determineHand(p *deck.Player) int {
	if p.Turn {
		if len(p.Hand) > 0 {
			return 1
		} else if len(p.Hand) == 0 && len(p.Pen) > 0 && len(p.Final) > 0 && draw.Size() == 0 {
			return 2
		} else if len(p.Hand) == 0 && len(p.Pen) == 0 && len(p.Final) > 0 && draw.Size() == 0 {
			return 3
		}
	}
	return 0
}

func judge(c *deck.Card) int {
	var top int = 0
	if seven {
		top = 7
	} else {
		top = center.Top().Num
	}

	if c.IsSpecial() {
		log.Println("card is special!")
		if c.Num == 2 {
			return 2
		} else if c.Num == 7 {
			seven = true
			return 7
		} else if c.Num == 10 {
			return 10
		}
	} else if c.Num > top {
		if seven {
			return 0
		}
		return 1
	}
	if seven {
		return 1
	}
	return 0
}
