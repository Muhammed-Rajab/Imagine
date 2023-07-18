package utils

import (
	"bytes"
	"errors"
	"image"
	"image/color"
	"math"

	_ "image/jpeg"

	_ "image/png"

	"github.com/disintegration/imaging"
)

type ImageUtils struct{}
type FilterFunc func(img *image.Image) *image.NRGBA

// Function which transforms image acc. to query
func (*ImageUtils) TransformImage(img *image.Image, query *TransformImageQuery) ([]byte, error) {

	// Decode config of the image and get the format
	// For now, just use JPEG
	format := imaging.JPEG

	// Converting img to imaging.Image to get better features
	// TODO: try refactoring this and you might end up with better results
	imagingImg := imaging.Clone(*img)

	// Getting all the necessary queries and Parsing them
	parsedQuery := query.Parse()

	// Rotate image
	if math.Mod(parsedQuery.Rotate, 360) != 0 {
		imagingImg = imaging.Rotate(imagingImg, parsedQuery.Rotate, color.Transparent)
	}

	// Scale image
	if parsedQuery.Scale != 1 {
		bounds := imagingImg.Bounds()
		w := bounds.Dx()
		h := bounds.Dy()
		imagingImg = imaging.Resize(imagingImg, w*int(parsedQuery.Scale), h*int(parsedQuery.Scale), imaging.NearestNeighbor)
	}

	// Resize image
	if parsedQuery.Resize[0] != -1 && parsedQuery.Resize[1] != -1 {
		imagingImg = imaging.Resize(imagingImg, parsedQuery.Resize[0], parsedQuery.Resize[1], imaging.NearestNeighbor)
	}

	// Encode the image back to image.Image
	var buf bytes.Buffer

	err := imaging.Encode(&buf, imagingImg, format)

	if err != nil {
		return []byte{}, errors.New("error occured while encoding transformed image, please try again later")
	}

	return buf.Bytes(), nil
}

// Function which applies filter
func (*ImageUtils) ApplyCustomFilter(img *image.Image, fn FilterFunc) ([]byte, error) {
	filteredImage := fn(img)

	var buf bytes.Buffer

	if err := imaging.Encode(&buf, filteredImage, imaging.JPEG); err != nil {
		return []byte{}, errors.New("error occured while encoding transformed image, please try again later")
	}

	return buf.Bytes(), nil
}

// Filters
func (*ImageUtils) Negative(img *image.Image) *image.NRGBA {

	output := imaging.Clone(*img)

	width := output.Bounds().Size().X
	height := output.Bounds().Size().Y

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Get the color of the pixel at (x, y)
			pixel := output.At(x, y)
			r, g, b, a := pixel.RGBA()

			// Modify the color values as per your custom filter logic
			// For example, you can invert the colors by subtracting each component from 255
			invertedR := 255 - uint8(r>>8)
			invertedG := 255 - uint8(g>>8)
			invertedB := 255 - uint8(b>>8)

			// Create a new color using the modified values
			invertedColor := color.RGBA{invertedR, invertedG, invertedB, uint8(a >> 8)}

			// Set the modified color to the pixel at (x, y)
			output.Set(x, y, invertedColor)
		}
	}

	return output
}

func (*ImageUtils) Grayscale(img *image.Image) *image.NRGBA {
	return imaging.Grayscale(*img)
}

func (*ImageUtils) Sepia(img *image.Image) *image.NRGBA {
	// Implement sepia here
	// return imaging.Grayscale(*img)

	bounds := (*img).Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	// Create a new NRGBA image for the modified image
	sepiaImg := image.NewNRGBA(bounds)

	// Iterate over each pixel in the original image
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			originalColor := (*img).At(x, y)
			rgbaColor := color.RGBAModel.Convert(originalColor).(color.RGBA)

			// Calculate sepia values
			sepiaR := uint8((float64(rgbaColor.R)*0.393 + float64(rgbaColor.G)*0.769 + float64(rgbaColor.B)*0.189))
			sepiaG := uint8((float64(rgbaColor.R)*0.349 + float64(rgbaColor.G)*0.686 + float64(rgbaColor.B)*0.168))
			sepiaB := uint8((float64(rgbaColor.R)*0.272 + float64(rgbaColor.G)*0.534 + float64(rgbaColor.B)*0.131))

			// Set the sepia color to the new image
			sepiaImg.Set(x, y, color.RGBA{R: sepiaR, G: sepiaG, B: sepiaB, A: rgbaColor.A})
		}
	}

	return sepiaImg
}
