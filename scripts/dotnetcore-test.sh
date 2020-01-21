#!/usr/bin/env bash
echo "Removing previous container..."
./clean.sh

echo "Building .NET Core image..."
pushd ../dotnetcore/ && ./build.sh && ./run.sh && popd

echo "Starting .NET Core test..."
./test.sh