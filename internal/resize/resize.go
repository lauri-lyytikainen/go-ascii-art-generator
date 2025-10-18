package resize

import (
	"image"
	"log"
)

func Resize(image image.Image, width int16, height int16) {

}

func FitImageToTerminal(imageWidth int, imageHeight int, terminalWidth int, terminalHeight int) (width int, height int) {
	const fontAspectRatio float32 = 2.0
	var scaleX = float32(terminalWidth) / float32(imageWidth)
	var scaleY = float32(terminalHeight) / float32(imageHeight)

	var scale = min(scaleX, scaleY)

	log.Println(scaleX)
	log.Println(scaleY)
	// TODO: Make the aspectratioed image full height again
	return int(float32(imageWidth) * scale), int(float32(imageHeight) * scale / fontAspectRatio)
}
