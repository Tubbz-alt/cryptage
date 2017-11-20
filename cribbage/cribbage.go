package cribbage

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/whereswaldon/cryptage/card"
	"strings"
)

type Card struct {
	Suit string
	Rank string
}

func (c *Card) MarshalText() ([]byte, error) {
	return []byte(c.Rank + " " + c.Suit), nil
}

func (c *Card) UnmarshalText(text []byte) error {
	split := strings.Split(string(text), " ")
	if len(split) < 2 {
		return fmt.Errorf("Invalid card: %v", text)
	}
	c.Rank = split[0]
	c.Suit = split[1]
	return nil
}

func (c *Card) String() string {
	return fmt.Sprintf("%s of %s", c.Rank, c.Suit)
}

var suits = []string{"Hearts", "Spades", "Clubs", "Diamonds"}
var ranks = []string{"Two", "Three", "Four", "Five", "Six", "Seven",
	"Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

func Cards() []card.CardFace {
	deck := make([]card.CardFace, len(suits)*len(ranks))
	for i, suit := range suits {
		for j, rank := range ranks {
			c := &Card{Suit: suit, Rank: rank}
			text, _ := c.MarshalText()
			deck[i*len(ranks)+j] = card.CardFace(text)
		}
	}
	return deck
}

type Cribbage struct {
	deck      Deck
	opponent  Opponent
	players   int
	playerNum int
}

type Deck interface {
	Draw(uint) (card.CardFace, error)
	Quit()
	Start([]card.CardFace) error
}

type Opponent interface {
	Send(message []byte) error
	Recieve() <-chan []byte
}

func NewCribbage(deck Deck, opp Opponent, playerNum int) (*Cribbage, error) {
	if deck == nil {
		return nil, fmt.Errorf("Cannot create Cribbage with nil deck")
	} else if playerNum < 1 || playerNum > 2 {
		return nil, fmt.Errorf("Illegal playerNum %d", playerNum)
	}
	return &Cribbage{
		deck:      deck,
		players:   2,
		playerNum: playerNum,
		opponent:  opp,
	}, nil
}

func (c *Cribbage) Hand() ([]*Card, error) {
	handSize := getHandSize(c.players)
	hand := make([]*Card, handSize)
	var index uint
	for i := range hand {
		if c.playerNum == 1 {
			index = 2 * uint(i)
		} else if c.playerNum == 2 {
			index = 2*uint(i) + 1
		} else {
			return nil, fmt.Errorf("Unsupported player number %d", c.playerNum)
		}

		current, err := c.deck.Draw(index)
		if err != nil {
			return nil, errors.Wrapf(err, "Unable to get hand")
		}
		hand[i] = &Card{}
		hand[i].UnmarshalText(current)
	}

	return hand, nil
}

func (c *Cribbage) Quit() error {
	c.deck.Quit()
	return nil
}

func getHandSize(numPlayers int) int {
	switch numPlayers {
	case 2:
		return 6
	case 3:
		fallthrough
	case 4:
		return 5
	default:
		return 0
	}
}
