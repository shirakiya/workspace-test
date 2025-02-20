package main

import (
	"fmt"

	"github.com/shirakiya/workspace-test/modulea"
	"github.com/shirakiya/workspace-test/moduleb"
)

func main() {
	a := modulea.Func1()
	b := moduleb.Func1()

	fmt.Println(a)
	fmt.Println(b)
}
