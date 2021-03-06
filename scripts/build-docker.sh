#!/bin/sh

# Exit on error
set -e

. scripts/utils.sh

VERSION=$1
IMAGE_NAME="doraboateng/api"
USAGE="Usage: ./run build-docker [VERSION]"

if [ "$VERSION" = "" ]; then
    VERSION=$(git describe --abbrev=0 --tags)
    # VERSION=$(date "+%y.%m.0")
fi

if [ "$VERSION" = "help" ] || [ "$VERSION" = "-h" ] || [ "$VERSION" = "--help" ];
then
    echo "$USAGE"
    exit 0
fi

TAGGED_IMAGE="$IMAGE_NAME:$VERSION"
LATEST_IMAGE="$IMAGE_NAME:latest"

echo ""
echo "Building \"$TAGGED_IMAGE\". Continue? (yes/[no])"
read -r CONFIRMATION

if [ "$CONFIRMATION" != "yes" ]; then
    exit 0
fi

if image_exists "$TAGGED_IMAGE"; then
    docker image rm --force "$TAGGED_IMAGE"
fi

docker build \
    --build-arg BUILD_VERSION="$VERSION" \
    --build-arg GIT_HASH="$(git rev-parse --short HEAD)" \
    --force-rm \
    --tag "$TAGGED_IMAGE" \
    --target prod \
    .

echo ""
echo "Update \"latest\" tag for \"$IMAGE_NAME\" (\"$LATEST_IMAGE\")? (yes/[no])"
read -r CONFIRM_TAG_LATEST

if [ "$CONFIRM_TAG_LATEST" = "yes" ]; then
    if image_exists "$LATEST_IMAGE"; then
        docker image rm --force "$LATEST_IMAGE"
    fi

    docker tag "$TAGGED_IMAGE" "$LATEST_IMAGE"
fi

echo ""
echo "Publish to Docker registry? (yes/[no])"
read -r CONFIRMATION

if [ "$CONFIRMATION" = "yes" ]; then
    get_env DOCKER_HUB_TOKEN | docker login \
        --username "$(get_env DOCKER_HUB_USERNAME)" \
        --password-stdin

    if [ "$CONFIRM_TAG_LATEST" = "yes" ]; then
        docker push "$LATEST_IMAGE"
    fi

    docker push "$TAGGED_IMAGE"
fi

echo ""
echo "Publish to Google Container registry? (yes/[no])"
read -r CONFIRMATION

if [ "$CONFIRMATION" = "yes" ]; then
    GCR_HOST="gcr.io"

    gcloud auth print-access-token | docker login \
        --username oauth2accesstoken \
        --password-stdin \
        https://"$GCR_HOST"
    
    docker tag "$TAGGED_IMAGE" "$GCR_HOST/$TAGGED_IMAGE"
    docker push "$TAGGED_IMAGE"
fi

echo ""
echo "Pruning Docker resources..."
echo ""

docker system prune
