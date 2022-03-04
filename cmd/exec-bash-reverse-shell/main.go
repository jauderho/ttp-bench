package main

import (
	"log"

	"github.com/tstromberg/ioc-bench/pkg/simulate"
)

func main() {
	if err := simulate.BashReverseShell(); err != nil {
		log.Fatalf("unexpected error: %v", err)
	}
}