package main

import (
	"go-serverless-demo/internal/api"
	"log"
)

func main() {
	log.Fatal(api.New().Start(":8080"))
}
