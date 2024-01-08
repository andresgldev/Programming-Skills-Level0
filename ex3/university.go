package main

import "fmt"

type User struct {
	nombre   string
	apellido string
	campus   string
	programa string
}

type Programas struct {
	codigo string
	nombre string
	slots  int
	users  []User
}

type Campus struct {
	codigo    string
	nombre    string
	programas map[string]Programas
}

var programasDisponibles = map[string]string{
	"cps": "Computer Science",
	"med": "Medicine",
	"mkt": "Marketing",
	"art": "Arts",
}

var campus = map[string]Campus{
	"lon": {
		codigo: "lon",
		nombre: "London",
		programas: map[string]Programas{
			"cps": {codigo: "cps", nombre: "Computer Science", slots: 1},
			"med": {codigo: "med", nombre: "Medicine", slots: 1},
			"mkt": {codigo: "mkt", nombre: "Marketing", slots: 1},
			"art": {codigo: "art", nombre: "Arts", slots: 1},
		},
	},
	"mch": {
		codigo: "mch",
		nombre: "Manchester",
		programas: map[string]Programas{
			"cps": {codigo: "cps", nombre: "Computer Science", slots: 3},
			"med": {codigo: "med", nombre: "Medicine", slots: 3},
			"mkt": {codigo: "mkt", nombre: "Marketing", slots: 3},
			"art": {codigo: "art", nombre: "Arts", slots: 3},
		},
	},
	"liv": {
		codigo: "liv",
		nombre: "Liverpool",
		programas: map[string]Programas{
			"cps": {codigo: "cps", nombre: "Computer Science", slots: 1},
			"med": {codigo: "med", nombre: "Medicine", slots: 1},
			"mkt": {codigo: "mkt", nombre: "Marketing", slots: 1},
			"art": {codigo: "art", nombre: "Arts", slots: 1},
		},
	},
}

func main() {
	welcome()

	user, login := login()
	if !login {
		return
	}

	fmt.Println("Bienvenido, ahora registre sus datos")
	nombre := getInput("Ingrese su Nombre: ")
	apellido := getInput("Ingrese su Apellido: ")

	user.nombre = nombre
	user.apellido = apellido

	for {
		fmt.Println()
		for i, name := range programasDisponibles {
			fmt.Printf("[%s] %s \n", i, name)
		}

		programa := getInput("Escoja el programa: ")

		_, existe := programasDisponibles[programa]
		if existe {
			user.programa = programa
			break
		} else {
			fmt.Println("Este programa no existe, escoja entre los disponibles")
		}
	}

	for {
		fmt.Println()
		for i, name := range campus {
			fmt.Printf("[%s] %s \n", i, name.nombre)
		}

		campusInput := getInput("Escoja el campus: ")

		campusEscogido, existe := campus[campusInput]
		if existe {
			carrera := campusEscogido.programas[user.programa]
			posicion := len(carrera.users)
			if carrera.slots >= posicion {
				user.campus = campusInput
				//carrera.users[posicion] = user
				fmt.Println("Usted fue enrolado en el campus ", campusEscogido.nombre, "en el programa ", carrera.nombre)
				break
			} else {
				fmt.Println("No DISPONIBLE")
			}

		} else {
			fmt.Println("Este programa no existe, escoja entre los disponibles")
		}
	}
}

func welcome() {
	fmt.Println("Bienvenido al sistema universitario")
}

func login() (User, bool) {
	var count int = 0
	for {
		login := getInput("Ingrese su usuario: ")
		pass := getInput("Ingrese su contraseña: ")
		fmt.Println(" ")

		if login == "user" && pass == "pass" {
			return User{}, true
		} else {
			count++
			fmt.Println("Credenciales Invalidas")
			if count >= 3 {
				fmt.Println("Su usuario a sido bloqueado")
				break
			}
		}
	}
	return User{}, false
}

func getInput(message string) string {
	fmt.Print(message)
	var txt string
	fmt.Scan(&txt)
	return txt
}
