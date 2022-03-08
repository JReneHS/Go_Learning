package main

import "fmt"

//inicializacion a nivel de paquetes
var aa int = 22       //full declaration
var is, js int = 1, 2 //tipos int is = 1 y js = 2

//Declaracion de variables en conjunto
var (
	id    int    = 2016630187
	name  string = "Juan Rene"
	lname string = "Hernandez Sanchez"
)

func main() {

	var a int      //declaracion
	a = 12         //inicializacion
	var b int = 12 //declaracion e inicializacion
	c := 33        //es tipo auto como C++
	var fl32 float32 = 3.1416
	var fl64 float64 = 3.3333
	var str string = "hola bb 7u7"

	var x1, x2, x3 int = 10, 20, 30

	fmt.Printf("%v, %v, %v\n", x1, x2, x3)

	var cs, python, java = true, false, "no!" //c = true | python = false | java = no!
	fmt.Println(is, js, cs, python, java)

	//se puede tener variables a nivel pakago
	//y en el scope con el mismo nombre
	fmt.Println("aa = ", aa)
	var aa int = 11
	fmt.Printf("var aa tipo  %T , valor = %v \n", aa, aa)
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("fl32 = ", fl32)
	fmt.Println("fl64 = ", fl64)
	fmt.Println("str = ", str)
	fmt.Println(id, name, lname)

	var p1 int = 36
	fmt.Printf("%v, %T \n", p1, p1)

	var p2 float32
	p2 = float32(p1) //usar el constructor para castear un tipo a otro
	fmt.Printf("%v, %T \n", p2, p2)

	var p3 string
	p3 = string(p1)
	fmt.Printf("%v, %T \n", p3, p3)

	//tipos primitivos

	var n bool = false
	var nn bool
	n1 := 1 == 1
	n2 := 2 == 1
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", n1, n1)
	fmt.Printf("%v, %T\n", n2, n2)
	fmt.Printf("%v, %T\n", nn, nn)

	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807
	fmt.Printf("%v, %T\n", i8, i8)
	fmt.Printf("%v, %T\n", i16, i16)
	fmt.Printf("%v, %T\n", i32, i32)
	fmt.Printf("%v, %T\n", i64, i64)

	var ui8 uint8 = 255
	var ui16 uint16 = 65535
	var ui32 uint32 = 1<<32 - 1
	var ui64 uint64 = 1<<64 - 1
	fmt.Printf("%v, %T\n", ui8, ui8)
	fmt.Printf("%v, %T\n", ui16, ui16)
	fmt.Printf("%v, %T\n", ui32, ui32)
	fmt.Printf("%v, %T\n", ui64, ui64)

	var com64 complex64 = 236 + 2i
	var com128 complex128 = 45 + 0.54i
	fmt.Printf("%v, %T\n", com64, com64)
	fmt.Printf("%v, %T\n", com128, com128)

}
