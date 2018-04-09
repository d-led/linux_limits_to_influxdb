#!/bin/bash

set -e

docker-compose up -d --build influxdb main main-debian

echo "sleeping a bit ..."
sleep 20

docker-compose exec influxdb influx -execute "show databases"
docker-compose exec influxdb influx -precision rfc3339 -execute "select * from llti..limits order by time desc limit 20"

docker-compose down --remove-orphans
