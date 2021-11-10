package main

import (
	"fmt"
	"github.com/kevin021788/kevingopkg"  // 使用自定义包
)

func main() {
	sum := kevingopkg.Add(10, 2)
	fmt.Println(sum)
}