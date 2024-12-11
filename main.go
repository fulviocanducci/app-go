package main

import (
	"fmt"

	"github.com.br/app-go/internal/calculate"
)

func main() {
	result := calculate.Sum(10, 20)
	fmt.Println(result)
}
