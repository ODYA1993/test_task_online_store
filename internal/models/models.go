package models

type Item struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Shelf struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Order struct {
	OrderID         int    `db:"order_id"`
	ID              int    `db:"id"`
	ItemID          int    `db:"item_id"`
	Quantity        int    `db:"quantity"`
	MainShelfID     int    `db:"main_shelf_id"`
	AdditionalShelf string `db:"additional_shelf"`
	Item            *Item  `db:"-"`
	MainShelf       *Shelf `db:"-"`
}

//// Product модель для товаров
//type Product struct {
//	ProductID   int    `db:"product_id"`
//	Name        string `db:"name"`
//	Description string `db:"description"`
//}
//
//// Order модель для заказов
//type Order struct {
//	OrderID     int `db:"order_id"`
//	OrderNumber int `db:"order_number"`
//}
//
//// Shelf модель для стеллажей
//type Shelf struct {
//	ShelfID int    `db:"shelf_id"`
//	Name    string `db:"name"`
//	IsMain  bool   `db:"is_main"`
//}
//
//// OrderProduct модель для товаров в заказах
//type OrderProduct struct {
//	OrderProductID int     `db:"order_product_id"`
//	OrderID        int     `db:"order_id"`
//	ProductID      int     `db:"product_id"`
//	Quantity       int     `db:"quantity"`
//	ShelfID        int     `db:"shelf_id"`
//	Product        Product `db:"product_id"`
//	Order          Order   `db:"order_id"`
//	Shelf          Shelf   `db:"shelf_id"`
//}
