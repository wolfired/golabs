package slot

// CurrencyTypeExp 货币类型
const (
	CurrencyTypeExp     uint16 = 0x01 //经验
	CurrencyTypeCrystal uint16 = 0x02 //水晶
	CurrencyTypeGold    uint16 = 0x03 //黄金
	CurrencyTypeCoin    uint16 = 0x04 //金币
)

// Currency 货币
type Currency struct {
	Slot
	Balance uint32 //余额
}
