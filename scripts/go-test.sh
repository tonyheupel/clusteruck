#!/usr/bin/env bash
echo "Removing previous container..."
./clean.sh

echo "Building Go image..."
pushd ../go/ && ./build.sh && ./run.sh && popd

echo "Starting Go test..."
./test.sh