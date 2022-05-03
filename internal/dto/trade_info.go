package dto

// Текущая статистика (выполнено и открыто ордеров)
type TradeInfo struct {
	ID               uint64
	ExternalID       uint64
	SellFilledOrders string
	SellOpenOrders   string
}
