package main //declaração de pacote; todo programa deve começar com ela.

import (
	"fmt"
)

//USAR UMA VARIÁVEL FORA DE UMA FUNÇÃO PERMITE QUE OUTRAS FUNÇÕES ACESSE ESSA VARIÁVEL

func main() {

	fmt.Println("Hello, World!")

	fmt.Println("Hello, my name is Carollina.")

	fmt.Println("1 + 1 =", 1+1)

	fmt.Println("1 + 1 =", 1.0+1.0)

	fmt.Println(len("Hello, World"))

	fmt.Println("Hello World"[1])

	fmt.Println("Hello, " + "World")

	fmt.Println(true && true)

	fmt.Println(true && false)

	fmt.Println(true || true)

	fmt.Println(true || false)

	fmt.Println(!true)

	fmt.Println("32.123 x 42.452 =", 32.123*42.452)

	fmt.Println((true && false) || (false && true) || !(false && false))

	var w string = "Hello, World"
	fmt.Println(w)

	var x string
	x = "Hello, World!"
	fmt.Println(x)

	var y string
	y = "first"
	fmt.Println(y)
	y = "second"
	fmt.Println(y)

	var z string
	z = "first "
	fmt.Println(z)
	z = z + "second"
	fmt.Println(z)

	var a string = "Hello"
	var b string = "World"
	fmt.Println(a == b)

	var c string = "Hello"
	var d string = "Hello"
	fmt.Println(c == d)

	e := "Hello, World!"
	fmt.Println(e)

	dogsName := "Ted and Thor"
	fmt.Println("My dog's names are,", dogsName)

	const f string = "Hello, World!"
	fmt.Println(f)
	//CONSTANTES SÃO VARIÁVEIS CUJOS VALORES NÃO PODEM SER ALTERADOS DEPOIS.

	/*var (
		a = 25
		b = 24
		c = 29
	)*/
	//DEFINIR DIVERSAS VARIÁVEIS; USAR "VAR" OU "CONST", SEGUIDA DE PARÊNTESES, COM CADA VARIÁVEL EM SUA PRÓPRIA LINHA.

	fmt.Println("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)

	fmt.Println("Coloque o Fahrenheit: ")
	var fahrenheit float64
	fmt.Scanf("%f", &fahrenheit)
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Println(celsius)

	fmt.Println("Insira o pés: ")
	var pes float64
	fmt.Scanf("%f", &pes)
	metros := (pes * 0.3048)
	fmt.Println(metros)
}
