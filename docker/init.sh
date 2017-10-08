#!/bin/sh

set -x

mkdir -p ./volumes/nginx/conf.d
mkdir -p ./volumes/mysql
mkdir -p ./volumes/nginx/www

docker-compose -p sabaramais up -d