package test

import (
	"testing"

	"github.com/mcsymiv/go-brand/format"
)

func TestFormat(t *testing.T) {

	format.FormatFiles("./", "suite")

}

func TestFeatures(t *testing.T) {

	format.FormatFeatures("./", "suite")

}
