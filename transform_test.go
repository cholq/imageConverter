package main

import (
	"errors"
	"image/color"
	"strings"
	"testing"
)

type TransformFnTest struct {
	name      string
	fnToTest  TransformFn
	input     [][]color.Color
	expected  [][]color.Color
	expectErr bool
	errText   string
}

type ProcessListTest struct {
	name      string
	list      []TransformationType
	input     [][]color.Color
	expected  [][]color.Color
	expectErr bool
	errText   string
}

type TransformOneByOneTest struct {
	name      string
	fnToTest  TransformSinglePixelFn
	input     [][]color.Color
	expected  [][]color.Color
	expectErr bool
	errText   string
}

type TransformNxNTest struct {
	name      string
	input     [][]color.Color
	expect    [][]color.Color
	expectErr bool
	errText   string
}

type GetPixelTest struct {
	name        string
	inputPixels [][]color.Color
	inputX      int
	inputY      int
	inputSize   int
	expected    []color.Color
	expectErr   bool
	errText     string
}

type SetPixelTest struct {
	name        string
	inputPixels [][]color.Color
	inputColor  color.Color
	inputX      int
	inputY      int
	inputSize   int
	expected    [][]color.Color
	expectErr   bool
	errText     string
}

func mockTransformFuncSucceed(original [][]color.Color) ([][]color.Color, error) {
	return original, nil
}

func mockTransformFuncFail(original [][]color.Color) ([][]color.Color, error) {
	return original, errors.New("Mock Error")
}

func mockTransformOneByOneSucceed(original color.Color) (color.Color, error) {
	return original, nil
}

func mockTransformOneByOneFailed(original color.Color) (color.Color, error) {
	return original, errors.New(("Mock Error"))
}

var emptyColorArray [][]color.Color

func TestTransformImage(t *testing.T) {
	var tests = []TransformFnTest{
		{"TransformImage-A", mockTransformFuncSucceed, emptyColorArray, emptyColorArray, false, ""},
		{"TransformImage-B", mockTransformFuncFail, emptyColorArray, emptyColorArray, true, "Mock Error"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := TransformImage(tt.fnToTest, tt.input)
			if err != nil && !tt.expectErr {
				t.Errorf("Test %s returned an unexpected error: %v", tt.name, err)
			}

			if err != nil && tt.expectErr {
				if !strings.Contains(err.Error(), tt.errText) {
					t.Errorf("Test %s returned incorrect error. Want: %s. Got: %v", tt.name, tt.errText, err)
				}
			}

			if err == nil && !tt.expectErr {
				if len(tt.expected) != len(result) {
					t.Errorf("Test %s returned invalid result: Expect: %v. Got: %v", tt.name, tt.expected, result)
				}
			}
		})
	}
}

