package convert

import (
	"fmt"
	tsize "github.com/kopoli/go-terminal-size"
	"github.com/lauri-lyytikainen/go-ascii-art-generator/internal/draw"
	"github.com/lauri-lyytikainen/go-ascii-art-generator/internal/loading"
	"github.com/lauri-lyytikainen/go-ascii-art-generator/internal/parameters"
	"github.com/lauri-lyytikainen/go-ascii-art-generator/internal/resize"
	"log"
)

func Convert(config *parameters.Config) {
	fmt.Println(config)
	m, err := loading.LoadImage(config.InputFile)
	if err != nil {
		log.Panic(err)
	}

	s, err := tsize.GetSize()
	if err != nil {
		log.Panic(err)
	}

	terminalImageWidth, terminalImageHeight := resize.FitImageToTerminal(m.Bounds().Max.X, m.Bounds().Max.Y, s.Width, s.Height)

	log.Println(m.Bounds())
	log.Println(s)
	log.Printf("Terminal image width %d height %d\n", terminalImageWidth, terminalImageHeight)

	draw.Draw(terminalImageWidth, terminalImageHeight)
}
