package example1

import (
	"container/list"
)

//菜单接口
type Menu interface {
	CreateIterator() Iterator
}

/*-------------------------------------------------------------------------------------------------------------------*/

//1.煎饼屋的菜单（早餐菜单）
type PancakeHouseMenu struct {
	menuItems *list.List
}

func NewPancakeHouseMenu() *PancakeHouseMenu {
	this := new(PancakeHouseMenu)
	this.menuItems = list.New()
	this.init()
	return this
}

func (this *PancakeHouseMenu)init() {
	this.addItem("K&B's Pancake Breakfast",
		"Pancakes with scrambled eggs, and toast",
		true,
		2.99)

	this.addItem("Regular Pancake Breakfast",
		"Pancakes with fried eggs, sausage",
		false,
		2.99)

	this.addItem("Blueberry Pancakes",
		"Pancakes made with fresh blueberries",
		true,
		3.49)

	this.addItem("Waffles",
		"Waffles, with your choice of blueberries or strawberries",
		true,
		3.59)
}
//添加菜单项
func (this *PancakeHouseMenu)addItem(name, description string, vegetarian bool, price float64) {
	menuItem := NewMenuItem(name, description, vegetarian, price)
	this.menuItems.PushBack(menuItem)
}

//获取菜单列表
func (this *PancakeHouseMenu)GetMenuItems() *list.List {
	return this.menuItems
}


//实现Menu接口的createIterator方法
func (this *PancakeHouseMenu)CreateIterator() Iterator {
	return NewPancakeHouseMenuIterator(this.menuItems)
}

//
func (this *PancakeHouseMenu)String() string {
	return "Objective Pancake House Menu"
}

// other menu methods here


//2.餐厅的菜单（午餐菜单）
type DinerMenu struct {
	menuItems []*MenuItem
}

func NewDinerMenu() *DinerMenu {
	this := new(DinerMenu)
	this.menuItems = make([]*MenuItem, 0, 10)
	this.init()
	return this
}

func (this *DinerMenu)init() {
	this.addItem("Vegetarian BLT",
		"(Fakin') Bacon with lettuce & tomato on whole wheat", true, 2.99)
	this.addItem("BLT",
		"Bacon with lettuce & tomato on whole wheat", false, 2.99)
	this.addItem("Soup of the day",
		"Soup of the day, with a side of potato salad", false, 3.29)
	this.addItem("Hotdog",
		"A hot dog, with saurkraut, relish, onions, topped with cheese",
		false, 3.05)
	this.addItem("Steamed Veggies and Brown Rice",
		"Steamed vegetables over brown rice", true, 3.99)
	this.addItem("Pasta",
		"Spaghetti with Marinara Sauce, and a slice of sourdough bread",
		true, 3.89)
}

func (this *DinerMenu)addItem(name, description string, vegetarian bool, price float64) {
	menuItem := NewMenuItem(name, description, vegetarian, price)
	this.menuItems = append(this.menuItems, menuItem)
}

func (this *DinerMenu)GetMenuItems() []*MenuItem {
	return this.menuItems
}

//实现Menu接口的CreateIterator方法
func (this *DinerMenu)CreateIterator()Iterator  {
	return NewDinerMenuIterator(this.menuItems)
	// To test Alternating menu items, comment out above line,
	// and uncomment the line below.
	//return new AlternatingDinerMenuIterator(menuItems)
}

// other menu methods here