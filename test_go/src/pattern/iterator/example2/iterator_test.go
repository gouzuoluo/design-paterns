package example2

import (
	"testing"
	"fmt"
)

func TestAll(t *testing.T) {
	var pancakeHouseMenu *PancakeHouseMenu = NewPancakeHouseMenu()//早餐菜单
	var dinerMenu *DinerMenu = NewDinerMenu()//午餐菜单
	var cafeMenu *CafeMenu = NewCafeMenu()//晚餐菜单

	// Without iterators
	//printMenu();

	//女服务员
	var waitress *Waitress = NewWaitress(pancakeHouseMenu, dinerMenu, cafeMenu)
	waitress.PrintMenu()
}

func printMenu() {
	var pancakeHouseMenu *PancakeHouseMenu = NewPancakeHouseMenu()//早餐菜单
	var dinerMenu *DinerMenu = NewDinerMenu()//午餐菜单

	breakfastItems := pancakeHouseMenu.GetMenuItems()
	lunchItems := dinerMenu.GetMenuItems()

	fmt.Println("USING FOR EACH")
	for e := breakfastItems.Front(); e != nil; e = e.Next() {
		menuItem := e.Value.(*MenuItem)
		fmt.Print(menuItem.GetName() + ",")
		fmt.Print(fmt.Sprintf("%f", menuItem.GetPrice()) + ",")
		fmt.Println(menuItem.GetDescription() + ",")
	}

	fmt.Println("USING FOR LOOPS")
	for _, menuItem := range lunchItems {
		fmt.Print(menuItem.GetName() + ",")
		fmt.Print(fmt.Sprintf("%f", menuItem.GetPrice()) + ",")
		fmt.Println(menuItem.GetDescription() + ",")
	}
}