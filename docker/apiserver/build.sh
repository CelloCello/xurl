#!/bin/bash

IMAGE=xurl/apiserver:latest
cmd="docker build -t $IMAGE -f ./docker/apiserver/Dockerfile ."
echo $cmd
eval $cmd
