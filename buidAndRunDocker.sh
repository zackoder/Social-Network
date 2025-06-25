#!/bin/bash
docker 
dockerd-rootless.sh & export DOCKER_HOST=unix://$XDG_RUNTIME_DIR/docker.sock