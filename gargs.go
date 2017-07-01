package gargs

import (
	"os"
	"strings"
)

var Args []string

var ArgsMap = make(map[string]string)

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
		return Args[index + 1] == following
	} else {
		return false
	}
}

// returns arg next to "val"
func Next(val string) (string, bool) {
	if contains, index := Contains(val); contains {
		return Args[index + 1], true
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

func init() {
	copy(Args, os.Args[1:])
	parseArgKeys()
}

func parseArgKeys() {
	for i, arg := range Args {
		if strings.Contains(arg, "=") {
			split := strings.Index(arg, "=")
			Args[i] = arg[0:split]
			ArgsMap[Args[i]] = arg[split+1:len(arg)]
		}
	}
}