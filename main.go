package main

import (
	"fmt"
	"log"
	"os"
)

func showHelp() {
	fmt.Println("imagesTx.exe -i <input file> -o <output file> [transformation flags]")
	fmt.Println("")
	fmt.Println("Transformation Flags:")
	fmt.Println("  -g     Convert image to grayscale")
	fmt.Println("  -gb    Convert image to grayscale, maintain blue value")
	fmt.Println("  -gg    Convert image to grayscale, maintain green value")
	fmt.Println("  -gr    Convert image to grayscale, maintain red value")
	fmt.Println("  -l     Shift Left (Red -> Blue, Green -> Red, Blue -> Green)")
	fmt.Println("  -p3    Pixelate the image in 3x3 blocks")
	fmt.Println("  -p10   Pixelate the image in 10x10 blocks")
	fmt.Println("  -p20   Pixelate the image in 20x20 blocks")
	fmt.Println("  -p50   Pixelate the image in 50x50 blocks")
	fmt.Println("  -r     Shift Right (Red -> Green, Green -> Blue, Blue -> Red)")
	fmt.Println("  -sgb   Swap green and blue values")
	fmt.Println("  -srb   Swap red and blue values")
	fmt.Println("  -srg   Swap red and green values")
	fmt.Println("")
	fmt.Println("Multiple transformation flags can be combined.  They are processed in the order they are listed.")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("  imagesTx.exe -i start.jpg -o result.jpg -gg -l")
	fmt.Println("")
	fmt.Println("")
}

func main() {
	params, err := parseParameters(os.Args[1:])
	if err != nil {
		log.Fatalf("Error parsing parameters: %v", err)
	}

	if params.showHelp {
		showHelp()
		return
	}

	img, err := openJpeg(params.inputFile)
	if err != nil {
		log.Fatal("Cannot open file: Aborting")
	}

	pixels, err := CreatePixelArrayFromImage(img)
	if err != nil {
		log.Fatal("Cannot generate array: Aborting")
	}

	pixels, err = ProcessListOfTransformations(pixels, params.transformList)
	if err != nil {
		log.Fatalf("Cannot Transform Images: %v", err)
	}

	err = writeJpeg(pixels, params.outputFile)
	if err != nil {
		log.Fatal("Cannot write file: Aborting")
	}
}
