package cribbage

const (
	STR_HELP = `Commands:
	hand - show current hand
	crib - show contents of crib
	toCrib <hand-index> - send the specified card to the crib
	cut - show the shared cut card
	cutAt - Choose the shared cut card for a hand
	seq - show the current sequence of played cards
	play <hand-index> - play the card at that index in your hand
	pass - choose not to play a card
	help - show this help message
	`
	STR_DRAW_INSTRUCTIONS          = "Type 'hand' to see your hand"
	STR_DISCARD_INSTRUCTIONS       = "Use 'toCrib' to send two cards to the crib"
	STR_DISCARD_WAIT_INSTRUCTIONS  = "Waiting for opponent to send two cards to the crib...\nCheck the crib periodically by typing 'crib'"
	STR_CUT_INSTRUCTIONS           = "Please cut the deck by typing 'cutAt <number>' to cut at the nth card."
	STR_CUT_WAIT_INSTRUCTIONS      = "Please wait for your opponent to cut the deck."
	STR_CIRCULAR_INSTRUCTIONS      = "Play a card with 'play <card-number' or 'pass' if you cannot."
	STR_CIRCULAR_WAIT_INSTRUCTIONS = "Wait for your opponent to play a card or pass."
	STR_INTERNAL_INSTRUCTIONS      = "Write these instructions please!"
	STR_CRIB_INSTRUCTIONS          = "Write these instructions please!"
	STR_END_INSTRUCTIONS           = "Write these instructions please!"
)
