package main

import (
	"fmt"
	"github.com/lauri-lyytikainen/go-ascii-art-generator/internal/loading"
	"github.com/lauri-lyytikainen/go-ascii-art-generator/internal/parameters"
)

func main() {
	config, configErr := parameters.ParseParameters()
	if configErr != nil {
		fmt.Printf("ERROR: Loading config failed: %s\n", configErr)
	}

	_, err := loading.LoadImage(config.InputFile)
	if err != nil {
		fmt.Printf("ERROR: Loading image failed: %s\n", err)
	}
}
