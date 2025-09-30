package main

import (
	"ServerOper/functions"
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	var socket net.Conn
	fmt.Println("*******************************************")
	fmt.Println("||       Cliente Operativos 2023.2       ||")
	fmt.Println("*******************************************")

	//Si los parametros brindados al ejecutar el programa en la terminal son menos de 4, termina el programa
	if len(os.Args) < 4 {
		fmt.Println("Parametros insuficientes")
		os.Exit(1)
	}

	// Obtener los argumentos
	ip := os.Args[1]
	puertoStr := os.Args[2]
	tiempoInfo := os.Args[3]

	//Juntar ip y puerto
	ip_puerto := ip + ":" + puertoStr
	//Crear la direccion tcp para conectarse al servidor
	tcpApdress, err := net.ResolveTCPAddr("tcp4", ip_puerto)
	if err != nil {
		fmt.Println("Error al resolver la dirección:", err)
		return
	}
	//Abrir un socket en la direccion obtenida
	socket, err = net.DialTCP("tcp", nil, tcpApdress)
	if err != nil {
		fmt.Println("Error al conectar en el puerto:", err)
		return
	}
	//Se crea un escritor en el socket
	env := bufio.NewWriter(socket)
	//Se envia el parametro recibido del tiempo entre reportes
	env.WriteString(tiempoInfo + "\n")
	env.Flush()

	//Se lanza el menu de cliente
	functions.Menu(&socket)
	socket.Close()
}
