package order

type IOrder interface {
	GetItems() []*OrderItems
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

func (o *Order) GetItems() []*OrderItems {
	return o.items
}

func (o *Order) GetStatus() string {
	return o.status
}

func (o *Order) SetStatus(status string) {
	o.status = status
}