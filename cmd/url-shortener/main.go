package main

import (
	"fmt"

	"url-shortener/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// init logger: slog

	// init storage: postgresql

	// init router: chi, chi render

	// init server:
}
