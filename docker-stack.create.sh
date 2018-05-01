#!/bin/sh

# NETWORK ---------------------------------------------
# docker network rm cellarstone-net
IS_NETWORK_EXIST=`docker network ls | grep cellarstone-net`
if [ "$IS_NETWORK_EXIST" != "" ]; then
	echo "NETWORK EXIST!"
    echo $IS_NETWORK_EXIST

else
	echo "NETWORK DOESN'T EXIST!"
    echo $IS_NETWORK_EXIST

    docker network create --driver overlay cellarstone-net

fi


# FLUENTD ---------------------------------------------
# docker service rm fluentd
IS_FLUENTD_EXIST=`docker service ps fluentd`
if [ "$IS_FLUENTD_EXIST" != "" ]; then
	echo "EXIST!"
    echo $IS_FLUENTD_EXIST

    docker service update --image cellarstone/cellar.hub.fluentd:0.50.0 \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --with-registry-auth \
                      fluentd

else
	echo "DOESN'T EXIST!"
    echo $IS_FLUENTD_EXIST

    docker service create --name fluentd \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --network cellarstone-net \
                      --publish 24224:24224 \
                      --with-registry-auth \
                      cellarstone/cellar.hub.fluentd:0.50.0

fi

# ELASTICSEARCH ------------------------------------------
# docker service rm elasticsearch
IS_ELASTIC_EXIST=`docker service ps elasticsearch`
if [ "$IS_ELASTIC_EXIST" != "" ]; then
	echo "EXIST!"
    echo $IS_ELASTIC_EXIST

    docker service update --image elasticsearch \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      elasticsearch

else
	echo "DOESN'T EXIST!"
    echo $IS_ELASTIC_EXIST

    docker service create --name elasticsearch \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --network cellarstone-net \
                      --mount type=bind,source=/data/cellarstone.hub/core/elasticsearch,target=/var/lib/elasticsearch \
                      --publish 9200:9200 \
                      elasticsearch

fi

# KIBANA --------------------------------------------------
# docker service rm kibana
IS_KIBANA_EXIST=`docker service ps kibana`
if [ "$IS_KIBANA_EXIST" != "" ]; then
	echo "EXIST!"
    echo $IS_KIBANA_EXIST

    docker service update --image kibana \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      kibana

else
	echo "DOESN'T EXIST!"
    echo $IS_KIBANA_EXIST

    docker service create --name kibana \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --network cellarstone-net \
                      --publish 5601:5601 \
                      kibana

fi

# MONGODB --------------------------------------------------
# docker service rm mongodb
IS_MONGO_EXIST=`docker service ps mongodb`
if [ "$IS_MONGO_EXIST" != "" ]; then
	echo "EXIST!"
    echo $IS_MONGO_EXIST

    docker service update --image cellarstone/cellar.hub.mongodb:0.50.0 \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --with-registry-auth \
                      mongodb

else
	echo "DOESN'T EXIST!"
    echo $IS_MONGO_EXIST

    docker service create --name mongodb \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --network cellarstone-net \
                      --mount type=bind,source=/data/cellarstone.hub/core/mongodb,target=/data/db \
                      --publish 27017:27017 \
                      --with-registry-auth \
                      cellarstone/cellar.hub.mongodb:0.50.0

fi

# MQTT --------------------------------------------------
# docker service rm mqtt
IS_MQTT_EXIST=`docker service ps mqtt`
if [ "$IS_MQTT_EXIST" != "" ]; then
	echo "EXIST!"
    echo $IS_MQTT_EXIST

    docker service update --image toke/mosquitto \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --with-registry-auth \
                      mqtt

else
	echo "DOESN'T EXIST!"
    echo $IS_MQTT_EXIST

    docker service create --name mqtt \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --network cellarstone-net \
                      --publish 1883:1883 \
                      --with-registry-auth \
                      toke/mosquitto

fi

# PROMETHEUS --------------------------------------------------
# docker service rm prometheus
IS_PROMETHEUS_EXIST=`docker service ps prometheus`
if [ "$IS_PROMETHEUS_EXIST" != "" ]; then
	echo "EXIST!"
    echo $IS_PROMETHEUS_EXIST

    docker service update --image cellarstone/cellar.hub.prometheus:0.50.0 \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --with-registry-auth \
                      prometheus

