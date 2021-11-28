package cmd

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/louismeunier/depicture/lib"
	"github.com/spf13/cobra"
)

var (
	maxColors int
)

var rootCmd = &cobra.Command{
	Use:   "depicture <image/path>",
	Short: "Extract the primary colors in an image",
	Long: `A CLI utility to extract a color scheme from a picture.
		See source code and more: https://github.com/louismeunier/depicture.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		imagePath := args[0]
		_, err := os.Stat(imagePath)
		if os.IsNotExist(err) {
			fmt.Printf("❌ Image \"%s\" is not found", imagePath)
			os.Exit(1)
		}

		imageReader, err := os.Open(imagePath)

		if err != nil {
			fmt.Println("❌ That file was found, but an error occurred opening it")
			os.Exit(1)
		}

		defer imageReader.Close()

		img, _, err := image.Decode(imageReader)

		if err != nil {
			fmt.Println("❌ Failed to decode that file")
			os.Exit(1)
		}

		breakdown, keys := lib.GetColorBreakDown(img)
		bounds := img.Bounds()

		for i := 0; i < maxColors; i++ {
			percent := float64(len(breakdown[keys[i]])) / float64(bounds.Max.X*bounds.Max.Y) * 100
			fmt.Printf("%f%% rgba(%s)\n", percent, keys[i])
		}
	},
}

func Execute() {
	rootCmd.PersistentFlags().IntVarP(&maxColors, "max-colors", "c", 3, "maximum colors to return")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
