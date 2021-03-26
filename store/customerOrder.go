package store

import order2 "oms/order"

type CustomerOrder struct {
	order      *order2.Order
	customerId int
}

func CreateCustomerOrder(id int, order *order2.Order) CustomerOrder {
	return CustomerOrder{customerId: id, order: order}
}