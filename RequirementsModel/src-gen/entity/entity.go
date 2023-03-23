package entity
import "time"

type Item struct {
	Barcode Int64 `db:"barcode"`
	Name String `db:"name"`
	Price Float64 `db:"price"`
	StockNumber Int64 `db:"stock_number"`
	Description String `db:"description"`
	BelongedShopShopId Int64 `db:"belonged_shop_shop_id"`
}
type User struct {
	UserId String `db:"user_id"`
	Username String `db:"username"`
}
type Sale struct {
	GenerateId Int64 `db:"generate_id"`
	Number Int64 `db:"number"`
	Price Float64 `db:"price"`
	ItemBarcode Int64 `db:"item_barcode"`
	BelongedOrderOrderId Int64 `db:"belonged_order_order_id"`
}
type Order struct {
	OrderId Int64 `db:"order_id"`
	TotalPrice Float64 `db:"total_price"`
	IsPayed Bool `db:"is_payed"`
	BelongedUserUserId String `db:"belonged_user_user_id"`
}
type Shop struct {
	ShopId Int64 `db:"shop_id"`
	ShopName String `db:"shop_name"`
}
