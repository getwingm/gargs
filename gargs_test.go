package gargs

import (
	"os"
	"testing"
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

func TestFlag(t *testing.T) {
	isFlag, flagType := Flag("arg4")

	if !isFlag {
		t.Error("expected arg4 to be flag", isFlag)
	}

	if flagType != long {
		t.Error("expected arg4 to be of long flag type", flagType)
	}
}

func TestFlag2(t *testing.T) {
	isFlag, flagType := Flag("arg5")

	if !isFlag {
		t.Error("expected arg5 to be flag", isFlag)
	}

	if flagType != short {
		t.Error("expected arg5 to be of short flag type", flagType)
	}
}

func TestFlag3(t *testing.T) {
	isFlag, flagType := Flag("arg2")

	if isFlag {
		t.Error("expected arg2 to not be a flag")
	}

	if flagType != -1 {
		t.Error("expected flag type to be -1")
	}
}

func TestFlag4(t *testing.T) {
	isFlag, flagType := Flag("arg77")

	if isFlag {
		t.Error("expected arg77 to not be a flag")
	}

	if flagType != -1 {
		t.Error("expected flag type to be -1")
	}
}

func TestFlagArgMix(t *testing.T) {
	isFlag, flagType := Flag("arg-six")

	if !isFlag {
		t.Error("expected arg-six to be a flag", FlagMap)
	}

	if flagType != long {
		t.Error("flag type expected to be long", flagType)
	}

	value, ok := ValueOf("arg-six")

	if !ok {
		t.Error("expected value of arg-six to be okay")
	}

	if value != "VALUE" {
		t.Error("expected value to be 'VALUE'")
	}

}

func TestMain(m *testing.M) {
	initialize()
	retCode := m.Run()
	os.Exit(retCode)
}

func initialize() {
	Args = []string{"arg1", "arg2", "arg3=value", "--arg4", "-arg5", "--arg-six=VALUE"}
	parseArgKeys()
}
