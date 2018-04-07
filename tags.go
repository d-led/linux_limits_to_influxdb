package main

import (
	"regexp"
	"strings"
)

func tags() map[string]string {
	res := map[string]string{}
	if output := executeResultOrEmpty("id -un"); output != "" {
		res["user"] = output
	}
	if output := executeResultOrEmpty("hostname"); output != "" {
		res["hostname"] = output
	}
	mergeIntoFirstStringMap(res, linuxRelease())
	return res
}

func executeResultOrEmpty(cmd string) string {
	if output, err := executeToString("/bin/sh", "-c", cmd); err == nil {
		return strings.TrimSpace(output)
	}
	return ""
}

func linuxRelease() map[string]string {
	res := map[string]string{}
	kv := parseRelease(executeResultOrEmpty("cat /etc/*-release"))
	for k, v := range kv {
		switch k {
		case "ID":
			res["distro_key"] = v
		case "VERSION_ID":
			res["distro_version"] = v
		}

	}
	return res
}

//output := executeResultOrEmpty("cat /etc/*-release")

func parseRelease(contents string) map[string]string {
	res := map[string]string{}

	valueRegex := regexp.MustCompile(`(?m)(\w.*?\w)\s*=\s*(.+)`)
	matches := valueRegex.FindAllStringSubmatch(contents, -1)
	for _, match := range matches {
		key := validIdentifierFrom(match[1])
		value := strings.Trim(strings.TrimSpace(match[2]), "\"")
		res[key] = value
	}
	return res
}
