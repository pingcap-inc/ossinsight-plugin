#!/bin/bash

HOST="192.168.31.94"
VOLUME_NAME="etcd-data"
REGISTRY="gcr.io/etcd-development/etcd"

# volume
docker volume rm ${VOLUME_NAME}
docker volume create --name ${VOLUME_NAME}

# container

docker stop etcd
docker rm etcd
docker run -it -d \
  -p 2379:2379 \
  -p 2380:2380 \
  --volume=${VOLUME_NAME}:/etcd-data \
  --name etcd ${REGISTRY}:latest \
  /usr/local/bin/etcd \
  --data-dir=/etcd-data --name node1 \
  --initial-advertise-peer-urls http://${HOST}:2380 --listen-peer-urls http://0.0.0.0:2380 \
  --advertise-client-urls http://${HOST}:2379 --listen-client-urls http://0.0.0.0:2379 \
  --initial-cluster node1=http://${HOST}:2380
