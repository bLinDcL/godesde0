package ejercicios

import "strconv"

func Ejercicio(input string) (num int, output string) {
	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, "Error al convertir numero" + err.Error()
	}

	if num > 100 {
		return num, "Es mayor a 100"
	} else {
		return num, "Es menor a 100"
	}
}
