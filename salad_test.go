package salad

import (
	"os"
	"testing"
)

func TestSave(t *testing.T) {
	save("salad:*")

	if os.Getenv("SALAD") == "salad:*" {
		t.Log("save test Ok")
	} else {
		t.Error("save test fail")
	}
}

func TestLoad(t *testing.T) {
	save("salad:app")
	if load() == "salad:app" {
		t.Log("load test pass")
	} else {
		t.Error("load test fail")
	}
}

func TestEnableSkip(t *testing.T) {
	enable("salad:-")
	if len(skips) > 0 {
		t.Log("enable skips pass")
	} else {
		t.Error("enable skips fail")
	}
}

func TestEnableNames(t *testing.T) {
	enable("salad:*")
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
