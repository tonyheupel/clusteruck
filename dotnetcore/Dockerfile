# Image with build tools installed -- 1.8GB!
# FROM mcr.microsoft.com/dotnet/core/sdk:3.1
# COPY . ./clusteruck
# WORKDIR /clusteruck/
# RUN dotnet build -c Release
# ENTRYPOINT ["dotnet", "run", "-c", "Release", "--no-build"]

# Multi-stage build so only runtime dependencies needed are in final image
# use :3.1-alpine as the tag instead of :3.1 for a smaller image
FROM mcr.microsoft.com/dotnet/core/sdk:3.1-alpine AS build
COPY . ./clusteruck
WORKDIR /clusteruck/
RUN dotnet build -c Release -o output

# Final image 
# (use :3.1-alpine as the tag instead of :3.1 for a smaller image
#  105MB instead of 208MB)
FROM mcr.microsoft.com/dotnet/core/aspnet:3.1-alpine AS runtime
COPY --from=build /clusteruck/output .
EXPOSE 8081
ENTRYPOINT ["dotnet", "clusteruck.dll", "--urls", "http://0.0.0.0:8081"]