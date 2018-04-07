package main

import (
	"os/exec"
	"strconv"
	"strings"
)

func mergeIntoFirst(first map[string]interface{}, second map[string]interface{}) {
	for k, v := range second {
		first[k] = v
	}
}

func mergeIntoFirstStringMap(first map[string]string, second map[string]string) {
	for k, v := range second {
		first[k] = v
	}
}

func executeToString(cmd string, args ...string) (string, error) {
	proc := exec.Command(cmd, args...)

	res, err := proc.Output()

	return string(res), err
}

func normalizeToInt(val string) int64 {
	val = strings.TrimSpace(val)

	if val == "unlimited" {
		return -1
	}

	if r, err := strconv.ParseInt(val, 10, 64); err == nil {
		return r
	}

	return -42
}

func toIntOrDefault(val string, defaultValue int64) int64 {
	if r, err := strconv.ParseInt(val, 10, 64); err == nil {
		return r
	}
	return defaultValue
}
