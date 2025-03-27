package test

import (
	"regexp"
	"testing"
)

func Hello(t *testing.T) {
	name, msg := "Gladys", "Glady"
	want := regexp.MustCompile(`\b` + name + `\b`)
	if !want.MatchString(msg) {
		t.Errorf(`Hello("Gladys") = %q, want match for %#q, nil`, msg, want)
	}
}
