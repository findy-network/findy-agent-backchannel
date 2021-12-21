#!/bin/bash

if [[ $(uname) == "Linux" ]] ; then
    docker run --rm --net=host eclipse/che-ip
else
    echo "host.docker.internal"
fi
