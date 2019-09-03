package model

import (
	"fmt"
	"testing"
)

func TestParseRate(t *testing.T) {
	r, err := ParseRate("10M")
	fmt.Println(r, err)
	fmt.Println(r.String())
}
