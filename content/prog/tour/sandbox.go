package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Println("¡Bienvenida al playground!")

	fmt.Println("Hora:", time.Now())

	fmt.Println("Si quieres abrir un fichero:")
	fmt.Println(os.Open("filename"))

	fmt.Println("O acceder a la red:")
	fmt.Println(net.Dial("tcp", "google.com"))
}
