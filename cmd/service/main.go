package main

import (
	"context"
	"fmt"
	"log"

	"github.com/pavelshustrov/service_b/client/v1"
)

func main() {
	fmt.Println("service A started")
	client := v1.New()
	id, err := client.Create(context.Background())
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println(id.UUID)
}
