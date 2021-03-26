package partner

type MenuItem struct {
	itemName string
	price    float64
	itemId   int
}

func (mi MenuItem) GetName() string {
	return mi.itemName
}

func (mi MenuItem) GetItemId() int {
	return mi.itemId
}

func (mi MenuItem) GetPrice() float64 {
	return mi.price
}
