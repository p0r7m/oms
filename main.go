package main

import "fmt"

// 3:30
func main() {
	store := CreateStore()
	cust1 := store.AddCustomer("prm")
	p1 := store.AddPartner("PizzaHut")
	p2 := store.AddPartner("Dominos")
	p1.capacity = 10
	p2.capacity = 5
	p1.id = 21
	p2.id = 95
	p1.AddMenuItem(1, "pizaa", 393.21)
	p1.AddMenuItem(2, "coke", 23.21)
	p1.AddMenuItem(3, "cheesy dip", 33)
	p1.AddMenuItem(4, "garlic bread", 89)

	p2.AddMenuItem(1, "pizaa", 493.91)
	p2.AddMenuItem(2, "coke", 33.84)
	p2.AddMenuItem(3, "cheesy dip", 20)
	p2.AddMenuItem(4, "garlic bread", 99.13)

	order := CreateOrder()

	order.AddItem(1,1)
	order.AddItem(2,2)
	order.AddItem(4,1)

	store.PlaceOrder(order, cust1)

	fmt.Println(order.status)
}
