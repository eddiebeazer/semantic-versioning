package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type VersionTest struct {
	Version        string
	Message        string
	ExpectedResult string
}

func TestGetNextVersion(t *testing.T) {
	TestCases := [4]VersionTest{
		{Version: "4.74.54", Message: "semver: major", ExpectedResult: "5.0.0"},
		{Version: "4.74.54", Message: "semver: minor", ExpectedResult: "4.75.0"},
		{Version: "4.74.54", Message: "semver: patch", ExpectedResult: "4.74.55"},
		{Version: "4.74.54", Message: "empty commit", ExpectedResult: "0.0.0"},
	}
	for _, testCase := range TestCases {
		version := GetNextVersion(testCase.Version, testCase.Message)
		fmt.Println(version)
		assert.Equal(t, version, testCase.ExpectedResult, "Mismatch for version type")
	}
}
