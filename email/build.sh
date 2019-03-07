#!/bin/sh

protoc -I pb/ pb/*.proto --go_out=plugins=grpc:pb
