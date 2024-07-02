package main

import (
	"errors"
	"image/color"
	"strings"
	"testing"
)

type ColorTest struct {
	name      string
	fnToTest  TransformRGBAValuesFn
	input     color.Color
	expected  color.Color
	expectErr bool
	errText   string
}

type TransformTest struct {
	name      string
	fnToTest  TransformSinglePixelFn
	input     color.Color
	expected  color.Color
	expectErr bool
	errText   string
}

type PixelBlockTest struct {
	name      string
	input     []color.Color
	expected  color.Color
	expectErr bool
	errText   string
}

type TestColor struct {
	rgb   color.RGBA
	color color.Color
}

var testWhite = SetTestColor(255, 255, 255)
var testBlack = SetTestColor(0, 0, 0)
var testGray = SetTestColor(127, 127, 127)
var testPink = SetTestColor(255, 0, 127)
var testBlue = SetTestColor(0, 0, 255)
var testGrayBlue = SetTestColor(127, 127, 255)
var testGreen = SetTestColor(0, 255, 0)
var testGrayGreen = SetTestColor(127, 255, 127)
var testRed = SetTestColor(255, 0, 0)
var testGrayRed = SetTestColor(255, 127, 127)

func SetTestColor(R uint8, G uint8, B uint8) TestColor {
	rgb := color.RGBA{R, G, B, 255}
	return TestColor{rgb, color.RGBAModel.Convert(rgb)}
}

func mockTransformRGBAGood(color.RGBA) (color.RGBA, error) {
	return testBlack.rgb, nil
}

func mockTransformRGBABad(color.RGBA) (color.RGBA, error) {
	return color.RGBA{}, errors.New("Dummy Error")
}

func TestSinglePixelTransformation(t *testing.T) {

	var tests = []ColorTest{
		{"TransformWorks-A", mockTransformRGBAGood, testWhite.rgb, testBlack.rgb, false, ""},
		{"TransformFails-A", mockTransformRGBABad, testWhite.rgb, nil, true, "could not transform RGBA data"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := SinglePixelTransformation(tt.input, tt.fnToTest)
			if err != nil && !tt.expectErr {
				t.Errorf("Test %s returned an unexpected error: %v", tt.name, err)
			}

			if err != nil && tt.expectErr {
				if !strings.Contains(err.Error(), tt.errText) {
					t.Errorf("Test %s returned incorrect error. Want: %s. Got: %v", tt.name, tt.errText, err)
				}
			}

			if err == nil && !tt.expectErr {
				if tt.expected != result {
					t.Errorf("Test %s returned invalid result: Expect: %v. Got: %v", tt.name, tt.expected, result)
				}
			}

		})
	}
}

func TestPixelBlockTransformation(t *testing.T) {
	var tests = []PixelBlockTest{
		{"EmptyBlock", []color.Color{}, nil, true, "could not average RGBA data"},
		{"SingleColor", []color.Color{testGray.color}, testGray.color, false, ""},
		{"MultiColor", []color.Color{testWhite.color, testBlack.color}, testGray.color, false, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := PixelBlockTransformation(tt.input)
			if err != nil && !tt.expectErr {
				t.Errorf("Test %s returned an unexpected error: %v", tt.name, err)
			}

			if err != nil && tt.expectErr {
				if !strings.Contains(err.Error(), tt.errText) {
					t.Errorf("Test %s returned incorrect error. Want: %s. Got: %v", tt.name, tt.errText, err)
				}
			}

			if err == nil && !tt.expectErr {
				if tt.expected != result {
					t.Errorf("Test %s returned invalid result: Expect: %v. Got: %v", tt.name, tt.expected, result)
				}
			}
		})
	}
}

func TestTranformations(t *testing.T) {
	var tests = []TransformTest{
		{"GrayscaleTransformation-A", GrayscaleTransformation, testPink.color, testGray.color, false, ""},
		{"GrayscaleAndBlue-A", GrayscaleAndBlueTransformation, testBlue.color, testGrayBlue.color, false, ""},
		{"GrayscaleAndGreen-A", GrayscaleAndGreenTransformation, testGreen.color, testGrayGreen.color, false, ""},
		{"GrayscaleAndRed-A", GrayscaleAndRedTransformation, testRed.color, testGrayRed.color, false, ""},
		{"ShiftLeft-A", ShiftLeftTransformation, testGreen.color, testRed.color, false, ""},
		{"ShiftRight-A", ShiftRightTransformation, testGreen.color, testBlue.color, false, ""},
		{"SwapGandB", SwapGandBTransformation, testGreen.color, testBlue.color, false, ""},
		{"SwapRandB", SwapRandBTransformation, testRed.color, testBlue.color, false, ""},
		{"SwapRandG", SwapRandGTransformation, testRed.color, testGreen.color, false, ""},
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
				if tt.expected != result {
					t.Errorf("Test %s returned invalid result: Expect: %v. Got: %v", tt.name, tt.expected, result)
				}
			}
		})
	}
}
