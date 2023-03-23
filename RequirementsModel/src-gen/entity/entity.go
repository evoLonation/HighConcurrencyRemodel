package entity
import "time"

type Item struct {
	Barcode int64 `db:"barcode"`
	Name string `db:"name"`
	Price float64 `db:"price"`
	StockNumber int64 `db:"stock_number"`
	Description string `db:"description"`
	BelongedShopShopId int64 `db:"belonged_shop_shop_id"`
}
type User struct {
	UserId string `db:"user_id"`
	Username string `db:"username"`
}
type Sale struct {
	GenerateId int64 `db:"generate_id"`
	Number int64 `db:"number"`
	Price float64 `db:"price"`
	ItemBarcode int64 `db:"item_barcode"`
	BelongedOrderOrderId int64 `db:"belonged_order_order_id"`
}
type Order struct {
	OrderId int64 `db:"order_id"`
	TotalPrice float64 `db:"total_price"`
	IsPayed bool `db:"is_payed"`
	BelongedUserUserId string `db:"belonged_user_user_id"`
}
type Shop struct {
	ShopId int64 `db:"shop_id"`
	ShopName string `db:"shop_name"`
}