else
	echo "DOESN'T EXIST!"
    echo $IS_PROMETHEUS_EXIST

    docker service create --name prometheus \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --network cellarstone-net \
                      --mount type=bind,source=/data/cellarstone.hub/core/prometheus,target=/data/prometheus \
                      --publish 9090:9090 \
                      --with-registry-auth \
                      cellarstone/cellar.hub.prometheus:0.50.0

fi

# PROMETHEUS gateway --------------------------------------------------
# docker service rm pushgateway
IS_PROMETHEUSGATEWAY_EXIST=`docker service ps pushgateway`
if [ "$IS_PROMETHEUSGATEWAY_EXIST" != "" ]; then
	echo "EXIST!"
    echo $IS_PROMETHEUSGATEWAY_EXIST

    docker service update --image prom/pushgateway \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      pushgateway

else
	echo "DOESN'T EXIST!"
    echo $IS_PROMETHEUSGATEWAY_EXIST

    docker service create --name pushgateway \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --network cellarstone-net \
                      --publish 9091:9091 \
                      prom/pushgateway

fi

# GRAFANA --------------------------------------------------
# docker service rm grafana
IS_GRAFANA_EXIST=`docker service ps grafana`
if [ "$IS_GRAFANA_EXIST" != "" ]; then
	echo "EXIST!"
    echo $IS_GRAFANA_EXIST

    docker service update --image grafana/grafana \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      pushgateway

else
	echo "DOESN'T EXIST!"
    echo $IS_GRAFANA_EXIST

    docker service create --name grafana \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --network cellarstone-net \
                      --mount type=bind,source=/data/cellarstone.hub/core/grafana,target=/var/lib/grafana \
                      --publish 3000:3000 \
                      grafana/grafana

fi

# NET-DATA --------------------------------------------------
# docker service rm sysmon
IS_SYSMON_EXIST=`docker service ps sysmon`
if [ "$IS_SYSMON_EXIST" != "" ]; then
	echo "EXIST!"
    echo $IS_SYSMON_EXIST

    docker service update --image titpetric/netdata \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      sysmon

else
	echo "DOESN'T EXIST!"
    echo $IS_SYSMON_EXIST

    docker service create --name sysmon \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --network cellarstone-net \
                      --mount type=bind,source=/proc,target=/host/proc:ro \
                      --mount type=bind,source=/sys,target=/host/sys:ro \
                      --publish 19999:19999 \
                      titpetric/netdata

fi

# TICK-STACK --------------------------------------------------

# TELEGRAF ---
# docker service rm telegraf
IS_TELEGRAF_EXIST=`docker service ps telegraf`
if [ "$IS_TELEGRAF_EXIST" != "" ]; then
	echo "EXIST!"
    echo $IS_TELEGRAF_EXIST

    docker service update --image cellarstone/cellar.hub.telegraf:0.50.0 \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --with-registry-auth \
                      telegraf

else
	echo "DOESN'T EXIST!"
    echo $IS_TELEGRAF_EXIST

    docker service create --name telegraf \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --network cellarstone-net \
                      --publish 8094:8094 \
                      --publish 8092:8092/udp \
                      --publish 8125:8125/udp \
                      --with-registry-auth \
                      cellarstone/cellar.hub.telegraf:0.50.0

fi

# INFLUXDB ---
# docker service rm influxdb
IS_INFLUX_EXIST=`docker service ps influxdb`
if [ "$IS_INFLUX_EXIST" != "" ]; then
	echo "EXIST!"
    echo $IS_INFLUX_EXIST

    docker service update --image influxdb:1.3.5 \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      influxdb

else
	echo "DOESN'T EXIST!"
    echo $IS_INFLUX_EXIST

    docker service create --name influxdb \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --network cellarstone-net \
                      --mount type=bind,source=/data/cellarstone.hub/core/influxdb,target=/var/lib/influxdb \
                      --publish 8086:8086 \
                      influxdb:1.3.5

fi

# KAPACITOR ---
# docker service rm kapacitor
IS_KAPACITOR_EXIST=`docker service ps kapacitor`
if [ "$IS_KAPACITOR_EXIST" != "" ]; then
	echo "EXIST!"
    echo $IS_KAPACITOR_EXIST

    docker service update --image kapacitor:1.3.3 \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      kapacitor

