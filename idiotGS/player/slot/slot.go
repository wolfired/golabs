package slot

// Slot 类型列表
const (
	SlotTypeGoodsInBag uint = 0x01 //背包中物品
	SlotTypeGoodsInBOX uint = 0x02 //仓库中物品
	SlotTypeGoodsInUse uint = 0x03 //使用中物品
	SlotTypeSkill      uint = 0x04 //技能
	SlotTypeBuff       uint = 0x05 //状态
	SlotTypeCurrency   uint = 0x06 //货币
)

// Slot 槽位
type Slot struct {
	//slot type 2^7
	//slot index 2^9
}
