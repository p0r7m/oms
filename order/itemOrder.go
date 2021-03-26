package order

type OrderItems struct {
	itemId int
	qty    int
}

func (oi OrderItems) GetItemId() int {
	return oi.itemId
}

func (oi OrderItems) GetQty() int {
	return oi.qty
}
