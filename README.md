# html-snapshot

一个可以快照网页的Go模块，将HTML转换成图片或PDF

A Go module that can snapshot web pages and convert HTML to images or PDFs

## 前提条件 Prerequisite

安装好NodeJS和NPM并配置好环境变量。必要时可能需要可用的上网工具

Install NodeJS and NPM and configure environment variables. Available Internet tools may be required when necessary

## 用法 Usage

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

	c.Src = "https://www.sina.com.cn"
	c.Dst = "sample3.pdf"

	err = c.MakePdfPage()
	if err != nil {
		println(err.Error())
	}
}

```
