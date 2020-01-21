#!/usr/bin/env bash
echo "Removing previous container..."
./clean.sh

echo "Building Node image..."
pushd ../node && ./build.sh && ./run.sh && popd

echo "Starting Node test..."
./test.sh