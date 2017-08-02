#!/usr/bin/env bash

swagger generate server -f swagger.json -t ./server
swagger generate client -f swagger.json -t ./client