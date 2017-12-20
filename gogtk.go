package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var outTE *walk.TextEdit
	MainWindow{
		Title:   "eztool",
		MinSize: Size{400, 400},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					PushButton{
						Text:     "启动Nginx",
						OnClicked: func() {
						},
					},
					PushButton{
						Text:      "启动redis",
						OnClicked: func() { },
					},
					PushButton{
						Text:      "一键安装composer",
						OnClicked: func() { },
					},
					PushButton{
						Text:      "一键安装Golang",
						OnClicked: func() { },
					},
				},
			},
		},
	}.Run()
}