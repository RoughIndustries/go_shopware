package models

import "testing"

func TestResults(t *testing.T) {
	got := Results{}
	if len(got.Results) > 0 {
		t.Errorf("Results(-1)")
	}
}
