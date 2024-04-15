package main

import (
	"fmt"

	"github.com/FedorKowarnow/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	fmt.Println("Hello", st)
}
