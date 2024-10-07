package main

import (
	"fmt"
	"time"
)

type Task struct {
	ID          uint
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func main() {
	fmt.Println("Hello World")
}
