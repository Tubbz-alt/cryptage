package deck

import (
	"io"
)

type Deck interface {
	Draw() (string, error)
	String() string
}

// ensure that *deck fulfills Deck interface
var _ Deck = &deck{}

type deck struct {
	cards      []card
	protocol   *Protocol
	connection io.ReadWriteCloser
}

// Draw draws a single card from the deck
func (d *deck) Draw() (string, error) {
	return "", nil
}

func (d *deck) String() string {
	return "Deck"
}

// NewDeck creates a deck of cards and assumes that the given
// io.ReadWriteCloser is a connection of some sort to another
// deck.
func NewDeck(deckConnection io.ReadWriteCloser) (Deck, error) {
	d := &deck{
		connection: deckConnection,
	}
	p, err := NewProtocol(deckConnection)
	if err != nil {
		return nil, err
	}
	d.protocol = p
	return d, nil
}
