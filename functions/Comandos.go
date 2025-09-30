package functions

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

// FUNCION QUE ENVIA CONSTANTEMENTE COMANDOS AL SERVIDOR
func EnviarComandos(socket *net.Conn) {
	//Crear lector del teclado
	lectorComando := bufio.NewReader(os.Stdin)
	//Crear escritor en el socket
	env := bufio.NewWriter(*socket)
	for {
		fmt.Print("$")
		//Leer comando del teclado
		comando, _ := lectorComando.ReadString('\n')
		// Enviar la longitud del comando antes de los datos
		var sizeComando uint16
		sizeComando = uint16(len(comando))
		err := binary.Write(*socket, binary.LittleEndian, sizeComando)
		if err != nil {
			fmt.Println("Error, Server desconectado")
			os.Exit(1)
		}

		// Enviar los datos del comando
		env.WriteString(comando)
		err = env.Flush()
		if err != nil {
			fmt.Println("Error, Servidor desconectado")
			os.Exit(1)
		}

		if comando == "bye\r\n" { //Si el comando que se escribe es "bye", acaba el programa
			fmt.Println("============================")
			fmt.Println("||    Gracias por usar    ||")
			fmt.Println("============================")
			os.Exit(0)
		}
	}
}

// FUNCION QUE RECIBE REPORTES DE RECURSOS Y SALIDAS DE COMANDOS CONSTANTEMENTE
func RecibeReporte(socket *net.Conn) {
	for {
		var mensajeLength uint16
		//Recibir longitud del mensaje
		err := binary.Read(*socket, binary.LittleEndian, &mensajeLength)
		if err != nil {
			return
		}
		//Crear un array de bytes
		mensajeBytes := make([]byte, mensajeLength)
		//Leer el mensaje en el array
		_, err = (*socket).Read(mensajeBytes)
		if err != nil {
			fmt.Println("Fin de conexion")
			return
		}
		//Convertir a string el mensaje
		mensaje := string(mensajeBytes)
		fmt.Println("Server:\n", mensaje)
	}
}
