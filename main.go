package main

import (
	"log"
	"os"
	"time"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	forever := len(os.Args) == 1
	log.Printf("Running forever: %v", forever)

	db, llti := defaultConfig()

	if forever {
		log.Printf("Delay: %vs", llti.DelaySeconds)
	}

	// call with more than 0 parameter to run only once
	run(forever, db, llti)
}

func run(forever bool, config *influxConfig, lconfig *lltiConfig) {
	log.Printf("Connecting to %v, db: %v", config.InfluxUrl, config.InfluxDb)

	sanityCheck()

	client, err := newInfluxClient(config)
	if err != nil {
		panic(err)
	}

	ensureDbExists(client, config.InfluxDb)

	for {
		fields := ulimits()
		mergeIntoFirst(fields, IpcsLimits())
		tags := tags()

		insertSinglePointNow(client, config, "limits", fields, tags)

		time.Sleep(time.Duration(lconfig.DelaySeconds) * time.Second)

		if !forever {
			break
		}
	}
}

type influxConfig struct {
	InfluxUrl  string
	InfluxDb   string
	InfluxUser string
	InfluxPass string
}

type lltiConfig struct {
	DelaySeconds int
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
		DelaySeconds: int(toIntOrDefault(os.Getenv("LLTI_DELAY_SECONDS"), 3)), // no range check
	}

	return config, lconfig
}
