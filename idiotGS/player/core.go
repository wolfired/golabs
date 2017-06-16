package player

// Core 核心
type Core struct {
	uuid [16]byte //唯一ID
	name string   //名字
	lv   byte     //等级
}

// Encode 编码
func (c *Core) Encode(raw []byte) {
}

// Decode 解码
func (c *Core) Decode(raw []byte) {
}
