package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Ejercicio2() {
	scanner := bufio.NewScanner(os.Stdin)
	var num int
	var err error
	for {
		fmt.Print("Ingrese numero : ")
		if scanner.Scan() {
			num, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Error al convertir numero: " + err.Error())
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", num, i, num*i)
	}
}
