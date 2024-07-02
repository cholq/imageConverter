package main

import (
	"strings"
	"testing"
)

type ParamTest struct {
	name                    string
	params                  []string
	expectedTransformations []TransformationType
	expectedShowHelp        bool
	expectedInputFile       string
	expectedOutputFile      string
	expectErr               bool
	errText                 string
}

func TestGetEmptyTransformationParams(t *testing.T) {
	result := getEmptyTransformationParams()

	if result.outputFile != "" {
		t.Errorf("getEmptyTransformationParams returned invalid OutputFile value: Expect: %v. Got: %v", "", result.outputFile)
	}

	if result.inputFile != "" {
		t.Errorf("getEmptyTransformationParams returned invalid InputFile value: Expect: %v. Got: %v", "", result.inputFile)
	}

	if result.showHelp != false {
		t.Errorf("getEmptyTransformationParams returned invalid showHelp value: Expect: %v. Got: %v", false, result.showHelp)
	}

	actualLength := len(result.transformList)
	if actualLength != 0 {
		t.Errorf("getEmptyTransformationParams returned incorrect number of transformations: Expect: %v. Got: %v", 0, actualLength)
	}
}

func TestParseParameters(t *testing.T) {
	var emptyParams []string
	showHelpParams := []string{"-help"}
	showHParams := []string{"-h"}
	swapRBParams := []string{"-srb"}
	swapRGParams := []string{"-srg"}
	swapGBParams := []string{"-sgb"}
	grayParams := []string{"-g"}
	grayBlueParams := []string{"-gb"}
	grayGreenParams := []string{"-gg"}
	grayRedParams := []string{"-gr"}
	shiftLeftParams := []string{"-l"}
	shiftRightParams := []string{"-r"}
	pixel3Params := []string{"-p3"}
	pixel10Params := []string{"-p10"}
	pixel20Params := []string{"-p20"}
	pixel50Params := []string{"-p50"}
	invalidParams := []string{"-x"}
	inputFileFlagOnly := []string{"-o", "abc.jpg", "-i"}
	inputFileFlagOnly2 := []string{"-i", "-o", "abc.jpg"}
	outputFileFlagOnly := []string{"-o", "-i", "xyz.jpg"}
	outputFileFlagOnly2 := []string{"-i", "xyz.jpg", "-o"}
	bothFileParams := []string{"-o", "abc.jpg", "-i", "xyz.jpg"}
	bothFileParams2 := []string{"-i", "xyz.jpg", "-o", "abc.jpg"}

	var emptyXfm []TransformationType
	swapRBXfm := []TransformationType{SwapRB}
	swapRGXfm := []TransformationType{SwapRG}
	swapGBXfm := []TransformationType{SwapGB}
	grayXfm := []TransformationType{Gray}
	grayBlueXfm := []TransformationType{GrayBlue}
	grayGreenXfm := []TransformationType{GrayGreen}
	grayRedXfm := []TransformationType{GrayRed}
	shiftLeftXfm := []TransformationType{ShiftLeft}
	shiftRightXfm := []TransformationType{ShiftRight}
	pixel3Xfm := []TransformationType{Pixel3}
	pixel10Xfm := []TransformationType{Pixel10}
	pixel20Xfm := []TransformationType{Pixel20}
	pixel50Xfm := []TransformationType{Pixel50}

	var tests = []ParamTest{
		{"NoParams", emptyParams, emptyXfm, false, "", "", true, "input file not properly defined"},
		{"InputFlagNoFile", inputFileFlagOnly, emptyXfm, false, "", "", true, "input file not properly defined"},
		{"InputFlagNoFile2", inputFileFlagOnly2, emptyXfm, false, "", "", true, "input file not properly defined"},
		{"OutputFlagNoFile", outputFileFlagOnly, emptyXfm, false, "", "", true, "output file not properly defined"},
		{"OutputFlagNoFile2", outputFileFlagOnly2, emptyXfm, false, "", "", true, "output file not properly defined"},
		{"BothFlagsWithFiles", bothFileParams, emptyXfm, false, "xyz.jpg", "abc.jpg", false, ""},
		{"BothFlagsWithFiles2", bothFileParams2, emptyXfm, false, "xyz.jpg", "abc.jpg", false, ""},
		{"ShowHelpOnly", showHelpParams, emptyXfm, true, "", "", false, ""},
		{"ShowHOnly", showHParams, emptyXfm, true, "", "", false, ""},
		{"SwapRBOnly", append(swapRBParams, bothFileParams...), swapRBXfm, false, "xyz.jpg", "abc.jpg", false, ""},
		{"SwapRBOnly", append(swapRBParams, bothFileParams...), swapRBXfm, false, "xyz.jpg", "abc.jpg", false, ""},
		{"SwapRGOnly", append(swapRGParams, bothFileParams...), swapRGXfm, false, "xyz.jpg", "abc.jpg", false, ""},
		{"SwapGBOnly", append(swapGBParams, bothFileParams...), swapGBXfm, false, "xyz.jpg", "abc.jpg", false, ""},
		{"GrayOnly", append(grayParams, bothFileParams...), grayXfm, false, "xyz.jpg", "abc.jpg", false, ""},
		{"GrayBlueOnly", append(grayBlueParams, bothFileParams...), grayBlueXfm, false, "xyz.jpg", "abc.jpg", false, ""},
		{"GrayGreenOnly", append(grayGreenParams, bothFileParams...), grayGreenXfm, false, "xyz.jpg", "abc.jpg", false, ""},
		{"GrayRedOnly", append(grayRedParams, bothFileParams...), grayRedXfm, false, "xyz.jpg", "abc.jpg", false, ""},
		{"ShiftLeftOnly", append(shiftLeftParams, bothFileParams...), shiftLeftXfm, false, "xyz.jpg", "abc.jpg", false, ""},
		{"ShiftRightOnly", append(shiftRightParams, bothFileParams...), shiftRightXfm, false, "xyz.jpg", "abc.jpg", false, ""},
		{"Pixel3Only", append(pixel3Params, bothFileParams...), pixel3Xfm, false, "xyz.jpg", "abc.jpg", false, ""},
		{"Pixel3Only", append(pixel10Params, bothFileParams...), pixel10Xfm, false, "xyz.jpg", "abc.jpg", false, ""},
		{"Pixel3Only", append(pixel20Params, bothFileParams...), pixel20Xfm, false, "xyz.jpg", "abc.jpg", false, ""},
		{"Pixel3Only", append(pixel50Params, bothFileParams...), pixel50Xfm, false, "xyz.jpg", "abc.jpg", false, ""},
		{"ParamCombo", append(append(swapRBParams, shiftLeftParams...), bothFileParams...), append(swapRBXfm, shiftLeftXfm...), false, "xyz.jpg", "abc.jpg", false, ""},
		{"InvalidCombo", append(append(swapRBParams, invalidParams...), bothFileParams...), swapRBXfm, false, "xyz.jpg", "abc.jpg", true, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseParameters(tt.params)
			if err != nil && !tt.expectErr {
				t.Errorf("Test %s returned an unexpected error: %v", tt.name, err)
			}

			if err != nil && tt.expectErr {
				if !strings.Contains(err.Error(), tt.errText) {
					t.Errorf("Test %s returned incorrect error. Want: %s. Got: %v", tt.name, tt.errText, err)
				}
			}

			if err == nil && tt.expectErr {
				t.Errorf("Test %v should have returned an error, but did not", tt.name)
			}

			if err == nil && !tt.expectErr {
				if tt.expectedShowHelp != result.showHelp {
					t.Errorf("Test %s returned invalid showHelp value: Expect: %v. Got: %v", tt.name, tt.expectedShowHelp, result.showHelp)
				}

				if tt.expectedInputFile != result.inputFile {
					t.Errorf("Test %s returned invalid InputFile value: Expect: %v. Got: %v", tt.name, tt.expectedInputFile, result.inputFile)
				}

				if tt.expectedOutputFile != result.outputFile {
					t.Errorf("Test %s returned invalid OutputFile value: Expect: %v. Got: %v", tt.name, tt.expectedOutputFile, result.outputFile)
				}

				expectedLength := len(tt.expectedTransformations)
				actualLength := len(result.transformList)
				if expectedLength != actualLength {
					t.Errorf("Test %s returned incorrect number of transformations: Expect: %v. Got: %v", tt.name, expectedLength, actualLength)
				} else {
					for vIdx, vVal := range result.transformList {
						if tt.expectedTransformations[vIdx] != vVal {
							t.Errorf("Test %s returned invalid Transformation at Index %d: Expect: %v. Got: %v", tt.name, vIdx, tt.expectedTransformations[vIdx], vVal)
						}
					}
				}
			}

		})
	}
}
