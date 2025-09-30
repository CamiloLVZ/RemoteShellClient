package functions

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"net"
	"os"
	"strings"
	"time"
)

func Menu(socket *net.Conn) {
	for {
		fmt.Println("MENU PRINCIPAL")
		fmt.Println("[1] Crear un usuario")
		fmt.Println("[2] Iniciar Sesion")
		fmt.Println("[0] Finalizar")

		fmt.Println("\nIngrese una opcion")
		opcion, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		opcion = strings.TrimSpace(opcion)

		switch opcion {
		case "0":
			fmt.Println("Saliendo del programa")
			return
		case "1":
			fmt.Println("----CREAR UN USUARIO----")
			fmt.Println("Ingrese nombre de usuario:")
			//Leer usuario
			user, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				fmt.Println("Error al leer usuario:", err)
				continue
			}
			//Crear escritor en el socket
			env := bufio.NewWriter(*socket)
			//Enviar usuario recibido
			env.WriteString(user)
			err = env.Flush()
			if err != nil {
				fmt.Println("Error, Servidor desconectado")
				return
			}
			fmt.Println("Ingrese Password:")

			//pedir contraseña de forma oculta
			password := PedirPassword() + "\n"

			//Enviar Contraseña
			env.WriteString(password)
			env.Flush()
			time.Sleep(200 * time.Millisecond)
			//Enviar opcion del usuario
			env.WriteString("registrar\n")
			err = env.Flush()
			if err != nil {
				fmt.Println("Error, Servidor desconectado")
				return
			}
			//Recibir respuesta del server
			respuesta, err := bufio.NewReader(*socket).ReadString('\n')
			fmt.Println(respuesta)

		case "2":
			fmt.Println("----LOGIN----")
			fmt.Println("Ingrese Usuario:")
			//Recibir user
			user, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				fmt.Println("Error al leer usuario:", err)
				continue
			}
			//Crear escritor y enviar usuario al server
			env := bufio.NewWriter(*socket)
			env.WriteString(user)

			err = env.Flush()
			if err != nil {
				fmt.Println("Error, Servidor desconectado")
				return
			}
			fmt.Println("Ingrese Password:")
			//Pedir contraseña en modo oculto
			password := PedirPassword() + "\n"

			//Enviar contraseña y opcion
			env.WriteString(password)
			env.Flush()
			time.Sleep(200 * time.Millisecond)
			env.WriteString("login\n")
			err = env.Flush()
			if err != nil {
				fmt.Println("Error, Servidor desconectado")
				return
			}
			loginExitoso := false
			//Recibir respuesta
			respuesta, err := bufio.NewReader(*socket).ReadString('\n')
			if respuesta == "succes\n" {
				//Si el login fue exitoso en el server
				loginExitoso = true
			} else if respuesta == "failed\n" {
				//Si el login fue fallido se repite el envio de contraseña las veces necesarias
				for {
					mensaje, err := bufio.NewReader(*socket).ReadString('\n')
					if err != nil {
						fmt.Println("Error, Servidor desconectado")
						return
					}
					if mensaje == "succes\n" {
						loginExitoso = true
						break
					} else if mensaje == "failed\n" {
						break
					}
					fmt.Println(mensaje)
					password = PedirPassword() + "\n"
					//Envio de contraseña
					env.WriteString(password)
					err = env.Flush()
					if err != nil {
						fmt.Println("Error, Servidor desconectado")
						return
					}
				}
			} else if respuesta == "notUser\n" {
				//Si el server dice que el usuario no existe
				fmt.Println("El usuario ingresado no existe")
			}
			if loginExitoso {
				//SI el login es exitoso, se da acceso a la shell y a enviar comandos
				fmt.Println("Login Exitoso!\nBienvenido a la Shell Remota")
				//Se lanza un hilo en segundo plano para recibir reportes
				go RecibeReporte(socket)
				//Se da acceso a enviar comando al servidor
				EnviarComandos(socket)
			}
		default:
			fmt.Println("Opción no válida. Intente de nuevo.")
		}
	}
}

// FUNCION QUE PIDE CONTRASEÑA SIN MOSTRAR LO QUE SE ESCRIBE EN LA TERMINAL
func PedirPassword() string {
	password, _ := terminal.ReadPassword(int(os.Stdin.Fd()))
	return string(password)
}
