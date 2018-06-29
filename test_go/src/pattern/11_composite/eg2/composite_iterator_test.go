package eg2

import (
	"testing"
)

func TestAll(t *testing.T) {

	//总菜单
	var allMenus MenuComponent = NewMenu("ALL MENUS", "All menus combined")

	//1.煎饼屋菜单（早餐菜单）
	var pancakeHouseMenu MenuComponent = NewMenu("PANCAKE HOUSE MENU", "Breakfast")
	pancakeHouseMenu.Add(NewMenuItem(
		"K&B's Pancake Breakfast",
		"Pancakes with scrambled eggs, and toast",
		true,
		2.99))
	pancakeHouseMenu.Add(NewMenuItem(
		"Regular Pancake Breakfast",
		"Pancakes with fried eggs, sausage",
		false,
		2.99))
	pancakeHouseMenu.Add(NewMenuItem(
		"Blueberry Pancakes",
		"Pancakes made with fresh blueberries, and blueberry syrup",
		true,
		3.49))
	pancakeHouseMenu.Add(NewMenuItem(
		"Waffles",
		"Waffles, with your choice of blueberries or strawberries",
		true,
		3.59))
	allMenus.Add(pancakeHouseMenu)


	//2.晚餐菜单
	var dinerMenu MenuComponent = NewMenu("DINER MENU", "Lunch")
	dinerMenu.Add(NewMenuItem(
		"Vegetarian BLT",
		"(Fakin') Bacon with lettuce & tomato on whole wheat",
		true,
		2.99))
	dinerMenu.Add(NewMenuItem(
		"BLT",
		"Bacon with lettuce & tomato on whole wheat",
		false,
		2.99))
	dinerMenu.Add(NewMenuItem(
		"Soup of the day",
		"A bowl of the soup of the day, with a side of potato salad",
		false,
		3.29))
	dinerMenu.Add(NewMenuItem(
		"Hotdog",
		"A hot dog, with saurkraut, relish, onions, topped with cheese",
		false,
		3.05))
	dinerMenu.Add(NewMenuItem(
		"Steamed Veggies and Brown Rice",
		"A medly of steamed vegetables over brown rice",
		true,
		3.99))

	dinerMenu.Add(NewMenuItem(
		"Pasta",
		"Spaghetti with Marinara Sauce, and a slice of sourdough bread",
		true,
		3.89))
	allMenus.Add(dinerMenu)

	//2.1甜点菜单
	var dessertMenu MenuComponent = NewMenu("DESSERT MENU", "Dessert of course!")
	dessertMenu.Add(NewMenuItem(
		"Apple Pie",
		"Apple pie with a flakey crust, topped with vanilla icecream",
		true,
		1.59))
	dessertMenu.Add(NewMenuItem(
		"Cheesecake",
		"Creamy New York cheesecake, with a chocolate graham crust",
		true,
		1.99))
	dessertMenu.Add(NewMenuItem(
		"Sorbet",
		"A scoop of raspberry and a scoop of lime",
		true,
		1.89))
	dinerMenu.Add(dessertMenu)


	NewWaitress(allMenus).PrintVegetarianMenu()
}