else
	echo "DOESN'T EXIST!"
    echo $IS_KAPACITOR_EXIST

    docker service create --name kapacitor \
                      --env KAPACITOR_HOSTNAME=kapacitor \
                      --env KAPACITOR_INFLUXDB_0_URLS_0=http://influxdb:8086 \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --network cellarstone-net \
                      --publish 9092:9092 \
                      kapacitor:1.3.3

fi

# CHRONOGRAF ---
# docker service rm chronograf
IS_CHRONOGRAF_EXIST=`docker service ps chronograf`
if [ "$IS_CHRONOGRAF_EXIST" != "" ]; then
	echo "EXIST!"
    echo $IS_CHRONOGRAF_EXIST

    docker service update --image chronograf:1.3.8 \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      chronograf

else
	echo "DOESN'T EXIST!"
    echo $IS_CHRONOGRAF_EXIST

    docker service create --name chronograf \
                      --env INFLUXDB_URL=http://influxdb:8086 \
                      --env KAPACITOR_URL=http://kapacitor:9092 \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --network cellarstone-net \
                      --publish 8888:8888 \
                      chronograf:1.3.8

fi



# HUB CORE - WEB --------------------------------------------------

# docker service rm cellar-hub-core-web
IS_HUBCOREWEB_EXIST=`docker service ps cellar-hub-core-web`
if [ "$IS_HUBCOREWEB_EXIST" != "" ]; then
	echo "EXIST!"
    echo $IS_HUBCOREWEB_EXIST

    docker service update --image cellarstone/cellar.hub.core.web:0.50.0 \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --log-driver fluentd \
                      --log-opt mode=non-blocking \
                      --log-opt tag="docker.cellar.hub.core.web" \
                      --with-registry-auth \
                      cellar-hub-core-web

else
	echo "DOESN'T EXIST!"
    echo $IS_HUBCOREWEB_EXIST

    docker service create --name cellar-hub-core-web \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --network cellarstone-net \
                      --label traefik.enable=true \
                      --label traefik.port=44401 \
                      --label traefik.docker.network=cellarstone-net \
                      --label traefik.backend=core-web \
                      --label traefik.frontend.rule=Host:web.cellarstone.hub \
                      --log-driver fluentd \
                      --log-opt mode=non-blocking \
                      --log-opt tag="docker.cellar.hub.core.web" \
                      --with-registry-auth \
                      --publish 44401:44401 \
                      cellarstone/cellar.hub.core.web:0.50.0     

fi

# HUB CORE - ADMIN --------------------------------------------------

# docker service rm cellar-hub-core-admin
IS_HUBCOREADMIN_EXIST=`docker service ps cellar-hub-core-admin`
if [ "$IS_HUBCOREADMIN_EXIST" != "" ]; then
	echo "EXIST!"
    echo $IS_HUBCOREADMIN_EXIST

    docker service update --image cellarstone/cellar.hub.core.admin:0.50.0 \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --log-driver fluentd \
                      --log-opt mode=non-blocking \
                      --log-opt tag="docker.cellar.hub.core.admin" \
                      --with-registry-auth \
                      cellar-hub-core-admin

else
	echo "DOESN'T EXIST!"
    echo $IS_HUBCOREADMIN_EXIST

    docker service create --name cellar-hub-core-admin \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --network cellarstone-net \
                      --label traefik.enable=true \
                      --label traefik.port=44402 \
                      --label traefik.docker.network=cellarstone-net \
                      --label traefik.backend=core-admin \
                      --label traefik.frontend.rule=Host:admin.cellarstone.hub \
                      --log-driver fluentd \
                      --log-opt mode=non-blocking \
                      --log-opt tag="docker.cellar.hub.core.admin" \
                      --with-registry-auth \
                      --publish 44402:44402 \
                      cellarstone/cellar.hub.core.admin:0.50.0     

fi

# HUB CORE - FILE --------------------------------------------------

# docker service rm cellar-hub-core-file
IS_HUBCOREFILE_EXIST=`docker service ps cellar-hub-core-file`
if [ "$IS_HUBCOREFILE_EXIST" != "" ]; then
	echo "EXIST!"
    echo $IS_HUBCOREFILE_EXIST

    docker service update --image cellarstone/cellar.hub.core.file:0.50.0 \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --log-driver fluentd \
                      --log-opt mode=non-blocking \
                      --log-opt tag="docker.cellar.hub.core.file" \
                      --with-registry-auth \
                      cellar-hub-core-file

