#!/bin/bash

# $1 is the serrvice name
# $2 is the service proto file name

protoc --proto_path=. --go_out=. --micro_out=. \
		$1/proto/$2/$2.proto