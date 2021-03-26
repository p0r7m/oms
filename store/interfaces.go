package store

import (
	"oms/customer"
	"oms/order"
	partner2 "oms/partner"
)

type IStore interface {
	AddCustomer(string) *customer.Customer
	AddPartner(string) *partner2.Partner
	PlaceOrder(*order.Order, *customer.Customer)
}
