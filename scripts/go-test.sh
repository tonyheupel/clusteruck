#!/usr/bin/env bash
echo "Removing previous container..."
pushd ../go/ && ./clean.sh

echo "Building Go image..."
./build.sh && ./run.sh && popd

echo "Starting Go test..."
ab -n 1000 -c 10 http://localhost:8081/