else
	echo "DOESN'T EXIST!"
    echo $IS_HUBCOREFILE_EXIST

    docker service create --name cellar-hub-core-file \
                      --env PORT=44404 \
                      --env DIRECTORY=/app/data \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --network cellarstone-net \
                      --label traefik.enable=true \
                      --label traefik.port=44404 \
                      --label traefik.docker.network=cellarstone-net \
                      --label traefik.backend=core-file \
                      --label traefik.frontend.rule=Host:file.cellarstone.hub \
                      --log-driver fluentd \
                      --log-opt mode=non-blocking \
                      --log-opt tag="docker.cellar.hub.core.file" \
                      --mount type=bind,source=/data/cellarstone.hub/core/file,target=/app/data \
                      --publish 44404:44404 \
                      --with-registry-auth \
                      cellarstone/cellar.hub.core.file:0.50.0     

fi


# HUB CORE - WEBSOCKETS --------------------------------------------------

# docker service rm cellar-hub-core-websockets
IS_HUBCOREWEBSOCKETS_EXIST=`docker service ps cellar-hub-core-websockets`
if [ "$IS_HUBCOREWEBSOCKETS_EXIST" != "" ]; then
	echo "EXIST!"
    echo $IS_HUBCOREWEBSOCKETS_EXIST

    docker service update --image cellarstone/cellar.hub.core.websockets:0.50.0 \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --log-driver fluentd \
                      --log-opt mode=non-blocking \
                      --log-opt tag="docker.cellar.hub.core.websockets" \
                      --with-registry-auth \
                      cellar-hub-core-websockets

else
	echo "DOESN'T EXIST!"
    echo $IS_HUBCOREWEBSOCKETS_EXIST

    docker service create --name cellar-hub-core-websockets \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --network cellarstone-net \
                      --label traefik.enable=true \
                      --label traefik.port=44406 \
                      --label traefik.docker.network=cellarstone-net \
                      --label traefik.backend=core-websockets \
                      --label traefik.frontend.rule=Host:websockets.cellarstone.hub \
                      --log-driver fluentd \
                      --log-opt mode=non-blocking \
                      --log-opt tag="docker.cellar.hub.core.websockets" \
                      --publish 44406:44406 \
                      --with-registry-auth \
                      cellarstone/cellar.hub.core.websockets:0.50.0     

fi





# HUB CORE - IOT --------------------------------------------------

# docker service rm cellar-hub-core-iot
IS_HUBCOREIOT_EXIST=`docker service ps cellar-hub-core-iot`
if [ "$IS_HUBCOREIOT_EXIST" != "" ]; then
	echo "EXIST!"
    echo $IS_HUBCOREIOT_EXIST

    docker service update --image cellarstone/cellar.hub.core.iot:0.50.0 \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --log-driver fluentd \
                      --log-opt mode=non-blocking \
                      --log-opt tag="docker.cellar.hub.core.iot" \
                      --with-registry-auth \
                      cellar-hub-core-iot

else
	echo "DOESN'T EXIST!"
    echo $IS_HUBCOREIOT_EXIST

    docker service create --name cellar-hub-core-iot \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --network cellarstone-net \
                      --env PORT=44403 \
                      --env MQTT_URL=mqtt \
                      --env MONGO_URL=mongodb \
                      --label traefik.enable=true \
                      --label traefik.port=44403 \
                      --label traefik.docker.network=cellarstone-net \
                      --label traefik.backend=core-iot \
                      --label traefik.frontend.rule=Host:iot.cellarstone.hub \
                      --log-driver fluentd \
                      --log-opt mode=non-blocking \
                      --log-opt tag="docker.cellar.hub.core.iot" \
                      --with-registry-auth \
                      --publish 44403:44403 \
                      cellarstone/cellar.hub.core.iot:0.50.0     

fi

# HUB CORE - USER --------------------------------------------------

# docker service rm cellar-hub-core-user
# IS_HUBCOREUSER_EXIST=`docker service ps cellar-hub-core-user`
# if [ "$IS_HUBCOREUSER_EXIST" != "" ]; then
# 	echo "EXIST!"
#     echo $IS_HUBCOREUSER_EXIST

