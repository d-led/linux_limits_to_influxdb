package main

import (
	"regexp"
	"strings"
)

func IpcsLimits() map[string]interface{} {
	if output, err := executeToString("/bin/sh", "-c", "ipcs -ls"); err == nil {
		return parseIpcsLimits(output)
	}
	return map[string]interface{}{}
}

func parseIpcsLimits(output string) map[string]interface{} {
	res := map[string]interface{}{}
	valueRegex := regexp.MustCompile(`(?m)(\w.*?\w)\s*=\s*(\S+)`)
	matches := valueRegex.FindAllStringSubmatch(output, -1)
	for _, match := range matches {
		key := validIdentifierFrom(match[1])
		value := normalizeToInt(match[2])
		res[key] = value
	}
	return res
}

func validIdentifierFrom(key string) string {
	var spaces = regexp.MustCompile(`\s+`)
	res := spaces.ReplaceAllString(strings.TrimSpace(key), `_`)
	return res
}
