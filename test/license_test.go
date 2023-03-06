package test

import (
	"fmt"
	"github.com/krumberg/gocopyleft"
	"testing"
)

func TestLicense(t *testing.T) {
	fmt.Println(gocopyleft.Stallman())
}
