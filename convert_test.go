package main

import (
	"image"
	"image/color"
	"testing"
)

func createGrayPixel() color.Color {
	grayRgba := color.RGBA{100, 100, 100, 255}
	grayPixel := color.RGBAModel.Convert(grayRgba)
	return grayPixel
}

func createGray2DArray() [][]color.Color {
	grayPixel := createGrayPixel()

	var xArray [][]color.Color
	for xIndex := 0; xIndex < 10; xIndex++ {
		var yArray []color.Color
		for yIndex := 0; yIndex < 10; yIndex++ {
			yArray = append(yArray, grayPixel)
		}
		xArray = append(xArray, yArray)
	}
	return xArray
}

func create2DArraySingleColor(c TestColor) [][]color.Color {

	var xArray [][]color.Color
	for xIndex := 0; xIndex < 10; xIndex++ {
		var yArray []color.Color
		for yIndex := 0; yIndex < 10; yIndex++ {
			yArray = append(yArray, c.color)
		}
		xArray = append(xArray, yArray)
	}
	return xArray
}

func createGray10x10Image() image.Image {
	upLeft := image.Point{0, 0}
	downRight := image.Point{10, 10}
	return image.NewGray(image.Rectangle{upLeft, downRight})
}

func TestCreatePixelArrayFromImage(t *testing.T) {
	img := createGray10x10Image()

	result, err := CreatePixelArrayFromImage(img)
	if err != nil {
		t.Errorf("Error creating Pixel array")
	}

	if len(result) != 10 {
		t.Errorf("Wrong dimension of first array")
	}

	if len(result[0]) != 10 {
		t.Errorf("Wrong dimension of second array")
	}
}

func TestCreateImageFromPixelArray(t *testing.T) {

	grayPixel := createGrayPixel()
	grayPixelArray := createGray2DArray()

	img, err := CreateImageFromPixelArray(grayPixelArray)
	if err != nil {
		t.Errorf("Error creating image")
	}

	for xIndex := 0; xIndex < 10; xIndex++ {
		for yIndex := 0; yIndex < 10; yIndex++ {
			if img.At(xIndex, yIndex) != grayPixel {
				t.Errorf("Wrong pixel value")
			}
		}
	}
}

func TestCreateImageFromPixelArrayFails(t *testing.T) {

	emptyPixelArray := [][]color.Color{}

	_, err := CreateImageFromPixelArray(emptyPixelArray)
	if err == nil {
		t.Errorf("Empty array should have thrown error")
	}
}
