package main

import (
	"simple_go_proj/internal/pkg/log"
	"fmt"
)

func main() {
	a := log.GetSomeVar()
	fmt.Print(a)
}