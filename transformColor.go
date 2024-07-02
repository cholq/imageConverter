package main

import (
	"errors"
	"image/color"
	"log"
)

type TransformSinglePixelFn func(color.Color) (color.Color, error)

func SinglePixelTransformation(original color.Color, transformRGBA TransformRGBAValuesFn) (color.Color, error) {
	pixelRGBA, ok := color.RGBAModel.Convert(original).(color.RGBA)
	if ok {
		newRGBA, err := transformRGBA(pixelRGBA)
		if err != nil {
			log.Printf("could not transform RGBA: %v", pixelRGBA)
			return nil, errors.New("could not transform RGBA data")
		}

		newPixel := color.RGBAModel.Convert(newRGBA)
		return newPixel, nil

	} else {
		log.Printf("could not transform pixel: %v", original)
		return nil, errors.New("could not transform pixel")
	}
}

func PixelBlockTransformation(originals []color.Color) (color.Color, error) {
	var rgbaPixels []color.RGBA
	for _, pixel := range originals {
		pixelRGBA, ok := color.RGBAModel.Convert(pixel).(color.RGBA)
		if ok {
			rgbaPixels = append(rgbaPixels, pixelRGBA)
		}
	}

	newRGBA, err := RGBAAverage(rgbaPixels)
	if err != nil {
		log.Printf("could not average RGBA: %v", err)
		return nil, errors.New("could not average RGBA data")
	}

	newPixel := color.RGBAModel.Convert(newRGBA)
	return newPixel, nil

}

func GrayscaleTransformation(original color.Color) (color.Color, error) {
	return SinglePixelTransformation(original, RGBAGrayscale)
}

func GrayscaleAndBlueTransformation(original color.Color) (color.Color, error) {
	return SinglePixelTransformation(original, RGBAGrayscaleAndBlue)
}

func GrayscaleAndGreenTransformation(original color.Color) (color.Color, error) {
	return SinglePixelTransformation(original, RGBAGrayscaleAndGreen)
}

func GrayscaleAndRedTransformation(original color.Color) (color.Color, error) {
	return SinglePixelTransformation(original, RGBAGrayscaleAndRed)
}

func ShiftLeftTransformation(original color.Color) (color.Color, error) {
	return SinglePixelTransformation(original, RGBAShiftLeft)
}

func ShiftRightTransformation(original color.Color) (color.Color, error) {
	return SinglePixelTransformation(original, RGBAShiftRight)
}

func SwapGandBTransformation(original color.Color) (color.Color, error) {
	return SinglePixelTransformation(original, RGBASwapGandB)
}

func SwapRandBTransformation(original color.Color) (color.Color, error) {
	return SinglePixelTransformation(original, RGBASwapRandB)
}

func SwapRandGTransformation(original color.Color) (color.Color, error) {
	return SinglePixelTransformation(original, RGBASwapRandG)
}
