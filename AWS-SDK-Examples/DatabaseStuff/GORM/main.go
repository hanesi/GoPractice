package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	var p Product
	fmt.Println(p)
}
