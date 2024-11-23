package main

import (
	"fmt"
	"go-basics/calcurator"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	fmt.Println(os.Getenv("GO_ENV"))
	fmt.Println(calcurator.Sum(5, 5))
	fmt.Println(calcurator.Multiply(3, 2))
	i := 1
	ui := uint16(2)
	fmt.Printf("i: %v %T\n", i, i)
	fmt.Printf("i: %[1]v %[1]T ui: %[2]v %[2]T \n", i, ui)

	pi, title := 3.14, "Go"
	fmt.Printf("pi: %v title: %v", pi, title)
}
