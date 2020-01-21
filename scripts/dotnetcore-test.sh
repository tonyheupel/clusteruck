#!/usr/bin/env bash
echo "Removing previous container..."
pushd ../dotnetcore/ && ./clean.sh

echo "Building .NET Core image..."
./build.sh && ./run.sh && popd

echo "Starting .NET Core test..."
ab -n 1000 -c 10 http://localhost:8081/