#     docker service update --image cellarstone/cellar.hub.core.user:0.50.0 \
#                       --replicas 1 \
#                       --update-parallelism 2 \
#                       --update-delay 5s \
#                       --update-order start-first \
#                       --restart-condition on-failure \
#                       --restart-delay 5s \
#                       --restart-max-attempts 3 \
#                       --restart-window 120s \
#                       --log-driver fluentd \
#                       --log-opt mode=non-blocking \
#                       --log-opt tag="docker.cellar.hub.core.user" \
#                       --with-registry-auth \
#                       cellar-hub-core-user

# else
# 	echo "DOESN'T EXIST!"
#     echo $IS_HUBCOREUSER_EXIST

#     docker service create --name cellar-hub-core-user \
#                       --replicas 1 \
#                       --update-parallelism 2 \
#                       --update-delay 5s \
#                       --update-order start-first \
#                       --restart-condition on-failure \
#                       --restart-delay 5s \
#                       --restart-max-attempts 3 \
#                       --restart-window 120s \
#                       --network cellarstone-net \
#                       --env PORT=44403 \
#                       --env MQTT_URL=mqtt \
#                       --env MONGO_URL=mongodb \
#                       --label traefik.enable=true \
#                       --label traefik.port=44407 \
#                       --label traefik.docker.network=cellarstone-net \
#                       --label traefik.backend=core-user \
#                       --label traefik.frontend.rule=Host:user.cellarstone.hub \
#                       --log-driver fluentd \
#                       --log-opt mode=non-blocking \
#                       --log-opt tag="docker.cellar.hub.core.user" \
#                       --with-registry-auth \
#                       --publish 44407:44407 \
#                       cellarstone/cellar.hub.core.user:0.50.0     

# fi

# HUB CORE - WORKFLOW --------------------------------------------------

# docker service rm cellar-hub-core-workflow
IS_HUBCOREWORKFLOW_EXIST=`docker service ps cellar-hub-core-workflow`
if [ "$IS_HUBCOREWORKFLOW_EXIST" != "" ]; then
	echo "EXIST!"
    echo $IS_HUBCOREWORKFLOW_EXIST

    docker service update --image cellarstone/cellar.hub.core.workflow:0.50.0 \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --log-driver fluentd \
                      --log-opt mode=non-blocking \
                      --log-opt tag="docker.cellar.hub.core.workflow" \
                      --with-registry-auth \
                      cellar-hub-core-workflow

else
	echo "DOESN'T EXIST!"
    echo $IS_HUBCOREWORKFLOW_EXIST

    docker service create --name cellar-hub-core-workflow \
                      --env PORT=44405 \
                      --env MONGO_URL=mongodb \
                      --env MQTT_URL=mqtt \
                      --env INFLUX_URL=http://influxdb:8086 \
                      --env WEBSOCKETS_URL=cellar-hub-core-websockets:44406 \
                      --env CELLAR_API_URL=cellar-hub-core-api:44413 \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --network cellarstone-net \
                      --label traefik.enable=true \
                      --label traefik.port=44405 \
                      --label traefik.docker.network=cellarstone-net \
                      --label traefik.backend=core-workflow \
                      --label traefik.frontend.rule=Host:workflow.cellarstone.hub \
                      --log-driver fluentd \
                      --log-opt mode=non-blocking \
                      --log-opt tag="docker.cellar.hub.core.workflow" \
                      --publish 44405:44405 \
                      --with-registry-auth \
                      cellarstone/cellar.hub.core.workflow:0.50.0     

fi

# HUB MODULE - OFFICE API --------------------------------------------------

# docker service rm cellar-hub-module-office-api
IS_HUBMODULE_OFFICEAPI_EXIST=`docker service ps cellar-hub-module-office-api`
if [ "$IS_HUBMODULE_OFFICEAPI_EXIST" != "" ]; then
	echo "EXIST!"
    echo $IS_HUBMODULE_OFFICEAPI_EXIST

    docker service update --image cellarstone/cellar.hub.module.office.api:0.50.0 \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --log-driver fluentd \
                      --log-opt mode=non-blocking \
                      --log-opt tag="docker.cellar-hub-module-office-api" \
                      --with-registry-auth \
                      cellar-hub-module-office-api

