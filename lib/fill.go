package lib

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"

	"github.com/anthonynsimon/bild/paint"
)

func Fill(oldImage image.Image, fillColor color.Color, colorName string) image.Image {
	new_image := paint.FloodFill(oldImage, image.Point{0, 0}, fillColor, 15)

	os.Mkdir("out", 0755)

	f, err := os.Create(fmt.Sprintf("./out/%s.jpeg", colorName))
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if err := jpeg.Encode(f, new_image, nil); err != nil {
		panic(err)
	}

	return new_image
}
