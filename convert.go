package main

import (
	"errors"
	"image"
	"image/color"
	"log"
)

func CreatePixelArrayFromImage(img image.Image) ([][]color.Color, error) {
	var pixels [][]color.Color

	imgSize := img.Bounds().Size()
	for xIndex := 0; xIndex < imgSize.X; xIndex++ {
		var yArray []color.Color
		for yIndex := 0; yIndex < imgSize.Y; yIndex++ {
			yArray = append(yArray, img.At(xIndex, yIndex))
		}
		pixels = append(pixels, yArray)
	}

	return pixels, nil

}

func CreateImageFromPixelArray(pixels [][]color.Color) (image.Image, error) {

	if len(pixels) == 0 {
		return nil, errors.New("pixel conversion on empty array is invalid")
	}

	imgRect := image.Rect(0, 0, len(pixels), len(pixels[0]))
	newImage := image.NewRGBA(imgRect)

	for xIndex := 0; xIndex < imgRect.Dx(); xIndex++ {
		pixelCol := pixels[xIndex]
		for yIndex := 0; yIndex < imgRect.Dy(); yIndex++ {
			singlePixel := pixelCol[yIndex]
			pixelRGBA, ok := color.RGBAModel.Convert(singlePixel).(color.RGBA)
			if ok {
				newImage.Set(xIndex, yIndex, pixelRGBA)
			} else {
				log.Printf("Could not convert pixel value: %v %v %v", xIndex, yIndex, singlePixel)
				return nil, errors.New("pixel conversion error")
			}
		}
	}

	return newImage, nil
}
