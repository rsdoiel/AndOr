//
// Package andor provides support for building simple digital
// object repositories in Go where objects are stored in a
// dataset collection and the UI of the repository is static
// HTML 5 documents using JavaScript to access a web API.
//
// @Author R. S. Doiel, <rsdoiel@library.caltech.edu>
//
package andor

import (
	"testing"
)

func TestAndOr(t *testing.T) {
	if AndOrUsers == "" {
		t.Errorf("expected AndOrUsers to non-empty")
	}
	if AndOrWorkflows == "" {
		t.Errorf("expected AndOrWorkflows to non-empty")
	}
}
