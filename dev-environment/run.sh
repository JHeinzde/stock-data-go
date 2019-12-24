#!/usr/bin/env bash

#Default setting
PROJECT_DIR=/home/jonathan/Projekte/stock-data-go

cd ${PROJECT_DIR}/stock-data-go
./stock-data-go &
cd ..
cd ${PROJECT_DIR}/stock-data
npm run serve &
cd ..
docker start nginx-proxy
