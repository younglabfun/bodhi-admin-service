#!/usr/bin/env bash

echo "build api..."
goctl api go -api admin.api -dir ../  --style=goZero --home ../../../../deploy/template
#goctl api format --dir ./
