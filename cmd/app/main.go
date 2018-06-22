package main

import (
	"simple_go_proj/internal/pkg/logwrapper"
)

func main() {
	wrapper := logwrapper.New()
	wrapper.L.Log()
}