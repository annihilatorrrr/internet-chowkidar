package main

import (
	"fmt"

	"github.com/gnulinuxindia/internet-chowkidar/di"
)

func main() {
	fmt.Println("Hello, World!")
	di.InjectDb()
}
