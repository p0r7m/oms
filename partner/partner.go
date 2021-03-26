package partner

type Partner struct {
	id              int
	name            string
	capacity        int
	menu            *Menu
	IPartner
}

func CreatePartner(name string) *Partner {
	return &Partner{name: name}
}

func (p *Partner) GetId() int {
	return p.id
}

func (p *Partner) IsEligible() bool {
	return p.capacity != 0
}

func (p *Partner) GetMenuWithItems() []MenuItem {
	return p.menu.GetMenuItems()
}

func (p *Partner) AddMenuItem(id int, name string, price float64) {

	if p.menu == nil {
		p.menu = &Menu{}
	}
	item := MenuItem{itemId: id, itemName: name, price: price}
	p.menu.items = append(p.menu.items, item)
}

func (p *Partner) SetId(id int) {
	p.id = id
}

func (p *Partner) SetCapacity(cap int) {
	p.capacity = cap
}



