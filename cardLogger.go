package izapple2

import (
	"fmt"
)

/*
Logger card. It never existed, I use it to trace accesses to the card.
*/

// CardLogger is a fake card to log soft switch invocations
type CardLogger struct {
	cardBase
}

// NewCardLogger creates a new VidHD card
func NewCardLogger() *CardLogger {
	var c CardLogger
	c.name = "Softswitch log card"
	return &c
}

func (c *CardLogger) assign(a *Apple2, slot int) {
	c.addCardSoftSwitches(func(address uint8, data uint8, write bool) uint8 {
		if write {
			fmt.Printf("[cardLogger] Write access to softswith 0x%x for slot %v, value 0x%02x.\n", address, slot, data)
		} else {
			fmt.Printf("[cardLogger] Read access to softswith 0x%x for slot %v.\n", address, slot)
		}
		return 0
	}, "LOGGER")

	c.cardBase.assign(a, slot)
}
