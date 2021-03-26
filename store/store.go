package store

import (
	"oms/customer"
	"oms/order"
	partner2 "oms/partner"
)

type Store struct {
	customers     []*customer.Customer
	partner       []*partner2.Partner
	customerOrder []CustomerOrder
	IStore
}

func CreateStore() *Store {
	return &Store{}
}

func (s *Store) AddCustomer(name string) *customer.Customer {
	c := customer.CreateCustomer(name)
	s.customers = append(s.customers, c)
	return c
}

func (s *Store) AddPartner(name string) *partner2.Partner {
	p := partner2.CreatePartner(name)
	s.partner = append(s.partner, p)
	return p
}

func (s *Store) PlaceOrder(order *order.Order, cust *customer.Customer) {
	custOrder := CreateCustomerOrder(cust.GetId(), order)
	s.customerOrder = append(s.customerOrder, custOrder)
	cust.PlaceOrder(order, s.partner)
}