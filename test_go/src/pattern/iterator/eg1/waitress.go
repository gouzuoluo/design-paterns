package eg1

import "fmt"

//女服务员
type Waitress struct {
	pancakeHouseMenu Menu //煎饼屋菜单
	dinerMenu        Menu //餐厅菜单
}

func NewWaitress(pancakeHouseMenu, dinerMenu Menu) *Waitress {
	this := new(Waitress)
	this.pancakeHouseMenu = pancakeHouseMenu
	this.dinerMenu = dinerMenu
	return this
}

//打印菜单
func (this *Waitress)PrintMenu() {
	var pancakeIterator Iterator = this.pancakeHouseMenu.CreateIterator()
	var dinerIterator Iterator = this.dinerMenu.CreateIterator()
	fmt.Println("MENU\n----\nBREAKFAST")
	this.printMenu(pancakeIterator)
	fmt.Println("\nLUNCH")
	this.printMenu(dinerIterator)
}

func (this *Waitress)printMenu(iterator Iterator) {
	for iterator.HasNext() {
		var menuItem *MenuItem = iterator.Next()
		fmt.Print(menuItem.GetName() + ",")
		fmt.Print(fmt.Sprintf("%f", menuItem.GetPrice()) + ",")
		fmt.Println(menuItem.GetDescription() + ",")
	}
}