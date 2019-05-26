#!/usr/bin/env bash

./scripts/setup.sh

# Load environment variables.
set -a
source ./.env
set +a

# Stop the container.
echo "Stopping container..."
API_CONTAINER=$(docker container ls | grep ${COMPOSE_PROJECT_NAME}_api | cut -c 1-12)
if [[ $API_CONTAINER != "" ]]; then
    docker stop $API_CONTAINER
fi