func TestProcessListOfTransformations(t *testing.T) {
	var tests = []ProcessListTest{
		{"swapRG", []TransformationType{SwapRG}, create2DArraySingleColor(testRed), create2DArraySingleColor(testGreen), false, ""},
		{"swapGB", []TransformationType{SwapGB}, create2DArraySingleColor(testGreen), create2DArraySingleColor(testBlue), false, ""},
		{"SwapRB", []TransformationType{SwapRB}, create2DArraySingleColor(testRed), create2DArraySingleColor(testBlue), false, ""},
		{"ShiftLeft", []TransformationType{ShiftLeft}, create2DArraySingleColor(testRed), create2DArraySingleColor(testBlue), false, ""},
		{"ShiftRight", []TransformationType{ShiftRight}, create2DArraySingleColor(testRed), create2DArraySingleColor(testGreen), false, ""},
		{"Gray", []TransformationType{Gray}, create2DArraySingleColor(testRed), create2DArraySingleColor(testGray), false, ""},
		{"GrayBlue", []TransformationType{GrayBlue}, create2DArraySingleColor(testBlue), create2DArraySingleColor(testGrayBlue), false, ""},
		{"GrayGreen", []TransformationType{GrayGreen}, create2DArraySingleColor(testGreen), create2DArraySingleColor(testGrayGreen), false, ""},
		{"GrayRed", []TransformationType{GrayRed}, create2DArraySingleColor(testRed), create2DArraySingleColor(testGrayRed), false, ""},
		{"Pixel-3", []TransformationType{Pixel3}, create2DArraySingleColor(testRed), create2DArraySingleColor(testRed), false, ""},
		{"Pixel-10", []TransformationType{Pixel10}, create2DArraySingleColor(testRed), create2DArraySingleColor(testRed), false, ""},
		{"Pixel-20", []TransformationType{Pixel20}, create2DArraySingleColor(testRed), create2DArraySingleColor(testRed), false, ""},
		{"Pixel-50", []TransformationType{Pixel50}, create2DArraySingleColor(testRed), create2DArraySingleColor(testRed), false, ""},
		{"Multiple 1", []TransformationType{ShiftLeft, ShiftRight}, create2DArraySingleColor(testRed), create2DArraySingleColor(testRed), false, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ProcessListOfTransformations(tt.input, tt.list)
			if err != nil && !tt.expectErr {
				t.Errorf("Test %s returned an unexpected error: %v", tt.name, err)
			}

			if err != nil && tt.expectErr {
				if !strings.Contains(err.Error(), tt.errText) {
					t.Errorf("Test %s returned incorrect error. Want: %s. Got: %v", tt.name, tt.errText, err)
				}
			}

			if err == nil && !tt.expectErr {
				if len(tt.expected) != len(result) {
					t.Errorf("Test %s returned invalid result length: Expect: %v. Got: %v", tt.name, len(tt.expected), len(result))
				} else {
					if len(tt.expected[0]) != len(result[0]) {
						t.Errorf("Test %s returned invalid result length2: Expect: %v. Got: %v", tt.name, len(tt.expected[0]), len(result[0]))
					}
				}
			}

		out:
			for xIndex := 0; xIndex < len(tt.expected); xIndex++ {
				for yIndex := 0; yIndex < len(tt.expected[0]); yIndex++ {
					if tt.expected[xIndex][yIndex] != result[xIndex][yIndex] {
						t.Errorf("Test %s returned invalid color: Expect: %v. Got: %v", tt.name, tt.expected[xIndex][yIndex], result[xIndex][yIndex])
						break out
					}
				}
			}

		})
	}
}

func TestTransformPixelsOneByOne(t *testing.T) {

	var colorArrayCol []color.Color
	colorArrayCol = append(colorArrayCol, testWhite.color)
	colorArrayCol = append(colorArrayCol, testBlack.color)
	var colorArrayData [][]color.Color
	colorArrayData = append(colorArrayData, colorArrayCol)
	colorArrayData = append(colorArrayData, colorArrayCol)

	var tests = []TransformOneByOneTest{
		{"TransformPixelsOneByOne-A", mockTransformOneByOneSucceed, colorArrayData, colorArrayData, false, ""},
		{"TransformPixelsOneByOne-B", mockTransformOneByOneFailed, colorArrayData, colorArrayData, true, "Mock Error"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := TransformPixelsOneByOne(tt.fnToTest, tt.input)
			if err != nil && !tt.expectErr {
				t.Errorf("Test %s returned an unexpected error: %v", tt.name, err)
			}

			if err != nil && tt.expectErr {
				if !strings.Contains(err.Error(), tt.errText) {
					t.Errorf("Test %s returned incorrect error. Want: %s. Got: %v", tt.name, tt.errText, err)
				}
			}

			if err == nil && !tt.expectErr {
				if len(tt.expected) != len(result) {
					t.Errorf("Test %s returned invalid result: Expect: %v. Got: %v", tt.name, tt.expected, result)
				}
			}
		})
	}
}

