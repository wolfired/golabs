package slot

// Slot 类型列表
const (
	SlotTypeGoodsInBag uint16 = 0x01 //背包中物品
	SlotTypeGoodsInBox uint16 = 0x02 //仓库中物品
	SlotTypeGoodsInUse uint16 = 0x03 //使用中物品
	SlotTypeSkill      uint16 = 0x04 //技能
	SlotTypeBuff       uint16 = 0x05 //状态
	SlotTypeCurrency   uint16 = 0x06 //货币
)

// Slot 槽位
type Slot uint16 //type 2^7, index 2^9

// Type 类型
func (s *Slot) Type() uint16 {
	return uint16(*s) >> 9
}

// Index 索引
func (s *Slot) Index() uint16 {
	return uint16(*s) & 0x1FF
}
