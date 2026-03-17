package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	project := os.Getenv("GOLANG_PROJECT")
	println(project)
}
