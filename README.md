# Bus Schedule API

[![Build Status](https://travis-ci.org/rodrigo-brito/bus-api-go.svg?branch=master)](https://travis-ci.org/rodrigo-brito/bus-api-go) [![Coverage Status](https://coveralls.io/repos/github/rodrigo-brito/bus-api-go/badge.svg)](https://coveralls.io/github/rodrigo-brito/bus-api-go)

Generic API for bus schedules in Go.<br>

### Development

Copy the settings from sample:
```bash
cp config/settings_sample.yaml config/settings.yaml
```
Start docker enviroment (`docker` and `docker-compose` required)
```bash
make run
```
Service avaliable in `http://localhost:5000`

Java Version with Spring Boot: https://github.com/rodrigo-brito/onibus-api

### Endpoints

- `/status` - Health check
- `/msg/mail` - Email service
- `/api/v1/bus` - All bus
- `/api/v1/bus/:id` - Bus by ID
- `/api/v1/bus/:id/schedule` - Schedules by bus ID
- `/api/v1/bus/:id/schedule/daytype` - Bus schedules separated by day type
- `/api/v1/company` - All companies
- `/api/v1/company/:id` - Company by ID
- `/api/v1/company/:id/bus` - Bus by company ID

### Projects that used it
Ônibus Sabará: https://horarios.sabaramais.com.br
