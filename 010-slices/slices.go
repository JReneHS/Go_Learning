package main

import "fmt"

func main() {
	str := "1234567890"
	var matriz []int //slice vacio y es nulo
	//el tamanio del arreglo es parte del tipo de dato
	matr := []int{1, 2, 3, 4}

	if matriz == nil {
		fmt.Println("esta vacio")
	}
	fmt.Println("long of matriz = ", len(matr))
	//la estructructura de un slice tiene
	// -> longitud del arreglo
	// -> puntero al arreglo
	// -> la capacidad
	fmt.Println(str)
	sliceOfStr := str[1:3]
	fmt.Println(sliceOfStr)

	//slice con Make
	slice := make([]int, 3) //inicializa y aloca en mem tipos:
	// -> Slice
	// -> Map
	// -> Chan
	slice2 := make([]int, 5, 10) //slice con 5 elementos y 10 de capacidad.
	fmt.Println("Slice -> ", slice, " con cap = ", cap(slice))
	//Como los slice tienen una capacidad fija se utiliza la funcion append
	//El cual te agrega un valor pero tambien aumenta la capacidad
	slice2 = append(slice2, 2)
	fmt.Println("Slice -> ", slice2, " con cap = ", cap(slice2))
}