func TestOneByOneTransformations(t *testing.T) {

	// var colorArrayWhiteCol []color.Color
	// colorArrayWhiteCol = append(colorArrayWhiteCol, testWhite.color)
	// colorArrayWhiteCol = append(colorArrayWhiteCol, testWhite.color)
	// var colorArrayWhiteData [][]color.Color
	// colorArrayWhiteData = append(colorArrayWhiteData, colorArrayWhiteCol)
	// colorArrayWhiteData = append(colorArrayWhiteData, colorArrayWhiteCol)

	var colorArrayGrayCol []color.Color
	colorArrayGrayCol = append(colorArrayGrayCol, testGray.color)
	colorArrayGrayCol = append(colorArrayGrayCol, testGray.color)
	var colorArrayGrayData [][]color.Color
	colorArrayGrayData = append(colorArrayGrayData, colorArrayGrayCol)
	colorArrayGrayData = append(colorArrayGrayData, colorArrayGrayCol)

	var colorArrayBlueCol []color.Color
	colorArrayBlueCol = append(colorArrayBlueCol, testBlue.color)
	colorArrayBlueCol = append(colorArrayBlueCol, testBlue.color)
	var colorArrayBlueData [][]color.Color
	colorArrayBlueData = append(colorArrayBlueData, colorArrayBlueCol)
	colorArrayBlueData = append(colorArrayBlueData, colorArrayBlueCol)

	var colorArrayGrayBlueCol []color.Color
	colorArrayGrayBlueCol = append(colorArrayGrayBlueCol, testGrayBlue.color)
	colorArrayGrayBlueCol = append(colorArrayGrayBlueCol, testGrayBlue.color)
	var colorArrayGrayBlueData [][]color.Color
	colorArrayGrayBlueData = append(colorArrayGrayBlueData, colorArrayGrayBlueCol)
	colorArrayGrayBlueData = append(colorArrayGrayBlueData, colorArrayGrayBlueCol)

	var colorArrayRedCol []color.Color
	colorArrayRedCol = append(colorArrayRedCol, testRed.color)
	colorArrayRedCol = append(colorArrayRedCol, testRed.color)
	var colorArrayRedData [][]color.Color
	colorArrayRedData = append(colorArrayRedData, colorArrayRedCol)
	colorArrayRedData = append(colorArrayRedData, colorArrayRedCol)

	var colorArrayRedGrayCol []color.Color
	colorArrayRedGrayCol = append(colorArrayRedGrayCol, testGrayRed.color)
	colorArrayRedGrayCol = append(colorArrayRedGrayCol, testGrayRed.color)
	var colorArrayRedGrayData [][]color.Color
	colorArrayRedGrayData = append(colorArrayRedGrayData, colorArrayRedGrayCol)
	colorArrayRedGrayData = append(colorArrayRedGrayData, colorArrayRedGrayCol)

	var colorArrayGreenCol []color.Color
	colorArrayGreenCol = append(colorArrayGreenCol, testGreen.color)
	colorArrayGreenCol = append(colorArrayGreenCol, testGreen.color)
	var colorArrayGreenData [][]color.Color
	colorArrayGreenData = append(colorArrayGreenData, colorArrayGreenCol)
	colorArrayGreenData = append(colorArrayGreenData, colorArrayGreenCol)

	var colorArrayGreenGrayCol []color.Color
	colorArrayGreenGrayCol = append(colorArrayGreenGrayCol, testGrayGreen.color)
	colorArrayGreenGrayCol = append(colorArrayGreenGrayCol, testGrayGreen.color)
	var colorArrayGreenGrayData [][]color.Color
	colorArrayGreenGrayData = append(colorArrayGreenGrayData, colorArrayGreenGrayCol)
	colorArrayGreenGrayData = append(colorArrayGreenGrayData, colorArrayGreenGrayCol)

	var tests = []TransformFnTest{
		{"Grayscale-A", Grayscale, colorArrayBlueData, colorArrayGrayData, false, ""},
		{"GrayAndBlue-A", GrayAndBlue, colorArrayBlueData, colorArrayGrayBlueData, false, ""},
		{"GrayAndGreen-A", GrayAndGreen, colorArrayGreenData, colorArrayGreenGrayData, false, ""},
		{"GrayAndRed-A", GrayAndRed, colorArrayRedData, colorArrayRedGrayData, false, ""},
		{"SwapRandGValues-A", SwapRandGValues, colorArrayRedData, colorArrayGreenData, false, ""},
		{"SwapRandBValues-A", SwapRandBValues, colorArrayRedData, colorArrayBlueData, false, ""},
		{"SwapGandBValues-A", SwapGandBValues, colorArrayGreenData, colorArrayBlueData, false, ""},
		{"ShiftRGBValuesLeft-A", ShiftRGBValuesLeft, colorArrayGreenData, colorArrayRedData, false, ""},
		{"ShiftRGBValuesLeft-B", ShiftRGBValuesLeft, colorArrayRedData, colorArrayBlueData, false, ""},
		{"ShiftRGBValuesLeft-C", ShiftRGBValuesLeft, colorArrayBlueData, colorArrayGreenData, false, ""},
		{"ShiftRGBValuesRight-A", ShiftRGBValuesRight, colorArrayGreenData, colorArrayBlueData, false, ""},
		{"ShiftRGBValuesRight-B", ShiftRGBValuesRight, colorArrayBlueData, colorArrayRedData, false, ""},
		{"ShiftRGBValuesRight-B", ShiftRGBValuesRight, colorArrayRedData, colorArrayGreenData, false, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.fnToTest(tt.input)
			if err != nil && !tt.expectErr {
				t.Errorf("Test %s returned an unexpected error: %v", tt.name, err)
			}

			if err != nil && tt.expectErr {
				if !strings.Contains(err.Error(), tt.errText) {
					t.Errorf("Test %s returned incorrect error. Want: %s. Got: %v", tt.name, tt.errText, err)
				}
			}

			if err == nil && !tt.expectErr {
				if len(tt.expected) != len(result) {
					t.Errorf("Test %s returned invalid result: Expect: %v. Got: %v", tt.name, tt.expected, result)
				}

				for xIndex := 0; xIndex < len(tt.expected); xIndex++ {
					for yIndex := 0; yIndex < len(tt.expected[0]); yIndex++ {
						if tt.expected[xIndex][yIndex] != result[xIndex][yIndex] {
							t.Errorf("Test %s returned invalid pixel result: Expect: %v. Got: %v", tt.name, tt.expected[xIndex][yIndex], result[xIndex][yIndex])
						}
					}
				}
			}
		})
	}
}

