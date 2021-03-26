package customer

import (
	"fmt"
	"oms/order"
	"oms/partner"
)

type PartnerMenu struct {
	partnerId int
	itemId    int
	itemName  string
	price     float64
}

type Customer struct {
	id   int
	name string
	ICustomer
}

func CreateCustomer(name string) *Customer {
	return &Customer{name: name}
}

func (c *Customer) PlaceOrder(order *order.Order, partner []*partner.Partner) {
	partnerMenu := c.getAllRestaurantItemsByPartnerId(partner)
	processOrder(order, partnerMenu)
}

func (c *Customer) GetId() int {
	return c.id
}

func  (c *Customer) getAllRestaurantItemsByPartnerId(partner []*partner.Partner) map[int][]PartnerMenu {
	pm := make(map[int][]PartnerMenu)

	for _, eachPartner := range partner {
		menuItems := eachPartner.GetMenuWithItems()
		if !eachPartner.IsEligible() || len(menuItems) == 0{
			continue
		}
		for _, eachItem := range menuItems {
			t := PartnerMenu{
				partnerId: eachPartner.GetId(),
				itemId: eachItem.GetItemId(),
				itemName: eachItem.GetName(),
				price: eachItem.GetPrice(),
			}
			pm[eachItem.GetItemId()] = append(pm[eachItem.GetItemId()], t)
		}
	}
	return pm
}

func processOrder(order *order.Order, partnerMenu map[int][]PartnerMenu) {
	var servingPartners []PartnerMenu
	// get all eligible partner list
	for _, item := range order.GetItems() {
		if _, ok := partnerMenu[item.GetItemId()]; ok {
			servingPartners = append(servingPartners, partnerMenu[item.GetItemId()]...)
		} else {
			order.SetStatus("PROCESSING_FAILED")
			return
		}
	}
	order.SetStatus("PROCESSED")
	fmt.Println(servingPartners)
}