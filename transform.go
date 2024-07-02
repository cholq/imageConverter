package main

import (
	"errors"
	"image/color"
	"log"
)

type TransformFn func([][]color.Color) ([][]color.Color, error)

func Grayscale(originalPixels [][]color.Color) ([][]color.Color, error) {
	return TransformPixelsOneByOne(GrayscaleTransformation, originalPixels)
}

func GrayAndBlue(originalPixels [][]color.Color) ([][]color.Color, error) {
	return TransformPixelsOneByOne(GrayscaleAndBlueTransformation, originalPixels)
}

func GrayAndGreen(originalPixels [][]color.Color) ([][]color.Color, error) {
	return TransformPixelsOneByOne(GrayscaleAndGreenTransformation, originalPixels)
}

func GrayAndRed(originalPixels [][]color.Color) ([][]color.Color, error) {
	return TransformPixelsOneByOne(GrayscaleAndRedTransformation, originalPixels)
}

func SwapRandGValues(originalPixels [][]color.Color) ([][]color.Color, error) {
	return TransformPixelsOneByOne(SwapRandGTransformation, originalPixels)
}

func SwapRandBValues(originalPixels [][]color.Color) ([][]color.Color, error) {
	return TransformPixelsOneByOne(SwapRandBTransformation, originalPixels)
}

func SwapGandBValues(originalPixels [][]color.Color) ([][]color.Color, error) {
	return TransformPixelsOneByOne(SwapGandBTransformation, originalPixels)
}

func ShiftRGBValuesLeft(originalPixels [][]color.Color) ([][]color.Color, error) {
	return TransformPixelsOneByOne(ShiftLeftTransformation, originalPixels)
}

func ShiftRGBValuesRight(originalPixels [][]color.Color) ([][]color.Color, error) {
	return TransformPixelsOneByOne(ShiftRightTransformation, originalPixels)
}

func Pixelate3x3(originalPixels [][]color.Color) ([][]color.Color, error) {
	return TransformPixelsThreeByThree(originalPixels)
}

func Pixelate10x10(originalPixels [][]color.Color) ([][]color.Color, error) {
	return TransformPixelsTenByTen((originalPixels))
}

func Pixelate20x20(originalPixels [][]color.Color) ([][]color.Color, error) {
	return TransformPixelsTwentyByTwenty((originalPixels))
}

func Pixelate50x50(originalPixels [][]color.Color) ([][]color.Color, error) {
	return TransformPixelsFiftyByFify((originalPixels))
}

func TransformImage(TxFn TransformFn, originalPixels [][]color.Color) ([][]color.Color, error) {
	newPixels, err := TxFn(originalPixels)
	if err != nil {
		log.Printf("Error transforming original image: %v", err)
		return nil, err
	}

	return newPixels, nil
}

func ProcessListOfTransformations(pixels [][]color.Color, transformationList []TransformationType) ([][]color.Color, error) {
	workingPixels := pixels
	var err error
	for _, transformVal := range transformationList {
		var transformedPixels [][]color.Color
		switch transformVal {
		case SwapRG:
			transformedPixels, err = TransformImage(SwapRandGValues, workingPixels)
		case SwapGB:
			transformedPixels, err = TransformImage(SwapGandBValues, workingPixels)
		case SwapRB:
			transformedPixels, err = TransformImage(SwapRandBValues, workingPixels)
		case Gray:
			transformedPixels, err = TransformImage(Grayscale, workingPixels)
		case GrayBlue:
			transformedPixels, err = TransformImage(GrayAndBlue, workingPixels)
		case GrayGreen:
			transformedPixels, err = TransformImage(GrayAndGreen, workingPixels)
		case GrayRed:
			transformedPixels, err = TransformImage(GrayAndRed, workingPixels)
		case ShiftLeft:
			transformedPixels, err = TransformImage(ShiftRGBValuesLeft, workingPixels)
		case ShiftRight:
			transformedPixels, err = TransformImage(ShiftRGBValuesRight, workingPixels)
		case Pixel3:
			transformedPixels, err = TransformImage(Pixelate3x3, workingPixels)
		case Pixel10:
			transformedPixels, err = TransformImage(Pixelate10x10, workingPixels)
		case Pixel20:
			transformedPixels, err = TransformImage(Pixelate20x20, workingPixels)
		case Pixel50:
			transformedPixels, err = TransformImage(Pixelate50x50, workingPixels)
		default:
			return workingPixels, errors.New("unknown transformation")
		}
		if err != nil {
			log.Fatalf("Cannot transform Image: %v: %v", transformVal, err)
			return workingPixels, err
		}
		workingPixels = transformedPixels
	}
	return workingPixels, nil
}

