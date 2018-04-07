package main

import (
	"log"
)

func ipcsLimits() map[string]interface{} {
	// output := "" //executeToString("/bin/sh", "-c", "ulimit "+flag)
	return nil
}

func parseIpcsLimits(output string) map[string]interface{} {
	res := map[string]interface{}{}
	log.Println(output)
	return res
}
