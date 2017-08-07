#!/usr/bin/env bash

cd ../cmd/receipt-api
GOOS=linux GOARCH=amd64 go build -o ./../../_output/receipt-api_linux_amd64 .