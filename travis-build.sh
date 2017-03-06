#!/usr/bin/env bash

# set -xe

if [[ ${TRAVIS_TAG} ]]; then
  TRAVIS_BRANCH=${TRAVIS_TAG}
fi

DOCKER_TAG=$(sed 's/\//_/g' <<< ${TRAVIS_BRANCH})

docker build -f Dockerfile.build -t mlaccetti/mcrouter-kubernetes-provisioner:build-${DOCKER_TAG} .
docker run -v $(pwd):/go/src/github.com/mlaccetti/mcrouter-kubernetes-provisioner mlaccetti/mcrouter-kubernetes-provisioner:build-${DOCKER_TAG}

docker build -t mlaccetti/mcrouter-kubernetes-provisioner:${DOCKER_TAG} .

docker login -u="${DOCKER_USERNAME}" -p="${DOCKER_PASSWORD}"
docker push mlaccetti/mcrouter-kubernetes-provisioner:${DOCKER_TAG}
