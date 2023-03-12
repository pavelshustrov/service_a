package main

import (
	"fmt"
	v1 "github.com/pavelshustrov/service_b/pkg/client/v1"
)

func main() {
	fmt.Println("service A started")
	v1.New()
}
