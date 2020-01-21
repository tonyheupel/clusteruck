#!/usr/bin/env bash
docker build -t tchype/dotnetcore-clusteruck . &&
  docker save -o dotnetcore-clusteruck.tar tchype/dotnetcore-clusteruck
