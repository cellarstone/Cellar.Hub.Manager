#!/bin/bash

# FLUENTD ---------------------------------------------
docker service rm fluentd

# ELASTICSEARCH ------------------------------------------
docker service rm elasticsearch

# KIBANA --------------------------------------------------
docker service rm kibana

# MONGODB --------------------------------------------------
docker service rm mongodb

# MQTT --------------------------------------------------
docker service rm mqtt

# PROMETHEUS --------------------------------------------------
docker service rm prometheus

# PROMETHEUS gateway --------------------------------------------------
docker service rm pushgateway

# GRAFANA --------------------------------------------------
docker service rm grafana

# NET-DATA --------------------------------------------------
docker service rm sysmon

# TICK-STACK --------------------------------------------------

# TELEGRAF ---
docker service rm telegraf

# INFLUXDB ---
docker service rm influxdb

# KAPACITOR ---
docker service rm kapacitor

# CHRONOGRAF ---
docker service rm chronograf


# HUB CORE - WEB --------------------------------------------------
docker service rm cellar-hub-core-web

# HUB CORE - ADMIN --------------------------------------------------
docker service rm cellar-hub-core-admin

# HUB CORE - FILE --------------------------------------------------
docker service rm cellar-hub-core-file

# HUB CORE - WEBSOCKETS --------------------------------------------------
docker service rm cellar-hub-core-websockets


# HUB CORE - IOT --------------------------------------------------
docker service rm cellar-hub-core-iot

# HUB CORE - USER --------------------------------------------------
docker service rm cellar-hub-core-user

# HUB CORE - WORKFLOW --------------------------------------------------
docker service rm cellar-hub-core-workflow



# HUB MODULE - OFFICE API --------------------------------------------------
docker service rm cellar-hub-module-office-api

# HUB MODULE - OFFICE METTINGROOMS --------------------------------------------------
docker service rm cellar-hub-module-office-meetingrooms

# HUB MODULE - OFFICE RECEPTION --------------------------------------------------
docker service rm cellar-hub-module-office-reception


# HUB PROXY --------------------------------------------------
docker service rm cellar-hub-proxy


temp=`docker ps --no-trunc | grep -c 'cellar'`
while [ $temp -gt 0 ]
do
    echo "wait for container"
    sleep 1
    temp=`docker ps --no-trunc | grep -c 'cellar'`
done



# -------------------------------------
# DOCKER CLEANING ALL
# -------------------------------------
docker system prune --force
