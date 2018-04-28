package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
)
func scan_demo() {
	var apple_str, orange_str string
	fmt.Scan(&apple_str, &orange_str)
	fmt.Println(apple_str)
	fmt.Println(orange_str)
}
func scanf_demo() {
	var apple_count int
	var apple_name string
	var apple_price float64
	fmt.Scanf("%d %s %f", &apple_count, &apple_name, &apple_price)
	fmt.Println(apple_count)
	fmt.Println(apple_name)
	fmt.Println(apple_price)
}
func scanln_demo() {
	var app_line string
	fmt.Scanln(&app_line)
	fmt.Println(app_line)
}
func main() {
	//scan_demo()
	//scanf_demo()
	//scanln_demo()

	cmdReader := bufio.NewReader(os.Stdin)
	cmdStr, _ := cmdReader.ReadString('\n')
	//这里把读取的数据后面的换行去掉，对于Mac是"\r"，Linux下面
	//是"\n"，Windows下面是"\r\n"，所以为了支持多平台，直接用
	//"\r\n"作为过滤字符
	cmdStr = strings.Trim(cmdStr, "\r\n")
	fmt.Println(cmdStr)
}