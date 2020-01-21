# Image with build tools installed -- almost 2GB!
# FROM mcr.microsoft.com/dotnet/core/sdk:3.1
# COPY . ./clusteruck
# WORKDIR /clusteruck/
# RUN dotnet build -c Release
# ENTRYPOINT ["dotnet", "run", "-c", "Release", "--no-build"]

# Multi-stage build so only runtime dependencies needed
FROM mcr.microsoft.com/dotnet/core/sdk:3.1 AS build
COPY . ./clusteruck
WORKDIR /clusteruck/
RUN dotnet build -c Release -o output

# Final image
FROM mcr.microsoft.com/dotnet/core/sdk:3.1 AS runtime
COPY --from=build /clusteruck/output .
EXPOSE 8081
ENTRYPOINT ["dotnet", "clusteruck.dll", "--urls", "http://0.0.0.0:8081"]