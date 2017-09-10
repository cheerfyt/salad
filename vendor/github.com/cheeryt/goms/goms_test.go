package goms

import (
	"testing"
)

func TestMs_1(t *testing.T) {
	if ms, _ := Ms("1s"); ms != 1000 {
		t.Error("1s test error")
	} else {
		t.Log("1s test pass")
	}
}

func TestMs_2(t *testing.T) {
	if ms, _ := Ms("1m"); ms != 60000 {
		t.Error("1m test error")
	} else {
		t.Log("1m test pass")
	}
}

func TestMs_3(t *testing.T) {
	if ms, _ := Ms("1d"); ms != 86400000 {
		t.Error("1d test error")
	} else {
		t.Log("1d test pass")
	}
}

func TestMs_4(t *testing.T) {
	if ms, _ := Ms("1y"); ms != 31536000000 {
		t.Error("1y test error")
	} else {
		t.Log("1y test pass")
	}
}