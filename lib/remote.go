package lib

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
)

func GetRemoteImage(url string) (image.Image, error) {
	request, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer request.Body.Close()

	img, _, err := image.Decode(request.Body)

	if err != nil {
		return nil, err
	}

	return img, nil
}