func TestTransformPixelsThreeByThree(t *testing.T) {
	var tests = []TransformNxNTest{
		{"OriginalEmpty", [][]color.Color{}, [][]color.Color{}, false, ""},
		{"OneRow", [][]color.Color{{testWhite.color, testBlack.color}}, [][]color.Color{{testGray.color, testGray.color}}, false, ""},
		{"ThreeByThree", [][]color.Color{
			{testWhite.color, testWhite.color, testWhite.color},
			{testGray.color, testGray.color, testGray.color},
			{testBlack.color, testBlack.color, testBlack.color},
		}, [][]color.Color{
			{testGray.color, testGray.color, testGray.color},
			{testGray.color, testGray.color, testGray.color},
			{testGray.color, testGray.color, testGray.color},
		}, false, ""},
		{"SixBySix", [][]color.Color{
			{testWhite.color, testWhite.color, testWhite.color, testWhite.color, testWhite.color, testWhite.color},
			{testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color},
			{testBlack.color, testBlack.color, testBlack.color, testBlack.color, testBlack.color, testBlack.color},
			{testWhite.color, testWhite.color, testWhite.color, testWhite.color, testWhite.color, testWhite.color},
			{testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color},
			{testBlack.color, testBlack.color, testBlack.color, testBlack.color, testBlack.color, testBlack.color},
		}, [][]color.Color{
			{testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color},
			{testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color},
			{testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color},
			{testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color},
			{testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color},
			{testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color},
		}, false, ""},
		{"SevenBySeven", [][]color.Color{
			{testWhite.color, testWhite.color, testWhite.color, testWhite.color, testWhite.color, testWhite.color, testWhite.color},
			{testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testWhite.color},
			{testBlack.color, testBlack.color, testBlack.color, testBlack.color, testBlack.color, testBlack.color, testWhite.color},
			{testWhite.color, testWhite.color, testWhite.color, testWhite.color, testWhite.color, testWhite.color, testRed.color},
			{testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testRed.color},
			{testBlack.color, testBlack.color, testBlack.color, testBlack.color, testBlack.color, testBlack.color, testRed.color},
			{testBlue.color, testBlue.color, testBlue.color, testGreen.color, testGreen.color, testGreen.color, testBlack.color},
		}, [][]color.Color{
			{testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testWhite.color},
			{testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testWhite.color},
			{testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testWhite.color},
			{testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testRed.color},
			{testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testRed.color},
			{testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testRed.color},
			{testBlue.color, testBlue.color, testBlue.color, testGreen.color, testGreen.color, testGreen.color, testBlack.color},
		}, false, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := TransformPixelsThreeByThree(tt.input)
			if err != nil && !tt.expectErr {
				t.Errorf("Test %s returned an unexpected error: %v", tt.name, err)
			}

			if err != nil && tt.expectErr {
				if !strings.Contains(err.Error(), tt.errText) {
					t.Errorf("Test %s returned incorrect error. Want: %s. Got: %v", tt.name, tt.errText, err)
				}
			}

			if err == nil && tt.expectErr {
				t.Errorf("Test %s should have returned an error, but did not", tt.name)
			}

			if err == nil && !tt.expectErr {
				if len(tt.expect) != len(result) {
					t.Errorf("Test %s returned invalid result: Expect: %v. Got: %v", tt.name, tt.expect, result)
				}

				for xIndex := 0; xIndex < len(tt.expect); xIndex++ {
					for yIndex := 0; yIndex < len(tt.expect[xIndex]); yIndex++ {
						if tt.expect[xIndex][yIndex] != result[xIndex][yIndex] {
							t.Errorf("Test %s returned invalid pixel result: Expect: %v. Got: %v", tt.name, tt.expect[xIndex][yIndex], result[xIndex][yIndex])
						}
					}
				}
			}
		})
	}
}

