#!/bin/bash

docker run -v ./config.yaml:/app/config.yaml --env-file .env token-cli $*
echo '' # i forgot to put \n in too many places