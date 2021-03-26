package partner

type Menu struct {
	name   string
	menuId int
	items  []MenuItem
}

func (m *Menu) GetMenuItems() []MenuItem {
	return m.items
}