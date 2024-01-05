package main

import (
	"fmt"
)

type User struct {
	user       string
	pass       string
	amount     float32
	failLogins int
}

func (u *User) increaseFailedLoginAttempts() {
	u.failLogins++
}
func (u *User) ResetFailedLoginAttempts() {
	u.failLogins = 0
}
func (u *User) shouldBeBlock() bool {
	return u.failLogins >= 3
}
func (u *User) withdraw(value float32) {
	u.amount -= value
}
func (u *User) deposit(value float32) {
	u.amount += value
}
func (u *User) getAmount() float32 {
	return u.amount
}

func newUser(user string, pass string) *User {
	return &User{
		user:       user,
		pass:       pass,
		amount:     2000,
		failLogins: 0,
	}
}

func getData() map[string]*User {
	var users = map[string]*User{}
	users["user01"] = newUser("user01", "password1")
	users["user02"] = newUser("user02", "password2")
	return users
}

func main() {
	welcome()

	users := getData()
	user := doLogin(users)
	if user == nil {
		return
	}

	sayHello(user)

	for {
		showMenu()
		var c int = getChoice()
		switch c {
		case 1:
			fmt.Println("Tu saldo es", user.getAmount())
		case 2:
			fmt.Println("Usted a escogido DEPOSITO")
			user.deposit(getAmount("Ingrese el valor a depositar: "))
			fmt.Println("Tu nuevo saldo es", user.getAmount())
		case 3:

			fmt.Println("Usted a escogido RETIRO")
			user.withdraw(getAmount("Ingrese el valor a retirar: "))
			fmt.Println("Tu nuevo saldo es", user.getAmount())
		case 4:
			fmt.Println("Usted a escogido TRANSFERIR")
			var cuentaATransferir string = getInput("Ingrese la cuenta a transferir: ")
			usuarioATransferir, existeUsuarioATransferir := findUser(cuentaATransferir, users)
			if existeUsuarioATransferir {
				monto := getAmount("Ingrese el valor a transferir: ")
				if monto <= user.getAmount() {
					usuarioATransferir.deposit(monto)
					user.withdraw(monto)
					fmt.Println("\nTransferencia exitosa, tu nuevo saldo es", user.getAmount())
				} else {
					fmt.Println("\nFondos insuficientes")
				}
			} else {
				fmt.Println("\nNo se ha encontrado la cuenta")
			}
		default:
			fmt.Println("Adios!")
			return
		}
	}
}

func welcome() {
	fmt.Print("\n\nBienvenido a su banca en linea!\n\n")
}

func sayHello(user *User) {
	fmt.Print("\nBienvenido ", user.user, "\n")
}

func doLogin(users map[string]*User) *User {
	for {
		username := getInput("\nIngrese su usuario: ")
		password := getInput("Ingrese su password: ")

		user, exist := findUser(username, users)
		if exist {
			if password == user.pass {
				user.ResetFailedLoginAttempts()
				return user
			} else {
				fmt.Println("Credenciales Invalidas")

				user.increaseFailedLoginAttempts()
				if user.shouldBeBlock() {
					fmt.Println("Su usuario a sido bloqueado")
					return nil
				}

				continue
			}
		} else {
			fmt.Println("No existe este usuario")
			continue
		}
	}
}

func findUser(username string, users map[string]*User) (*User, bool) {
	user, exist := users[username]
	return user, exist
}

func showMenu() {
	fmt.Println("\nQue deseas hacer ahora?")
	fmt.Println("1. Ver el saldo")
	fmt.Println("2. Depositar")
	fmt.Println("3. Retirar")
	fmt.Println("4. Transferir")
	fmt.Print("5. Salir \n\n\n")
}

func getChoice() int {
	fmt.Print("Escoge una opcion (1-5): ")
	var choice int
	fmt.Scan(&choice)
	return choice
}

func getInput(message string) string {
	fmt.Print(message)
	var text string
	fmt.Scan(&text)
	return text
}

func getAmount(message string) float32 {
	fmt.Print(message)
	var amount float32
	fmt.Scan(&amount)
	return amount
}
