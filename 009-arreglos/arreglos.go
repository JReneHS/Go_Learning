package main

import (
	"fmt"
)

func main() {
	//esta cosa es estatica y se inicializa a un valor predeterminado
	var array [10]int         //en este caso el varor predeterminado es 0
	var array2 [10]string     //con Strings es vacio
	array3 := [5]int{1, 2, 3} //los que no se inicializan son 0
	fmt.Println(array)
	fmt.Println(array2)
	fmt.Println(array3)

	fmt.Println("La longitud de array1 es:", len(array))

	for i := 0; i < len(array); i++ {
		fmt.Println("'", array[i], "'")
	}

	var matriz [3][2]int
	s := 1
	fmt.Println(matriz)
	//Relleno de la matriz
	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[0]); j++ {
			matriz[i][j] = s
			s++
		}
	}

	fmt.Println(matriz)
	//impresion de la matriz
	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[0]); j++ {
			fmt.Printf("%d, ", matriz[i][j])
		}
		fmt.Println()
	}
}