func TransformPixelsOneByOne(TxPixel TransformSinglePixelFn, originalPixels [][]color.Color) ([][]color.Color, error) {
	var newPixels [][]color.Color
	for xIndex := 0; xIndex < len(originalPixels); xIndex++ {
		pixelCol := originalPixels[xIndex]
		var newCol []color.Color
		for yIndex := 0; yIndex < len(pixelCol); yIndex++ {
			newPixel, err := TxPixel(pixelCol[yIndex])
			if err != nil {
				log.Printf("error transforming 1x1: %v", err)
				return nil, err
			}
			newCol = append(newCol, newPixel)
		}
		newPixels = append(newPixels, newCol)
	}
	return newPixels, nil
}

func TransformPixelsThreeByThree(originalPixels [][]color.Color) ([][]color.Color, error) {
	return TransformPixelsNxN(originalPixels, 3)
}

func TransformPixelsTenByTen(originalPixels [][]color.Color) ([][]color.Color, error) {
	return TransformPixelsNxN(originalPixels, 10)
}

func TransformPixelsTwentyByTwenty(originalPixels [][]color.Color) ([][]color.Color, error) {
	return TransformPixelsNxN(originalPixels, 20)
}

func TransformPixelsFiftyByFify(originalPixels [][]color.Color) ([][]color.Color, error) {
	return TransformPixelsNxN(originalPixels, 50)
}

func TransformPixelsNxN(originalPixels [][]color.Color, size int) ([][]color.Color, error) {
	var transformedPixels [][]color.Color
	transformedPixels = originalPixels

	for xIndex := 0; xIndex < len(originalPixels); xIndex += size {
		for yIndex := 0; yIndex < len(originalPixels[xIndex]); yIndex += size {

			pixelBlock, err := getPixelBlock(originalPixels, xIndex, yIndex, size)
			if err != nil {
				return nil, errors.New("could not get pixel block")
			}

			newColor, err := PixelBlockTransformation(pixelBlock)
			if err != nil {
				return nil, errors.New("could not calculate pixel block color")
			}

			transformedPixels, err = setPixelBlock(transformedPixels, newColor, xIndex, yIndex, size)
			if err != nil {
				return nil, errors.New("could not set pixel block")
			}

		}
	}

	return transformedPixels, nil
}

func getPixelBlock(originalPixels [][]color.Color, startX int, startY int, size int) ([]color.Color, error) {
	var pixelsInBlock []color.Color

	if startX >= len(originalPixels) {
		return pixelsInBlock, errors.New("x value too big for available pixel array")
	}

	if startY >= len(originalPixels[startX]) {
		return pixelsInBlock, errors.New("y value too big for available pixel array")
	}

	for xIndex := startX; xIndex < len(originalPixels) && xIndex < startX+size; xIndex++ {
		for yIndex := startY; yIndex < len(originalPixels[xIndex]) && yIndex < startY+size; yIndex++ {
			pixelsInBlock = append(pixelsInBlock, originalPixels[xIndex][yIndex])
		}
	}
	return pixelsInBlock, nil
}

func setPixelBlock(originalPixels [][]color.Color, newPixel color.Color, startX int, startY int, size int) ([][]color.Color, error) {

	if startX >= len(originalPixels) {
		return nil, errors.New("x value too big for available pixel array")
	}

	if startY >= len(originalPixels[startX]) {
		return nil, errors.New("y value too big for available pixel array")
	}

	for xIndex := startX; xIndex < len(originalPixels) && xIndex < startX+size; xIndex++ {
		for yIndex := startY; yIndex < len(originalPixels[xIndex]) && yIndex < startY+size; yIndex++ {
			originalPixels[xIndex][yIndex] = newPixel
		}
	}
	return originalPixels, nil
}
