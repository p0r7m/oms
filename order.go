package main

type OrderItems struct {
	itemId int
	qty    int
}

type IOrder interface {
	GetItems()
}

type Order struct {
	orderId int
	items   []*OrderItems
	status  string
	IOrder
}

func CreateOrder() *Order {
	return &Order{}
}

func (o *Order) AddItem(itemId int, qty int) {
	it := &OrderItems{itemId: itemId, qty: qty}
	o.items = append(o.items, it)
}