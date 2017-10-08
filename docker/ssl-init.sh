#!/bin/sh

GetCert() {
    docker run -it \
        --rm \
        -v ${PWD}/letsencrypt/etc:/etc/letsencrypt \
        -v ${PWD}/letsencrypt/lib:/var/lib/letsencrypt \
        -v ${PWD}/letsencrypt/www:/var/www \
        certbot/certbot -t certonly --webroot -w /var/www \
        --keep-until-expiring \
        $@
}

echo "Fetching SSL..."
GetCert -d onibus.sabaramais.com.br
