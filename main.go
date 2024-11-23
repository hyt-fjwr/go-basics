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
	fmt.Printf("i: %v %T\n", i, i)
}
