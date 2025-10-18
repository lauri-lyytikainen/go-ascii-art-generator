package loading

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func LoadImage(filePath string) (image.Image, error) {
	fmt.Printf("Attempting to load image from %s\n", filePath)

	reader, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	// bounds := m.Bounds()

	return m, nil
}
