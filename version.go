package kevingopkg

import "fmt"

const VERSION = 1.1

const name = "kevin gopkg"

func printName() {
	fmt.Println(name)
}

func PrintName() {
	fmt.Println(name)
}

func init() {
	fmt.Println("kevingopkg init")
}
