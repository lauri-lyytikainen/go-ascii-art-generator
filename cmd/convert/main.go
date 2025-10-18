package main

import (
	"github.com/lauri-lyytikainen/go-ascii-art-generator/internal/convert"
	"github.com/lauri-lyytikainen/go-ascii-art-generator/internal/parameters"
	"log"
)

func main() {
	config, err := parameters.ParseParameters()
	if err != nil {
		log.Panic(err)
	}

	convert.Convert(config)
}
