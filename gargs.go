package gargs

import (
	"os"
	"strings"
)

var Args []string

var ArgsMap = make(map[string]string)
var FlagMap = make(map[string]FlagType)

func Contains(val string) (bool, int) {
	for i, arg := range Args {
		if arg == val {
			return true, i
		}
	}

	return false, -1
}

// returns whether or not arg next to "val" is equal to "following"
func Follows(val string, following string) bool {
	if contains, index := Contains(val); contains {
		return Args[index+1] == following
	} else {
		return false
	}
}

// returns arg next to "val"
func Next(val string) (string, bool) {
	if contains, index := Contains(val); contains {
		if index+1 >= len(Args) {
			return "", false
		} else {
			return Args[index+1], true
		}
	} else {
		return "", false
	}
}

// if arg contains key and value in string, extract value
func ValueOf(val string) (string, bool) {
	if value, ok := ArgsMap[val]; ok {
		return value, true
	} else {
		return "", false
	}
}

type FlagType int

const (
	long  FlagType = iota // a flag like "--post-data"
	short                 // a flag like "-h"
)

func Flag(val string) (exists bool, flagType FlagType) {
	flagVal, ok := FlagMap[val]

	exists = ok

	if !ok {
		flagType = -1
	} else {
		flagType = flagVal
	}

	return
}

func ParseArgs() {
	Args = make([]string, len(os.Args[1:]))
	copy(Args, os.Args[1:])
	parseArgKeys()
}

func isFlag(val string) (bool, FlagType) {
	if len(val) > 1 {
		if val[0] == '-' {
			if val[1] == '-' {
				return true, long
			} else {
				return true, short
			}
		}
	}

	return false, -1
}

func parseArgKeys() {
	for i, arg := range Args {
		var f FlagType
		if isFlag, flagType := isFlag(arg); isFlag {
			if flagType == long {
				Args[i] = arg[2:]
			} else {
				Args[i] = arg[1:]
			}

			f = flagType
		} else {
			f = -1
		}

		key, value := splitVal(Args[i])

		if value != "" {
			Args[i] = key
			ArgsMap[key] = value
		}

		if f != -1 {
			FlagMap[key] = f
		}
	}
}

func splitVal(arg string) (key string, value string) {
	if strings.Contains(arg, "=") {
		split := strings.Index(arg, "=")
		key = arg[0:split]
		value = arg[split+1:]
		return
	} else {
		return arg, ""
	}
}
