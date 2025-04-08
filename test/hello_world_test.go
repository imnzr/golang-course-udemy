package test

import (
	"testing"

	"github.com/imnzr/golang-course-udemy/function"
)

func TestHelloWorld(t *testing.T) {
	result := function.HelloWorld("Nizar Fadilah")
	if result != "Nizar Fadilah" {
		panic("Result is not Nizar Fadilah")
	}
}
