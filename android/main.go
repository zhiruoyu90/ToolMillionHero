package main

import (
	"fmt"
	"../baidu/ocr"
)

func main() {
	ocr.GetImageText("screenshot_1.png")
	fmt.Println("ok")
}
