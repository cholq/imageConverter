package main

import (
	"errors"
	"image/color"
)

type TransformRGBAValuesFn func(color.RGBA) (color.RGBA, error)

func calcGrayscaleValue(original color.RGBA) uint8 {
	minValue := min(original.R, original.G, original.B)
	maxValue := max(original.R, original.G, original.B)
	grayscale := (uint16(minValue) + uint16(maxValue)) / 2
	return uint8(grayscale)
}

func RGBAAverage(original []color.RGBA) (color.RGBA, error) {
	rTotalVal := 0
	gTotalVal := 0
	bTotalVal := 0
	var newColor color.RGBA

	for _, pixel := range original {
		rTotalVal += int(pixel.R)
		gTotalVal += int(pixel.G)
		bTotalVal += int(pixel.B)
	}

	if len(original) == 0 {
		return newColor, errors.New("cannot divide by zero")
	}
	newColor.R = uint8(rTotalVal / len(original))
	newColor.G = uint8(gTotalVal / len(original))
	newColor.B = uint8(bTotalVal / len(original))
	newColor.A = 255

	return newColor, nil
}

func RGBAGrayscale(original color.RGBA) (color.RGBA, error) {
	grayscale := calcGrayscaleValue(original)
	original.R = grayscale
	original.G = grayscale
	original.B = grayscale
	return original, nil
}

func RGBAGrayscaleAndBlue(original color.RGBA) (color.RGBA, error) {
	grayscale := calcGrayscaleValue(original)
	original.R = grayscale
	original.G = grayscale
	return original, nil
}

func RGBAGrayscaleAndGreen(original color.RGBA) (color.RGBA, error) {
	grayscale := calcGrayscaleValue(original)
	original.R = grayscale
	original.B = grayscale
	return original, nil
}

func RGBAGrayscaleAndRed(original color.RGBA) (color.RGBA, error) {
	grayscale := calcGrayscaleValue(original)
	original.G = grayscale
	original.B = grayscale
	return original, nil
}

func RGBAShiftLeft(original color.RGBA) (color.RGBA, error) {
	oldR := original.R
	oldG := original.G
	oldB := original.B
	original.R = oldG
	original.G = oldB
	original.B = oldR
	return original, nil
}

func RGBAShiftRight(original color.RGBA) (color.RGBA, error) {
	oldR := original.R
	oldG := original.G
	oldB := original.B
	original.R = oldB
	original.G = oldR
	original.B = oldG
	return original, nil
}

func RGBASwapGandB(original color.RGBA) (color.RGBA, error) {
	oldG := original.G
	oldB := original.B
	original.G = oldB
	original.B = oldG
	return original, nil
}

func RGBASwapRandB(original color.RGBA) (color.RGBA, error) {
	oldR := original.R
	oldB := original.B
	original.R = oldB
	original.B = oldR
	return original, nil
}

func RGBASwapRandG(original color.RGBA) (color.RGBA, error) {
	oldR := original.R
	oldG := original.G
	original.R = oldG
	original.G = oldR
	return original, nil
}
