#!/bin/bash

IMAGE=xurl/f2e:latest
cmd="docker build -t $IMAGE -f ./docker/f2e/Dockerfile ."
echo $cmd
eval $cmd
