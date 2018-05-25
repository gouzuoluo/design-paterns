package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

//做一个巨丑无比的登录框
//这里用到了LineEdit、LineEdit控件
func F2() {
	var usernameTE, passwordTE *walk.LineEdit
	MainWindow{
		Title:   "登录",
		MinSize: Size{270, 290},
		Layout:  VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2, Spacing: 10},
				Children: []Widget{
					VSplitter{
						Children: []Widget{
							Label{
								Text: "用户名",
							},
						},
					},
					VSplitter{
						Children: []Widget{
							LineEdit{
								MinSize:  Size{160, 0},
								AssignTo: &usernameTE,
							},
						},
					},
					VSplitter{
						Children: []Widget{
							Label{MaxSize: Size{160, 40},
								Text: "密码",
							},
						},
					},
					VSplitter{
						Children: []Widget{
							LineEdit{
								MinSize:  Size{160, 0},
								AssignTo: &passwordTE,
							},
						},
					},
				},
			},

			PushButton{
				Text:    "登录",
				MinSize: Size{120, 50},
				OnClicked: func() {
					if usernameTE.Text() == "" {
						var tmp walk.Form
						walk.MsgBox(tmp, "用户名为空", "", walk.MsgBoxIconInformation)
						return
					}
					if passwordTE.Text() == "" {
						var tmp walk.Form
						walk.MsgBox(tmp, "密码为空", "", walk.MsgBoxIconInformation)
						return
					}
				},
			},
		},
	}.Run()
}
