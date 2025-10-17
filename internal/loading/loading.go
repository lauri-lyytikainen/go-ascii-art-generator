package loading

import (
	"errors"
	"fmt"
	"image"
)

func LoadImage(filePath string) (image.Image, error) {
	fmt.Printf("Attempting to load image from %s\n", filePath)
	return nil, errors.New("Not implemented")
}
