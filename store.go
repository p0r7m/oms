package main

type IStore interface {
	AddCustomer(string)
	AddPartner(string)
}


type Store struct {
	customers     []*Customer
	partner       []*Partner
	customerOrder []CustomerOrder
	IStore
}

func CreateStore() *Store {
	return &Store{}
}

func (s *Store) AddCustomer(name string) *Customer {
	c := CreateCustomer(name)
	s.customers = append(s.customers, c)
	return c
}

func (s *Store) AddPartner(name string) *Partner {
	p := CreatePartner(name)
	s.partner = append(s.partner, p)
	return p
}

func (s *Store) PlaceOrder(order *Order, cust *Customer) {
	custOrder := CustomerOrder{customerId: cust.id, order: order}
	s.customerOrder = append(s.customerOrder, custOrder)
	cust.PlaceOrder(order, s.partner)
}