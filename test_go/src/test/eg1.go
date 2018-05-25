//使用第三方库lxn/walk，进行Windows GUI编程。
//go get github.com/lxn/walk

package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"strings"
)

func F1() {
	var inTE, outTE *walk.TextEdit

	MainWindow{
		Title:   "SCREAMO",
		MinSize: Size{600, 400},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			PushButton{
				Text: "SCREAM",
				OnClicked: func() {
					outTE.SetText(strings.ToUpper(inTE.Text()))
				},
			},
		},
	}.Run()
}

//go build生成go_windows_gui.exe
/*
go_windows_gui.exe.manifest:
	<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
	<assembly xmlns="urn:schemas-microsoft-com:asm.v1" manifestVersion="1.0">
	<assemblyIdentity version="1.0.0.0" processorArchitecture="*" name="SomeFunkyNameHere" type="win32"/>
	<dependency>
	<dependentAssembly>
	<assemblyIdentity type="win32" name="Microsoft.Windows.Common-Controls" version="6.0.0.0" processorArchitecture="*" publicKeyToken="6595b64144ccf1df" language="*"/>
	</dependentAssembly>
	</dependency>
	</assembly>

manifest是一种xml文件，标明所依赖的side-by-side组建。
如果用VS开发，可以Set通过porperty->configuration properties->linker->manifest file->Generate manifest To Yes来自
动创建manifest来指定系统的和CRT的assembly版本。

详解
观察上面的manifest文件：

<xml>这是xml声明：
版本号----<?xml version="1.0"?>。
这是必选项。 尽管以后的 XML 版本可能会更改该数字，但是 1.0 是当前的版本。

编码声明------<?xml version="1.0" encoding="UTF-8"?>
这是可选项。 如果使用编码声明，必须紧接在 XML 声明的版本信息之后，并且必须包含代表现有字符编码的值。

standalone表示该xml是不是独立的，如果是yes，则表示这个XML文档时独立的，不能引用外部的DTD规范文件；
如果是no，则该XML文档不是独立的，表示可以用外部的DTD规范文档。

<dependency>这一部分指明了其依赖于一个库：
<assemblyIdentity>属性里面还分别是：
type----系统类型，
version----版本号，
processorArchitecture----平台环境，
publicKeyToken----公匙

*/
