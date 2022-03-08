package main

import "fmt"

func duplicar(pntr *int) {
	*pntr = *pntr * 2
}

func main() {
	var w, x, y, z int
	x = 10
	y = 20
	z = 30

	/// hay tres formas de inicializar punteros
	//Con sintaxis tipo C/C++
	var p1 *int
	p1 = &x
	//Con constructor
	var p2 = new(int)
	p2 = &y
	//con auto formato
	p3 := &z
	p4 := &w
	*p4 = 8
	fmt.Printf("x = %v -> p1 => %v (%T) DIR(%v)|(%v)\n", x, *p1, p1, p1, &x)
	fmt.Printf("y = %v -> p1 => %v (%T) DIR(%v)|(%v)\n", y, *p2, p2, p2, &y)
	fmt.Printf("z = %v -> p1 => %v (%T) DIR(%v)|(%v)\n", z, *p3, p3, p3, &z)
	fmt.Printf("w = %2v -> p1 => %2v (%T) DIR(%v)|(%v)\n", w, *p4, p4, p4, &w)

}
