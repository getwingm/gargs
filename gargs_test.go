package gargs

import (
	"testing"
	"os"
)

func TestContains(t *testing.T) {
	contains, index := Contains("arg1")

	if !contains {
		t.Error("Expected Contains() to be true")
	}

	if index == -1 {
		t.Error("Expected Contains() to return non-negative value")
	}
}

func TestContains2(t *testing.T) {
	contains, index := Contains("arg77")

	if contains {
		t.Error("Expected Contains() to be false")
	}

	if index != -1 {
		t.Error("Expected Contains() to return negative value")
	}
}

func TestFollows(t *testing.T) {
	if !Follows("arg1", "arg2") {
		t.Error("Expected arg2 after arg1")
	}
}

func TestFollows2(t *testing.T) {
	if Follows("arg1", "args3") {
		t.Error("Expected Follows() to return false when args1 and args3 passed in")
	}
}

func TestNext(t *testing.T) {
	result, ok := Next("arg1")

	if !ok {
		t.Error("expected Next() to be ok")
	}

	if result != "arg2" {
		t.Error("expected arg2 to follow arg1")
	}
}

func TestValueOf(t *testing.T) {
	value, ok := ValueOf("arg3")

	if !ok {
		t.Error("epxected ValueOf() to be okay")
	}

	if value != "value" {
		t.Error("expected TestValueOf to return 'value'")
	}
}

func TestMain(m *testing.M) {
	initialize()
	retCode := m.Run()
	os.Exit(retCode)
}

func initialize() {
	Args = []string{"arg1", "arg2", "arg3=value"}
	parseArgKeys()
}
