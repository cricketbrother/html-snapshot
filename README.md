# html-snapshot

一个可以快照网页的Go模块，将HTML转换成图片或PDF

## 前提条件

按照好NodeJS和NPM，并配置好环境变量。必要时，可能需要可用的上网工具

## 用法

```go
package main

import hs "github.com/cricketbrother/html-snapshot"

func main() {
	var c = &hs.Converter{
		NodePath: "node",
		Src:      "file:///absolute/path/to/sample.html",
		Dst:      "sample1.png",
		Width:    1920,
		Height:   1080,
	}

	err := c.MakeScreenshotPage()
	if err != nil {
		println(err.Error())
	}

	c.Element = "canvas"
	c.Dst = "sample2.png"

	err = c.MakeScreenshotElement()
	if err != nil {
		println(err.Error())
	}

	c.Dst = "sample3.pdf"

	err = c.MakePdfPage()
	if err != nil {
		println(err.Error())
	}
}

```
