package main

import (
	influx "github.com/influxdata/influxdb/client/v2"
	"log"
	"regexp"
	"time"
)

func newInfluxClient(config *influxConfig) (influx.Client, error) {
	c, err := influx.NewHTTPClient(influx.HTTPConfig{
		Addr:     config.InfluxUrl,
		Username: config.InfluxUser,
		Password: config.InfluxPass,
	})

	return c, err
}

func query(client influx.Client, db string, qs string) (res []influx.Result, err error) {
	q := influx.Query{
		Command:  qs,
		Database: db,
	}

	if response, err := client.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}

	return res, nil
}

func databaseExists(client influx.Client, db string) (bool, error) {
	minimalNameSanityCheck(db)

	res, err := query(client, db, "show databases")

	if err == nil {
		dbs := res[0].Series[0].Values
		for _, dbName := range dbs {
			if dbName[0] == db {
				return true, nil
			}
		}
	}
	return false, err
}

func tryCreateDb(client influx.Client, db string) {
	// ignore return values for now
	minimalNameSanityCheck(db)
	log.Println(query(client, db, "CREATE DATABASE "+db))
}

func minimalNameSanityCheck(name string) {
	if len(name) > 120 {
		panic("DB name too long: " + name)
	}

	nameRegex := regexp.MustCompile(`^\S+$`)
	if !nameRegex.MatchString(name) {
		panic("DB name suspicious: " + name)
	}
}

func ensureDbExists(client influx.Client, db string) {
	log.Printf("Ensuring, database '%v' exists\n", db)
	exists, err := databaseExists(client, db)

	// panic if couldn't query
	if err != nil {
		panic(err)
	}

	if exists {
		return
	}

	tryCreateDb(client, db)

	exists, err = databaseExists(client, db)
	if !exists {
		panic("Couldn't find or create db '" + db + "'")
	}
	if err != nil {
		panic(err)
	}
}

func insertSinglePointNow(client influx.Client, config *influxConfig, limits string, fields map[string]interface{}, tags map[string]string) {
	bp, err := influx.NewBatchPoints(influx.BatchPointsConfig{
		Database:  config.InfluxDb,
		Precision: "us",
	})

	if err != nil {
		panic(err)
	}

	pt, err := influx.NewPoint(limits, tags, fields, time.Now())

	if err != nil {
		panic(err.Error())
	}

	bp.AddPoint(pt)

	client.Write(bp)
}
