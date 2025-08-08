package main

import (
	"fmt"

	"github.com/bLinDcL/godesde0/variables"
)

func main() {
	estado, texto := variables.ConvertToText(1993)

	fmt.Println(estado)
	fmt.Println(texto)
}
