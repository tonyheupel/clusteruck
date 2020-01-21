# Original, most basic Dockerfile for go
#FROM golang:onbuild
#EXPOSE 8081

# Multi-stage: build in one container, static link and run in another
FROM golang:alpine as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

# Use a non-OS image so it's plain Linux
FROM scratch
COPY --from=builder /build/main /app/
EXPOSE 8080 8081
WORKDIR /app
CMD ["./main"]