func TestTransformPixelsTenByTen(t *testing.T) {
	var tests = []TransformNxNTest{
		{"OriginalEmpty", [][]color.Color{}, [][]color.Color{}, false, ""},
		{"OneRow", [][]color.Color{{testWhite.color, testBlack.color, testWhite.color, testBlack.color, testWhite.color, testBlack.color, testWhite.color, testBlack.color, testWhite.color, testBlack.color}}, [][]color.Color{{testGray.color, testGray.color}}, false, ""},
		{"ThreeByThree", [][]color.Color{
			{testWhite.color, testWhite.color, testWhite.color},
			{testGray.color, testGray.color, testGray.color},
			{testBlack.color, testBlack.color, testBlack.color},
		}, [][]color.Color{
			{testGray.color, testGray.color, testGray.color},
			{testGray.color, testGray.color, testGray.color},
			{testGray.color, testGray.color, testGray.color},
		}, false, ""},
		{"TenByTen", [][]color.Color{
			{testWhite.color, testWhite.color, testWhite.color, testWhite.color, testWhite.color, testWhite.color, testWhite.color, testWhite.color, testWhite.color, testWhite.color},
			{testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color},
			{testBlack.color, testBlack.color, testBlack.color, testBlack.color, testBlack.color, testBlack.color, testBlack.color, testBlack.color, testBlack.color, testBlack.color},
			{testWhite.color, testWhite.color, testWhite.color, testWhite.color, testWhite.color, testWhite.color, testWhite.color, testWhite.color, testWhite.color, testWhite.color},
			{testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color},
			{testBlack.color, testBlack.color, testBlack.color, testBlack.color, testBlack.color, testBlack.color, testBlack.color, testBlack.color, testBlack.color, testBlack.color},
			{testWhite.color, testWhite.color, testWhite.color, testWhite.color, testWhite.color, testWhite.color, testWhite.color, testWhite.color, testWhite.color, testWhite.color},
			{testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color},
			{testBlack.color, testBlack.color, testBlack.color, testBlack.color, testBlack.color, testBlack.color, testBlack.color, testBlack.color, testBlack.color, testBlack.color},
			{testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color},
		}, [][]color.Color{
			{testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color},
			{testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color},
			{testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color},
			{testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color},
			{testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color},
			{testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color},
			{testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color},
			{testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color},
			{testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color},
			{testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color, testGray.color},
		}, false, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := TransformPixelsTenByTen(tt.input)
			if err != nil && !tt.expectErr {
				t.Errorf("Test %s returned an unexpected error: %v", tt.name, err)
			}

			if err != nil && tt.expectErr {
				if !strings.Contains(err.Error(), tt.errText) {
					t.Errorf("Test %s returned incorrect error. Want: %s. Got: %v", tt.name, tt.errText, err)
				}
			}

			if err == nil && tt.expectErr {
				t.Errorf("Test %s should have returned an error, but did not", tt.name)
			}

			if err == nil && !tt.expectErr {
				if len(tt.expect) != len(result) {
					t.Errorf("Test %s returned invalid result: Expect: %v. Got: %v", tt.name, tt.expect, result)
				}

				for xIndex := 0; xIndex < len(tt.expect); xIndex++ {
					for yIndex := 0; yIndex < len(tt.expect[xIndex]); yIndex++ {
						if tt.expect[xIndex][yIndex] != result[xIndex][yIndex] {
							t.Errorf("Test %s returned invalid pixel result: Expect: %v. Got: %v", tt.name, tt.expect[xIndex][yIndex], result[xIndex][yIndex])
						}
					}
				}
			}
		})
	}
}

