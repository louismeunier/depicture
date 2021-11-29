package lib

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"sort"
	"strconv"

	"github.com/schollz/progressbar/v3"
)

type Color struct {
	Name       string
	Percentage float64
	Points     [][2]int
}

func rgbToString(u uint32) string {
	return strconv.Itoa(int(uint8(u)))
}

func GetColorBreakDown(img image.Image) []Color {
	bounds := img.Bounds()
	area := bounds.Dx() * bounds.Dy()

	m := make(map[string][][2]int)

	bar := progressbar.NewOptions(
		area,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetDescription("Detecting colors"),
	)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			bar.Add(1)
			r, g, b, a := img.At(x, y).RGBA()
			r_str, g_str, b_str, a_str :=
				rgbToString(r),
				rgbToString(g),
				rgbToString(b),
				rgbToString(a)

			coord := [2]int{x, y}

			m[r_str+","+g_str+","+b_str+","+a_str] = append(m[r_str+","+g_str+","+b_str+","+a_str], coord)
		}
	}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	// try to reduce number of loops by sorting and creating summary in same loop?

	sort.Slice(keys, func(i, j int) bool {
		return len(m[keys[i]]) > len(m[keys[j]])
	})

	var summary []Color

	for _, v := range keys {
		summary = append(summary,
			Color{
				Name:       v,
				Percentage: float64(len(m[v])) / float64(area) * 100.0,
				Points:     m[v],
			},
		)
	}

	return summary
}
