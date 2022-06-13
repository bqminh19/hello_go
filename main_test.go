package main

import (
	"errors"
	"github.com/pborman/uuid"
	"log"
	"strings"
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
	err := testEllipsis(Point{0, 5}, 15)
	if err != nil {
		log.Fatal(err)
	}
}

func TestUUID(t *testing.T) {
	uuidWH := uuid.NewRandom()
	uuids := strings.Replace(uuidWH.String(), "-", "", -1)
	if len(uuids) == 0 {
		log.Fatal(errors.New("can not use package uuid"))
	}
}
