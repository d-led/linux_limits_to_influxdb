package main

import (
	"log"
)

func sanityCheck() {
	if v, e := ulimit("-a"); e == nil {
		log.Printf(v)
	} else {
		panic(e)
	}
}

func ulimits() map[string]interface{} {
	res := map[string]interface{}{}

	// -t: cpu time (seconds)             unlimited
	// -d: data seg size (kb)             unlimited
	// -s: stack size (kb)                8192
	// -c: core file size (blocks)        0
	// -m: resident set size (kb)         unlimited
	// -l: locked memory (kb)             64
	// -p: processes                      1048576
	// -n: file descriptors               1048576
	// -v: address space (kb)             unlimited
	// -w: locks                          unlimited
	// -e: scheduling priority            0
	// -r: real-time priority             0

	flags := map[string]string{
		"-w": "max_locks",
		"-n": "max_file_descriptors",
		"-p": "max_processes",
	}

	for flag, field := range flags {
		if val, err := ulimit(flag); err == nil {
			res[field] = normalizeToInt(val)
		}
	}

	return res
}

func ulimit(flag string) (string, error) {
	return executeToString("/bin/sh", "-c", "ulimit "+flag)
}
