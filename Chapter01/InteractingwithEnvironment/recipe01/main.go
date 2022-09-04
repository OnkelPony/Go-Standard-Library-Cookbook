package main

import (
	"log"
	"runtime"
)

const info = `
Application %s starting.
The binary was build by GO: %s
It is running on %s processor with %d cores`

func main() {
	log.Printf(info, "Example", runtime.Version(), runtime.GOARCH, runtime.NumCPU())
}
