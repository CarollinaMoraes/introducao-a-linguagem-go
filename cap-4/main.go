package main

import "fmt"

func main() {

	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	/*for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%d even\n", i)
			// par
		} else {
			fmt.Printf("%d odd\n", i)
			//ímpar
		}
	}*/

	// if i % 2 == 0 {
	// 	// divisível por 2
	// 	else if i % 3 == 0 {
	// 		// divisível por 3
	// } else if i % 4 == 0 {
	// 	// divisível por 4
	// }
	// }

	// for i := 1; i <= 10; i++ {
	// 	if i %2 == 0 {
	// 		fmt.Println(i, "even")
	// 	} else {
	// 		fmt.Println(i, "odd")
	// 	}
	// 	}

	for i := 1; i <= 10; i++ {
		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("Ome")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		default:
			fmt.Println("Unknown Number")

		}

	}

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}

	for i := 1; i <= 100; i++ {
		 if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println(i, "Fizz")
		} else if i%5 == 0 {
			fmt.Println(i, "Buzz")
		} 
	}

}
