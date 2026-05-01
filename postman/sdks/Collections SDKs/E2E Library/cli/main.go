package main

import (
	_ "example.com/e2e-library/cmd/books"
	_ "example.com/e2e-library/config"
	"example.com/e2e-library/root"
	_ "example.com/e2e-library/setupauth"
)

func main() {
	root.Execute()
}
