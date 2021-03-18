package main

import (
	"fmt"

	"github.com/joubertredrat-tests/ebanx-dev-test-golang-2k21/pkg"
)

func main() {
	fmt.Println("Server running at http://0.0.0.0:8000")
	pkg.Run()
}
