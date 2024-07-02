package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

func openJpeg(path string) (image.Image, error) {
	fileReader, err := os.Open(path)
	if err != nil {
		log.Printf("Error opening file: %s", err)
		return nil, err
	}
	defer fileReader.Close()

	img, _, err := image.Decode(fileReader)
	if err != nil {
		log.Printf("Error decoding image data: %s", err)
		return nil, nil
	}

	return img, nil
}

func writeJpeg(pixels [][]color.Color, filePath string) error {

	newImage, err := CreateImageFromPixelArray(pixels)
	if err != nil {
		log.Printf("Could not create new image: %v", err)
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		log.Printf("Could not write new image: %v", err)
		return err
	}
	defer file.Close()

	err = jpeg.Encode(file, newImage, nil)

	return err

}

func writePng(pixels [][]color.Color, filePath string) error {

	newImage, err := CreateImageFromPixelArray(pixels)
	if err != nil {
		log.Printf("Could not create new image: %v", err)
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		log.Printf("Could not write new image: %v", err)
		return err
	}
	defer file.Close()

	err = png.Encode(file, newImage)

	return err

}