else
	echo "DOESN'T EXIST!"
    echo $IS_HUBMODULE_OFFICEAPI_EXIST

    docker service create --name cellar-hub-module-office-api \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --network cellarstone-net \
                      --label traefik.enable=true \
                      --label traefik.port=44513 \
                      --label traefik.docker.network=cellarstone-net \
                      --label traefik.backend=office-api \
                      --label traefik.frontend.rule=Host:officeapi.cellarstone.hub \
                      --log-driver fluentd \
                      --log-opt mode=non-blocking \
                      --log-opt tag="docker.cellar-hub-module-office-api" \
                      --publish 44513:44513 \
                      --with-registry-auth \
                      cellarstone/cellar.hub.module.office.api:0.50.0

fi

# HUB MODULE - OFFICE METTINGROOMS --------------------------------------------------

# docker service rm cellar-hub-module-office-meetingrooms
IS_HUBMODULE_OFFICEMEETINGS_EXIST=`docker service ps cellar-hub-module-office-meetingrooms`
if [ "$IS_HUBMODULE_OFFICEMEETINGS_EXIST" != "" ]; then
	echo "EXIST!"
    echo $IS_HUBMODULE_OFFICEMEETINGS_EXIST

    docker service update --image cellarstone/cellar.hub.module.office.meetingrooms:0.50.0 \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --log-driver fluentd \
                      --log-opt mode=non-blocking \
                      --log-opt tag="docker.cellar-hub-module-office-meetingrooms" \
                      --with-registry-auth \
                      cellar-hub-module-office-meetingrooms

else
	echo "DOESN'T EXIST!"
    echo $IS_HUBMODULE_OFFICEMEETINGS_EXIST

    docker service create --name cellar-hub-module-office-meetingrooms \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --network cellarstone-net \
                      --label traefik.enable=true \
                      --label traefik.port=44511 \
                      --label traefik.docker.network=cellarstone-net \
                      --label traefik.backend=office-meetingrooms \
                      --label traefik.frontend.rule=Host:meetingrooms.cellarstone.hub \
                      --log-driver fluentd \
                      --log-opt mode=non-blocking \
                      --log-opt tag="docker.cellar-hub-module-office-meetingrooms" \
                      --publish 44511:44511 \
                      --with-registry-auth \
                      cellarstone/cellar.hub.module.office.meetingrooms:0.50.0

fi

# HUB MODULE - OFFICE RECEPTION --------------------------------------------------

# docker service rm cellar-hub-module-office-reception
IS_HUBMODULE_OFFICERECEPTION_EXIST=`docker service ps cellar-hub-module-office-reception`
if [ "$IS_HUBMODULE_OFFICERECEPTION_EXIST" != "" ]; then
	echo "EXIST!"
    echo $IS_HUBMODULE_OFFICERECEPTION_EXIST

    docker service update --image cellarstone/cellar.hub.module.office.reception:0.50.0 \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --log-driver fluentd \
                      --log-opt mode=non-blocking \
                      --log-opt tag="docker.cellar-hub-module-office-reception" \
                      --with-registry-auth \
                      cellar-hub-module-office-reception

else
	echo "DOESN'T EXIST!"
    echo $IS_HUBMODULE_OFFICERECEPTION_EXIST

    docker service create --name cellar-hub-module-office-reception \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --network cellarstone-net \
                      --label traefik.enable=true \
                      --label traefik.port=44512 \
                      --label traefik.docker.network=cellarstone-net \
                      --label traefik.backend=office-reception \
                      --label traefik.frontend.rule=Host:reception.cellarstone.hub \
                      --log-driver fluentd \
                      --log-opt mode=non-blocking \
                      --log-opt tag="docker.cellar-hub-module-office-reception" \
                      --publish 44512:44512 \
                      --with-registry-auth \
                      cellarstone/cellar.hub.module.office.reception:0.50.0

fi

# HUB MODULE - OFFICE CAFE --------------------------------------------------

