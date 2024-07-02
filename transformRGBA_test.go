package main

import (
	"image/color"
	"strings"
	"testing"
)

type GrayscaleTest struct {
	name      string
	input     color.RGBA
	expected  uint8
	expectErr bool
	errText   string
}

type AverageValueTest struct {
	name      string
	input     []color.RGBA
	expected  color.RGBA
	expectErr bool
	errText   string
}

type RBGATest struct {
	name      string
	fnToTest  TransformRGBAValuesFn
	input     color.RGBA
	expected  color.RGBA
	expectErr bool
	errText   *string
}

func TestRGBAAverage(t *testing.T) {
	var emptyInput []color.RGBA
	singleInput := []color.RGBA{{255, 255, 255, 255}}
	twoInputs := []color.RGBA{{255, 255, 255, 255}, {0, 0, 0, 255}}
	var tests = []AverageValueTest{
		{"EmptyInput", emptyInput, color.RGBA{}, true, "divide by zero"},
		{"SingleInput", singleInput, color.RGBA{255, 255, 255, 255}, false, ""},
		{"TwoInputs", twoInputs, color.RGBA{127, 127, 127, 255}, false, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := RGBAAverage(tt.input)
			if err != nil && !tt.expectErr {
				t.Errorf("Average Value Test %s returned an unexpected error: %v", tt.name, err)
			}

			if err != nil && tt.expectErr {
				if !strings.Contains(err.Error(), tt.errText) {
					t.Errorf("Average Value Test %s returned incorrect error. Want: %s. Got: %v", tt.name, tt.errText, err)
				}
			}

			if err == nil && !tt.expectErr {
				if tt.expected != result {
					t.Errorf("Average Value Test %s returned invalid result: Expect: %v. Got: %v", tt.name, tt.expected, result)
				}
			}

		})
	}
}

func TestCalcGrayscaleValue(t *testing.T) {
	var tests = []GrayscaleTest{
		{"calcGrayscale-A", color.RGBA{0, 0, 0, 0}, 0, false, ""},
		{"calcGrayscale-B", color.RGBA{255, 255, 255, 255}, 255, false, ""},
		{"calcGrayscale-C", color.RGBA{0, 255, 255, 255}, 127, false, ""},
		{"calcGrayscale-D", color.RGBA{0, 100, 200, 255}, 100, false, ""},
		{"calcGrayscale-E", color.RGBA{100, 100, 100, 255}, 100, false, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calcGrayscaleValue(tt.input)
			if tt.expected != result {
				t.Errorf("Test %s returned invalid result: Expect: %v. Got: %v", tt.name, tt.expected, result)
			}
		})
	}
}

func TestRGBATransformations(t *testing.T) {
	var tests = []RBGATest{
		{"Grayscale-A", RGBAGrayscale, color.RGBA{10, 20, 30, 100}, color.RGBA{20, 20, 20, 100}, false, nil},
		{"GrayscaleAndBlue-A", RGBAGrayscaleAndBlue, color.RGBA{10, 20, 30, 100}, color.RGBA{20, 20, 30, 100}, false, nil},
		{"GrayscaleAndGreen-A", RGBAGrayscaleAndGreen, color.RGBA{10, 21, 30, 100}, color.RGBA{20, 21, 20, 100}, false, nil},
		{"GrayscaleAndRed-A", RGBAGrayscaleAndRed, color.RGBA{10, 20, 30, 100}, color.RGBA{10, 20, 20, 100}, false, nil},
		{"ShiftLeft-A", RGBAShiftLeft, color.RGBA{10, 20, 30, 100}, color.RGBA{20, 30, 10, 100}, false, nil},
		{"ShiftRight-A", RGBAShiftRight, color.RGBA{10, 20, 30, 100}, color.RGBA{30, 10, 20, 100}, false, nil},
		{"SwapGandB-A", RGBASwapGandB, color.RGBA{10, 20, 30, 100}, color.RGBA{10, 30, 20, 100}, false, nil},
		{"SwapRandB-A", RGBASwapRandB, color.RGBA{10, 20, 30, 100}, color.RGBA{30, 20, 10, 100}, false, nil},
		{"SwapRandG-A", RGBASwapRandG, color.RGBA{10, 20, 30, 100}, color.RGBA{20, 10, 30, 100}, false, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.fnToTest(tt.input)
			if err != nil && !tt.expectErr {
				t.Errorf("Test %s returned an unexpected error: %v", tt.name, err)
			}

			if err != nil && tt.expectErr {
				if !strings.Contains(err.Error(), *tt.errText) {
					t.Errorf("Test %s returned incorrect error. Want: %s. Got: %v", tt.name, *tt.errText, err)
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
