package customer

import (
	"oms/order"
	"oms/partner"
)

type ICustomer interface {
	PlaceOrder(*order.Order, []*partner.Partner)
}
