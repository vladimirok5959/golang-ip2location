VERSION="1.0.0"
DOCKER_IMG_NAME := golang-ip2location
CURRENT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
DATA_DIR := ${CURRENT_DIR}/data

.PHONY: default
default: test

include .makefile/minimal.makefile
include .makefile/build.makefile
include .makefile/docker.makefile

test:
	go test `go list ./... \
		| grep -v cmd/ip2location \
		| grep -v internal/consts \
		| grep -v internal/server/web`

debug:
	@-rm ./bin/ip2location
	make all
	./bin/ip2location \
		-data_dir ${DATA_DIR} \
		-deployment development \
		-host 0.0.0.0 \
		-port 8080 \
		-web_url http://localhost:8080/ \
		color=always

docker-test:
	docker run --rm \
		--network host \
		--name ${DOCKER_IMG_NAME}-test \
		-e ENV_DATA_DIR="/app/data" \
		-e ENV_DEPLOYMENT="deployment" \
		-e ENV_HOST="127.0.0.1" \
		-e ENV_PORT="8080" \
		-e ENV_WEB_URL="http://localhost:8080/" \
		-v /etc/timezone:/etc/timezone:ro \
		-v ${CURRENT_DIR}/data:/app/data \
		-it ${DOCKER_IMG_NAME}:latest
