package main

import (
	"errors"
	"fmt"
	"strings"
)

type Transformation struct {
	transformList []TransformationType
	inputFile     string
	outputFile    string
	showHelp      bool
}

type TransformationType int64

const (
	Undefined TransformationType = iota
	Gray
	GrayBlue
	GrayGreen
	GrayRed
	Pixel3
	Pixel10
	Pixel20
	Pixel50
	ShiftLeft
	ShiftRight
	SwapGB
	SwapRB
	SwapRG
)

func getEmptyTransformationParams() Transformation {
	var transformParams Transformation
	transformParams.inputFile = ""
	transformParams.outputFile = ""
	transformParams.showHelp = false
	transformParams.transformList = []TransformationType{}

	return transformParams
}

func parseParameters(args []string) (Transformation, error) {

	transformParams := getEmptyTransformationParams()
	var returnImmediately bool
	nextValueInput := false
	nextValueOutput := false

	for _, a := range args {

		returnImmediately = false

		firstChar := a[0:1]
		if firstChar == "-" {
			// This is a flag, don't use value as a file name
			nextValueInput = false
			nextValueOutput = false
		}

		if nextValueInput {
			transformParams.inputFile = a
			nextValueInput = false
		} else if nextValueOutput {
			transformParams.outputFile = a
			nextValueOutput = false
		} else {
			switch a {
			case "-i":
				nextValueInput = true
			case "-o":
				nextValueOutput = true
			case "-help":
				fallthrough
			case "-h":
				transformParams = getEmptyTransformationParams()
				transformParams.showHelp = true
				returnImmediately = true
			case "-g":
				transformParams.transformList = append(transformParams.transformList, Gray)
			case "-gb":
				transformParams.transformList = append(transformParams.transformList, GrayBlue)
			case "-gg":
				transformParams.transformList = append(transformParams.transformList, GrayGreen)
			case "-gr":
				transformParams.transformList = append(transformParams.transformList, GrayRed)
			case "-p3":
				transformParams.transformList = append(transformParams.transformList, Pixel3)
			case "-p10":
				transformParams.transformList = append(transformParams.transformList, Pixel10)
			case "-p20":
				transformParams.transformList = append(transformParams.transformList, Pixel20)
			case "-p50":
				transformParams.transformList = append(transformParams.transformList, Pixel50)
			case "-sgb":
				transformParams.transformList = append(transformParams.transformList, SwapGB)
			case "-srb":
				transformParams.transformList = append(transformParams.transformList, SwapRB)
			case "-srg":
				transformParams.transformList = append(transformParams.transformList, SwapRG)
			case "-l":
				transformParams.transformList = append(transformParams.transformList, ShiftLeft)
			case "-r":
				transformParams.transformList = append(transformParams.transformList, ShiftRight)
			default:
				return transformParams, fmt.Errorf("unknown transformation flag: %v", a)
			}
		}

		if returnImmediately {
			return transformParams, nil
		}
	}

	if strings.TrimSpace(transformParams.inputFile) == "" {
		return transformParams, errors.New("input file not properly defined")
	}

	if strings.TrimSpace(transformParams.outputFile) == "" {
		return transformParams, errors.New("output file not properly defined")
	}

	return transformParams, nil
}