# docker service rm cellar-hub-module-office-cafe
IS_HUBMODULE_OFFICECAFE_EXIST=`docker service ps cellar-hub-module-office-cafe`
if [ "$IS_HUBMODULE_OFFICECAFE_EXIST" != "" ]; then
	echo "EXIST!"
    echo $IS_HUBMODULE_OFFICECAFE_EXIST

    docker service update --image cellarstone/cellar.hub.module.office.cafe:0.50.0 \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --log-driver fluentd \
                      --log-opt mode=non-blocking \
                      --log-opt tag="docker.cellar-hub-module-office-cafe" \
                      --with-registry-auth \
                      cellar-hub-module-office-cafe

else
	echo "DOESN'T EXIST!"
    echo $IS_HUBMODULE_OFFICECAFE_EXIST

    docker service create --name cellar-hub-module-office-cafe \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --network cellarstone-net \
                      --label traefik.enable=true \
                      --label traefik.port=44514 \
                      --label traefik.docker.network=cellarstone-net \
                      --label traefik.backend=office-cafe \
                      --label traefik.frontend.rule=Host:cafe.cellarstone.hub \
                      --log-driver fluentd \
                      --log-opt mode=non-blocking \
                      --log-opt tag="docker.cellar-hub-module-office-cafe" \
                      --publish 44514:44514 \
                      --with-registry-auth \
                      cellarstone/cellar.hub.module.office.cafe:0.50.0

fi


# HUB MODULE - OFFICE WELCOME --------------------------------------------------

# docker service rm cellar-hub-module-office-welcome
IS_HUBMODULE_OFFICEWELCOME_EXIST=`docker service ps cellar-hub-module-office-welcome`
if [ "$IS_HUBMODULE_OFFICEWELCOME_EXIST" != "" ]; then
	echo "EXIST!"
    echo $IS_HUBMODULE_OFFICEWELCOME_EXIST

    docker service update --image cellarstone/cellar.hub.module.office.welcome:0.50.0 \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --log-driver fluentd \
                      --log-opt mode=non-blocking \
                      --log-opt tag="docker.cellar-hub-module-office-welcome" \
                      --with-registry-auth \
                      cellar-hub-module-office-welcome

else
	echo "DOESN'T EXIST!"
    echo $IS_HUBMODULE_OFFICEWELCOME_EXIST

    docker service create --name cellar-hub-module-office-welcome \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --network cellarstone-net \
                      --label traefik.enable=true \
                      --label traefik.port=44515 \
                      --label traefik.docker.network=cellarstone-net \
                      --label traefik.backend=office-welcome \
                      --label traefik.frontend.rule=Host:welcome.cellarstone.hub \
                      --log-driver fluentd \
                      --log-opt mode=non-blocking \
                      --log-opt tag="docker.cellar-hub-module-office-welcome" \
                      --publish 44515:44515 \
                      --with-registry-auth \
                      cellarstone/cellar.hub.module.office.welcome:0.50.0

fi




# HUB PROXY --------------------------------------------------

# docker service rm cellar-hub-proxy
IS_HUBPROXY_EXIST=`docker service ps cellar-hub-proxy`
if [ "$IS_HUBPROXY_EXIST" != "" ]; then
	echo "EXIST!"
    echo $IS_HUBPROXY_EXIST

    docker service update --image cellarstone/cellar.hub.proxy:0.50.0 \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --log-driver fluentd \
                      --log-opt mode=non-blocking \
                      --log-opt tag="docker.cellar.hub.proxy" \
                      --with-registry-auth \
                      cellar-hub-proxy

else
	echo "DOESN'T EXIST!"
    echo $IS_HUBPROXY_EXIST

    docker service create --name cellar-hub-proxy \
                      --constraint=node.role==manager \
                      --mount type=bind,source=/var/run/docker.sock,target=/var/run/docker.sock \
                      --replicas 1 \
                      --update-parallelism 2 \
                      --update-delay 5s \
                      --update-order start-first \
                      --restart-condition on-failure \
                      --restart-delay 5s \
                      --restart-max-attempts 3 \
                      --restart-window 120s \
                      --network cellarstone-net \
                      --log-driver fluentd \
                      --log-opt mode=non-blocking \
                      --log-opt tag="docker.cellar.hub.proxy" \
                      --publish 80:80 \
                      --publish 8080:8080 \
                      --with-registry-auth \
                      cellarstone/cellar.hub.proxy:0.50.0 \
                      --docker \
                      --docker.swarmmode \
                      --docker.domain=cellar.hub \
                      --docker.watch \
                      --api \
                      --web
fi