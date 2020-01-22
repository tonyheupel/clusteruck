#!/usr/bin/env bash
docker run --name clusteruck -p 8081:8081 --cpus="1.0" --memory="100m" -d tchype/node-clusteruck
