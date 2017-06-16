package player

import (
	"github.com/satori/go.uuid"
)

// Player 玩家
type Player struct {
	Core
}

// MakePlayer 创建一个Player
func MakePlayer() *Player {
	p := new(Player)
	p.uuid = uuid.NewV4()
	return p
}
