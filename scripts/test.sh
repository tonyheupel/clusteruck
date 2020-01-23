#!/usr/bin/env bash
ab -n 1000 -c 10 http://localhost:8081/
# ab -n 10000 -c 200 http://localhost:8081/
