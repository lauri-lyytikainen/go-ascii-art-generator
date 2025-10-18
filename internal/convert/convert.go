package convert

import (
	"fmt"
	"github.com/lauri-lyytikainen/go-ascii-art-generator/internal/loading"
	"github.com/lauri-lyytikainen/go-ascii-art-generator/internal/parameters"
	"log"
)

func Convert(config *parameters.Config) {
	fmt.Println(config)
	m, err := loading.LoadImage(config.InputFile)
	if err != nil {
		log.Panic(err)
	}
	log.Println(m.Bounds())
}
