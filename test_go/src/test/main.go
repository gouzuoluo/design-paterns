package main

import (
	"html/template"
	"log"
	"bytes"
	"fmt"
)

func main() {
	// 声明模板内容
	const tpl = ` 
        <font>哇哦,</font><font color=#008000 size=30>{{ .AccName }}</font><font>在</font><font color=#008000 size=30><gamename gameid = "{{ .GameId }}"  subgameid = "{{ .SubGameId }}" gamelevel = "{{ .GameLevel }}" /></font><font>赢得了</font><font color=#008000 size=30>{{ .Coin }}</font>
`

	// 创建一个新的模板，并且载入内容
	t, err := template.New("跑马灯").Parse(tpl)
	if err != nil {
		log.Fatal(err)
	}

	//定义传入到模板的数据，并在终端打印
	data := struct {
		AccName   string
		GameId    string
		SubGameId string
		GameLevel string
		Coin      string
		B         string
	}{
		AccName:   "张三",
		GameId:    "slots",
		SubGameId: "wangbao",
		Coin:      "500",
		B:         "2222",
	}
	//err = t.Execute(os.Stdout, data)
	//if err != nil {
	//	log.Fatal(err)
	//}

	buf := bytes.Buffer{}
	err = t.Execute(&buf, data)

	fmt.Println(buf.String())
}
