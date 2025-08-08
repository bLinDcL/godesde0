package iteraciones

import (
	"fmt"
)

func Iterar() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}
