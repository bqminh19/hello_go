package main

import (
	"errors"
	"testing"
)

func testEllipsis(i ...interface{}) error {
	a, aok := i[0].(Point)
	b, bok := i[1].(int)

	if !aok || !bok {
		return errors.New("unexpected values")
	}

	a.x = a.y + b
	a.y = a.x - b

	return nil
}

func TestEllipsis(t *testing.T) {
	testEllipsis(Point{0, 5}, 15)
}
