#!/bin/bash

docker run -p $2:50051 --env-file ../$1/.env $1