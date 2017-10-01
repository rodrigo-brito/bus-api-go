# Bus Schedule API

[![Build Status](https://travis-ci.org/rodrigo-brito/bus-api-go.svg?branch=master)](https://travis-ci.org/rodrigo-brito/bus-api-go) [![Coverage Status](https://coveralls.io/repos/github/rodrigo-brito/bus-api-go/badge.svg)](https://coveralls.io/github/rodrigo-brito/bus-api-go)

Generic API for bus schedules in Go.<br>
Base API used for the application: https://horarios.sabaramais.com.br/

### Development

Copy the settings and docker compose from sample: 
```bash
cp config/settings_sample.yaml config/settings.yaml
cp docker-compose_sample.yaml docker-compose.yaml
```
Start docker enviroment (`docker` and `docker-compose` required)
```
docker-compose up -d
```
Service avaliable in `http://localhost:5000`

Java Version with Spring Boot: https://github.com/rodrigo-brito/onibus-api
