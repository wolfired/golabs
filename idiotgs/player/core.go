package player

import (
	"github.com/wolfired/golabs/idiotgs/player/slot"
)

// Core 核心
type Core struct {
	uuid   [16]byte                  //唯一ID
	name   string                    //名字
	race   byte                      //种族
	gender byte                      //性别
	slots  map[slot.Slot]interface{} //槽位
}

// Encode 编码
func (c *Core) Encode(raw []byte) {
}

// Decode 解码
func (c *Core) Decode(raw []byte) {
}
