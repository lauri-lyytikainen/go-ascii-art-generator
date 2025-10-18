package convert

import (
	"fmt"
	"github.com/lauri-lyytikainen/go-ascii-art-generator/internal/loading"
	"github.com/lauri-lyytikainen/go-ascii-art-generator/internal/parameters"
)

func Convert(config *parameters.Config) {
	fmt.Println(config)
	m, err := loading.LoadImage(config.InputFile)
	if err != nil {
		fmt.Println(err)
		panic("Loading image failed")
	}
	fmt.Println(m.Bounds())
}
