package lib

import (
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"sort"
	"strconv"

	"github.com/schollz/progressbar/v3"
)

type Color struct {
	Name       string
	RGBA       color.RGBA
	Percentage float64
	Points     [][2]int
}

func rgbToString(u uint32) string {
	return strconv.Itoa(int(uint8(u)))
}

func GetColorBreakDown(img image.Image) []Color {
	bounds := img.Bounds()
	area := bounds.Dx() * bounds.Dy()

	m := make(map[color.Color][][2]int)

	bar := progressbar.NewOptions(
		area,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetDescription("Detecting colors"),
	)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			bar.Add(1)
			rgba := img.At(x, y)
			// r_str, g_str, b_str, a_str :=
			// 	rgbToString(r),
			// 	rgbToString(g),
			// 	rgbToString(b),
			// 	rgbToString(a)

			coord := [2]int{x, y}

			m[rgba] = append(m[rgba], coord)
		}
	}

	keys := make([]color.Color, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	// try to reduce number of loops by sorting and creating summary in same loop?

	sort.Slice(keys, func(i, j int) bool {
		return len(m[keys[i]]) > len(m[keys[j]])
	})

	var summary []Color

	for _, v := range keys {
		r, g, b, a := v.RGBA()
		r_8, g_8, b_8, a_8 := uint8(r), uint8(g), uint8(b), uint8(a)
		r_s, g_s, b_s, a_s := strconv.Itoa(int(r_8)), strconv.Itoa(int(g_8)), strconv.Itoa(int(b_8)), strconv.Itoa(int(a_8))

		summary = append(summary,
			Color{
				Name:       fmt.Sprintf("(%s, %s, %s, %s)", r_s, g_s, b_s, a_s),
				RGBA:       color.RGBA{r_8, g_8, b_8, a_8},
				Percentage: float64(len(m[v])) / float64(area) * 100.0,
				Points:     m[v],
			},
		)
	}

	return summary
}
