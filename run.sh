#!/bin/bash

aws ecr get-login-password --region ap-southeast-1 | docker login --username AWS --password-stdin 697559720460.dkr.ecr.ap-southeast-1.amazonaws.com

docker compose -f ./docker/docker-compose.yml up -d