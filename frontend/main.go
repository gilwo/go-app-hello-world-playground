package main

import (
	"goappex"
	_ "goappex/frontcode"
)

func main() {

	if goappex.Mainfront == nil {
		panic("front code not found")
	}
	goappex.Mainfront()
}
