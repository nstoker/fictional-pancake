package main

import (
	"fmt"

	"github.com/nstoker/fictional-pancake/internal/version"
)

func main() {
	fmt.Printf("Hello, world, from %s\n", version.Version())
}
