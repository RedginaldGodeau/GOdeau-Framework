package Controllers

import (
	"GOdeau/godeau/Controllers/Methods"
	"testing"
)

func TestCanAccessMethodGoodAccess(t *testing.T) {

	got := Methods.List(Methods.GET)
	recive := "GET"

	if !canAccess(got, recive) {
		t.Errorf("got false, wanted true")
	}
}

func TestCanAccessMethodBadAccess(t *testing.T) {

	got := Methods.List(Methods.GET)
	recive := "POST"

	if canAccess(got, recive) {
		t.Errorf("got true, wanted false")
	}
}
