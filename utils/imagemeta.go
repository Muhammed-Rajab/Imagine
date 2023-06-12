package utils

import (
	"fmt"
	"mime/multipart"

	_ "image/jpeg"
	_ "image/png"

	"github.com/disintegration/imaging"
)

// Metadata related stuff
type ImageMeta struct {
	Valid      bool   `json:"is_valid"`
	Format     string `json:"format"`
	Resolution [2]int `json:"resolution"`
}

func (*ImageUtils) GetMetaData(file *multipart.File) *ImageMeta {

	meta := new(ImageMeta)
	meta.Valid = false
	meta.Format = ""
	meta.Resolution = [2]int{0, 0}

	// Check for validity
	img, err := imaging.Decode(*file)

	if err != nil {
		fmt.Println("Image isn't valid!")
		return meta
	}
	meta.Valid = true

	// Get resolution
	W := img.Bounds().Dx()
	H := img.Bounds().Dy()
	meta.Resolution = [2]int{W, H}

	return meta
}
