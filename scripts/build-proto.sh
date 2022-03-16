#!/bin/bash

protoc --proto_path=. --go_out=. --micro_out=. \
		$1/proto/consignment/consignment.proto