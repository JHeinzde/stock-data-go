#!/usr/bin/env bash

#Default setting
PROJECT_DIR=/home/jonathan/Projekte/stock-data-go

cd ${PROJECT_DIR}/stock-data-go
go build ./
cd ../
cd cd ${PROJECT_DIR}/stock-data/
npm build
