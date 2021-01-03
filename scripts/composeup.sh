#!/usr/bin/env bash
echo ">>>>>>>>>>>>>>>>>BLACKSPACE PLATFORM<<<<<<<<<<<<<<<<<<<<<"
echo " Starting All Docker Containers "
docker network create web 
docker-compose up