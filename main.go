package main

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
	// "github.com/influxdata/influxdb/client/v2"
)

type influxConfig struct {
	InfluxUrl  string
	InfluxDb   string
	InfluxUser string
	InfluxPass string
}

type lltiConfig struct {
	DelaySeconds int
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	runForever(defaultConfig())
}

func runForever(config *influxConfig, lconfig *lltiConfig) {
	log.Printf("Connecting to %v, db: %v", config.InfluxUrl, config.InfluxDb)

	if v, e := ulimit("-a"); e == nil {
		log.Printf(v)
	} else {
		panic(e)
	}

	for {
		values := ulimits()
		mergeIntoFirst(values, IpcsLimits())
		tags := tags()

		log.Println(values)
		log.Println(tags)

		time.Sleep(time.Duration(lconfig.DelaySeconds) * time.Second)
	}
}

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

func defaultConfig() (*influxConfig, *lltiConfig) {
	config := &influxConfig{
		InfluxUrl:  os.Getenv("INFLUX_URL"),
		InfluxDb:   os.Getenv("INFLUX_DB"),
		InfluxUser: os.Getenv("INFLUX_USER"),
		InfluxPass: os.Getenv("INFLUX_PWD"),
	}

	if config.InfluxUrl == "" {
		config.InfluxUrl = "http://localhost:8086"
	}

	if config.InfluxDb == "" {
		config.InfluxDb = "llti"
	}

	lconfig := &lltiConfig{
		DelaySeconds: int(toIntOrDefault(os.Getenv("LLTI_DELAY"), 3)), // no range check
	}

	return config, lconfig
}

func executeToString(cmd string, args ...string) (string, error) {
	proc := exec.Command(cmd, args...)

	res, err := proc.Output()

	return string(res), err
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
