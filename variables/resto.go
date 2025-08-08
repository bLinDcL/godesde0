package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Felipe"
	Estado = true
	Sueldo = 2992.10
	Fecha = time.Now()

	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha.Date())
}

func ConvertToText(numero int) (bool, string) {
	text := strconv.Itoa(numero)

	return true, text
}
