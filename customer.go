package main

import "fmt"

type PartnerMenu struct {
	partnerId int
	itemId    int
	itemName  string
	price     float64
}

type ICustomer interface {
	Register()
	PlaceOrder(Order)
	GetAllPartners()
}

type Customer struct {
	id   int
	name string
	ICustomer
}

func CreateCustomer(name string) *Customer {
	return &Customer{name: name}
}

func (c *Customer) PlaceOrder(order *Order, partner []*Partner) {
	partnerMenu := c.getAllRestaurantItemsByPartnerId(partner)
	processOrder(order, partnerMenu)
}

func (c *Customer) GetId() int {
	return c.id
}

func  (c *Customer) getAllRestaurantItemsByPartnerId(partner []*Partner) map[int][]PartnerMenu {
	pm := make(map[int][]PartnerMenu)

	for _, eachPartner := range partner {
		if !eachPartner.IsEligible() || len(eachPartner.menu.items) == 0{
			continue
		}
		for _, eachItem := range eachPartner.menu.items {
			t := PartnerMenu{
				partnerId: eachPartner.id,
				itemId: eachItem.itemId,
				itemName: eachItem.itemName,
				price: eachItem.price,
			}
			pm[eachItem.itemId] = append(pm[eachItem.itemId], t)
		}
	}
	return pm
}

func processOrder(order *Order, partnerMenu map[int][]PartnerMenu) {
	var servingPartners []PartnerMenu
	// get all eligible partner list
	for _, item := range order.items {
		if _, ok := partnerMenu[item.itemId]; ok {
			servingPartners = append(servingPartners, partnerMenu[item.itemId]...)
		} else {
			order.status = "PROCESSING_FAILED"
			return
		}
	}
	order.status = "PROCESSED"
	fmt.Println(servingPartners)
}