func TestGetPixelBlock(t *testing.T) {
	var tests = []GetPixelTest{
		{"NoOriginal", [][]color.Color{}, 0, 0, 1, nil, true, "x value too big"},
		{"BadX", [][]color.Color{{testGray.color, testGray.color}}, 1, 0, 1, nil, true, "x value too big"},
		{"BadY", [][]color.Color{{testGray.color, testGray.color}}, 0, 2, 1, nil, true, "y value too big"},
		{"FullBlock", [][]color.Color{
			{testGray.color, testGray.color, testGray.color, testGray.color},
			{testGray.color, testBlue.color, testRed.color, testGray.color},
			{testGray.color, testWhite.color, testBlack.color, testGray.color},
			{testGray.color, testGray.color, testGray.color, testGray.color},
		}, 1, 1, 2, []color.Color{testBlue.color, testRed.color, testWhite.color, testBlack.color}, false, ""},
		{"PartialXBlock", [][]color.Color{
			{testGray.color, testGray.color, testGray.color, testGray.color},
			{testGray.color, testGray.color, testGray.color, testBlue.color},
			{testGray.color, testGray.color, testGray.color, testWhite.color},
			{testGray.color, testGray.color, testGray.color, testGray.color},
		}, 1, 3, 2, []color.Color{testBlue.color, testWhite.color}, false, ""},
		{"PartialYBlock", [][]color.Color{
			{testGray.color, testGray.color, testGray.color, testGray.color},
			{testGray.color, testGray.color, testGray.color, testGray.color},
			{testGray.color, testGray.color, testGray.color, testGray.color},
			{testGray.color, testWhite.color, testBlack.color, testGray.color},
		}, 3, 1, 2, []color.Color{testWhite.color, testBlack.color}, false, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := getPixelBlock(tt.inputPixels, tt.inputX, tt.inputY, tt.inputSize)
			if err != nil && !tt.expectErr {
				t.Errorf("Test %s returned an unexpected error: %v", tt.name, err)
			}

			if err != nil && tt.expectErr {
				if !strings.Contains(err.Error(), tt.errText) {
					t.Errorf("Test %s returned incorrect error. Want: %s. Got: %v", tt.name, tt.errText, err)
				}
			}

			if err == nil && !tt.expectErr {
				if len(tt.expected) != len(result) {
					t.Errorf("Test %s returned invalid result: Expect: %v. Got: %v", tt.name, tt.expected, result)
				}

				for xIndex := 0; xIndex < len(tt.expected); xIndex++ {
					if tt.expected[xIndex] != result[xIndex] {
						t.Errorf("Test %s returned invalid pixel result: Expect: %v. Got: %v", tt.name, tt.expected[xIndex], result[xIndex])
					}
				}
			}
		})
	}

}

