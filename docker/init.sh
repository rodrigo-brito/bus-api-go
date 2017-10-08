#!/bin/sh

if [ "$PROJECT_ENV" = "development" ]; then
    echo "-- ENV DEVELOPMENT --"
    BASE_URL="sabaramais.io"
else
    echo "-- CAUTION: ENV PRODUCTION --"
    BASE_URL="sabaramais.com.br"
fi

#!/bin/sh

GetCert() {
        docker run -it \
                --rm \
                -v ${PWD}/letsencrypt/etc:/etc/letsencrypt \
                -v ${PWD}/letsencrypt/lib:/var/lib/letsencrypt \
                -v ${PWD}/letsencrypt/www:/var/www/.well-known \
                certbot/certbot -t certonly --webroot -w /var/www \
                --keep-until-expiring \
                $@
}

set -x

mkdir -p ./volumes/nginx/conf.d
mkdir -p ./volumes/mysql
mkdir -p ./volumes/nginx/www

echo "Getting certificates..."
# GetCert -d www.$BASE_URL -d $BASE_URL -d onibus.$BASE_URL

docker-compose -p sabaramais up -d
# docker-compose logs

