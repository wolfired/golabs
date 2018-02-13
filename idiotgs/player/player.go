package player

import (
	"github.com/satori/go.uuid"

	"github.com/wolfired/golabs/idiotgs/player/slot"
)

// Player 玩家
type Player struct {
	Core
}

// MakePlayer 创建一个Player
func MakePlayer() *Player {
	p := new(Player)
	p.uuid = uuid.NewV4()
	p.slots = make(map[slot.Slot]interface{}, 1)

	c := slot.Currency{Balance: 0}
	p.slots[c.Slot] = &c

	return p
}