func TestSetPixelBlock(t *testing.T) {
	var tests = []SetPixelTest{
		{"NoOriginal1", [][]color.Color{}, testGray.color, 0, 0, 1, nil, true, "x value too big"},
		{"NoOriginal2", [][]color.Color{}, testGray.color, 1, 1, 3, nil, true, "x value too big"},
		{"BadOriginal1", [][]color.Color{{testGray.color, testGray.color}}, testGray.color, 0, 2, 3, nil, true, "y value too big"},
		{"ReplaceFullBlock", [][]color.Color{
			{testGray.color, testGray.color, testGray.color, testGray.color},
			{testGray.color, testGray.color, testGray.color, testGray.color},
			{testGray.color, testGray.color, testGray.color, testGray.color},
			{testGray.color, testGray.color, testGray.color, testGray.color},
		}, testRed.color, 1, 1, 2, [][]color.Color{
			{testGray.color, testGray.color, testGray.color, testGray.color},
			{testGray.color, testRed.color, testRed.color, testGray.color},
			{testGray.color, testRed.color, testRed.color, testGray.color},
			{testGray.color, testGray.color, testGray.color, testGray.color},
		}, false, ""},
		{"ReplacePartialXBlock", [][]color.Color{
			{testGray.color, testGray.color, testGray.color, testGray.color},
			{testGray.color, testGray.color, testGray.color, testGray.color},
			{testGray.color, testGray.color, testGray.color, testGray.color},
			{testGray.color, testGray.color, testGray.color, testGray.color},
		}, testRed.color, 1, 3, 2, [][]color.Color{
			{testGray.color, testGray.color, testGray.color, testGray.color},
			{testGray.color, testGray.color, testGray.color, testRed.color},
			{testGray.color, testGray.color, testGray.color, testRed.color},
			{testGray.color, testGray.color, testGray.color, testGray.color},
		}, false, ""},
		{"ReplacePartialYBlock", [][]color.Color{
			{testGray.color, testGray.color, testGray.color, testGray.color},
			{testGray.color, testGray.color, testGray.color, testGray.color},
			{testGray.color, testGray.color, testGray.color, testGray.color},
			{testGray.color, testGray.color, testGray.color, testGray.color},
		}, testRed.color, 3, 1, 2, [][]color.Color{
			{testGray.color, testGray.color, testGray.color, testGray.color},
			{testGray.color, testGray.color, testGray.color, testGray.color},
			{testGray.color, testGray.color, testGray.color, testGray.color},
			{testGray.color, testRed.color, testRed.color, testGray.color},
		}, false, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := setPixelBlock(tt.inputPixels, tt.inputColor, tt.inputX, tt.inputY, tt.inputSize)
			if err != nil && !tt.expectErr {
				t.Errorf("Test %s returned an unexpected error: %v", tt.name, err)
			}

			if err != nil && tt.expectErr {
				if !strings.Contains(err.Error(), tt.errText) {
					t.Errorf("Test %s returned incorrect error. Want: %s. Got: %v", tt.name, tt.errText, err)
				}
			}

			if err == nil && !tt.expectErr {
				if len(tt.expected) != len(result) {
					t.Errorf("Test %s returned invalid result: Expect: %v. Got: %v", tt.name, tt.expected, result)
				}

				for xIndex := 0; xIndex < len(tt.expected); xIndex++ {
					for yIndex := 0; yIndex < len(tt.expected[0]); yIndex++ {
						if tt.expected[xIndex][yIndex] != result[xIndex][yIndex] {
							t.Errorf("Test %s returned invalid pixel result: Expect: %v. Got: %v", tt.name, tt.expected[xIndex][yIndex], result[xIndex][yIndex])
						}
					}
				}
			}
		})
	}
}
