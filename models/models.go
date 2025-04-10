package models

type Order struct {
	Id        uint64 `db:"id" json:"id"`
	Status    string `db:"status" json:"status"`
	Quantity  uint32 `db:"quantity" json:"quantity"`
	ProductID uint64 `db:"product_id" json:"product_id"`
}
