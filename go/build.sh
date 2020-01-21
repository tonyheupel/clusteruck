#!/usr/bin/env bash
docker build -t tchype/go-clusteruck . &&
  docker save -o go-clusteruck.tar tchype/go-clusteruck
