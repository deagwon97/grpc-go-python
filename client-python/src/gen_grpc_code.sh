#!/bin/bash

python -m grpc_tools.protoc -I ../proto/ --python_out=helloworld --grpc_python_out=helloworld ../proto/helloworld.proto

echo "" >> helloworld/__init__.py

python -m grpc_tools.protoc -I ../proto/ --python_out=countries --grpc_python_out=countries ../proto/countries.proto

echo "" >> countries/__init__.py