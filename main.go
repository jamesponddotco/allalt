package main

import (
	"os"

	"git.sr.ht/~jamesponddotco/allalt/internal/app"
)

func main() {
	os.Exit(app.Run(os.Args))
}
