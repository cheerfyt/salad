package swgo

import (
	"os"
	"testing"
)

func TestSave(t *testing.T) {
	save("swgo:*")
	if os.Getenv("SWGO") == "swgo:*" {
		t.Log("save test Ok")
	} else {
		t.Error("save test fail")
	}
}

func TestLoad(t *testing.T) {
	save("swgo:app")
	if load() == "swgo:app" {
		t.Log("load test pass")
	} else {
		t.Error("load test fail")
	}
}

func TestEnableSkip(t *testing.T) {
	enable("swgo:-")
	if len(skips) > 0 {
		t.Log("enable skips pass")
	} else {
		t.Error("enable skips fail")
	}
}

func TestEnableNames(t *testing.T) {
	enable("swgo:*")
	if len(names) > 0 {
		t.Log("enable names pass")
	} else {
		t.Error("enable names fail")
	}
}

func TestIngnoreCase(t *testing.T) {
	if ignoreCase("*") == "(?i)*" {
		t.Log("ignoreCase test pass")
	} else {
		t.Error("ignoreCase test fail")
	}
}
