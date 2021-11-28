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

func rgbToString(u uint32) string {
	return strconv.Itoa(int(uint8(u)))
}

func GetColorBreakDown(img image.Image) (map[string][][2]int, []string) {
	bounds := img.Bounds()
	m := make(map[string][][2]int)

	bar := progressbar.NewOptions(
		bounds.Dx()*bounds.Dy(),
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

	sort.Slice(keys, func(i, j int) bool {
		return len(m[keys[i]]) > len(m[keys[j]])
	})

	return m, keys
}
