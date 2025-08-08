package main

import (
	"fmt"
	"runtime"

	"github.com/bLinDcL/godesde0/ejercicios"
	"github.com/bLinDcL/godesde0/variables"
)

func main() {
	estado, texto := variables.ConvertToText(1993)
	fmt.Println(estado)
	fmt.Println(texto)

	if os := runtime.GOOS; os == "darwin" || os == "Linux." {
		fmt.Println("OS X")
	} else {
		fmt.Println("Windows")
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "Linux":
		fmt.Println("Linux")
	default:
		fmt.Println("Windows")
	}

	num, output := ejercicios.Ejercicio("99")
	fmt.Println(num)
	fmt.Println(output)
}
