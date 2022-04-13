#!/bin/bash

protoc --proto_path=../proto --go_out=plugins=grpc:helloworld ../proto/helloworld.proto

protoc --proto_path=../proto --go_out=plugins=grpc:countries ../proto/countries.proto