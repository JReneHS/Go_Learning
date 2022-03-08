package main

import (
	"bufio" //maneja un buffer de lectura y escritura
	"fmt"
	"os"
	"strconv"
)

func main() {
	edad := "22"
	edad_int, _ := strconv.Atoi(edad) //el opetrador "_" es para inorar

	fmt.Println(edad_int + 10)

	fmt.Print("Esta funcion imprime en la misma linea\n")
	fmt.Printf("Esta usa formatos de impresion de tipos y datos\n")
	fmt.Println("Esta cosa es igual a Print pero agrega salto de linea al final")
	fmt.Println("Ingresa tu Edad: ")
	fmt.Scanf("%s\n", &edad)
	fmt.Println("Mi edad es: " + edad)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingresa tu nombre completo:")
	nombre, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Holaaa " + nombre)
	}
}
