package main

import (
	"github.com/lauri-lyytikainen/go-ascii-art-generator/internal/convert"
	"github.com/lauri-lyytikainen/go-ascii-art-generator/internal/parameters"
)

func main() {
	config, err := parameters.ParseParameters()
	if err != nil {
		panic("Loading config failed")
	}

	convert.Convert(config)
}
