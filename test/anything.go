package main

import "fmt"

func main() {
	fmt.Println("Hello GoLang")
	saitama()
	fmt.Println("Boss of Genos")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

}

func saitama() {
	fmt.Println("I'm Saitama")
}
