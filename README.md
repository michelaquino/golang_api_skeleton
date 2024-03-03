# Golang API Skeleton
> A simple API skeleton writen in Go


[![Build Status](https://travis-ci.org/michelaquino/golang_api_skeleton.svg?branch=master)](https://travis-ci.org/michelaquino/golang_api_skeleton)
[![License][license-image]][license-url]


## Includes
  - [Zap - Uber Log library](https://github.com/uber-go/zap)
  - [Echo Framework](https://github.com/labstack/echo)
  - [mgo - MongoDB driver](https://github.com/go-mgo/mgo/tree/v2)
  - [Go-Redis](github.com/go-redis/redis)

  - [Prometheus - Monitoring system](https://github.com/prometheus) 

## Dependencies

- Docker
- Docker compose

## Configuration
- Docker Compose
    - Nginx with `proxy_pass` pre configurated
    - API
    - MongoDB
    - Redis

## Run
`make run`

## Usage
`curl http://localhost/healthcheck`

`curl -i -X POST -H 'Content-Type: application/json' -d '{"name": "user name", "email": "user@email.com"}' http://localhost/user`

### Prometheus
Access http://localhost:9090 to view Prometheus metrics

## Release History

* 0.3.0
    * Work in progress

## Meta

Michel Aquino – [@michelaquino](https://github.com/michelaquino)
Vinicius Souza – [@vsouza](https://github.com/vsouza)


[license-image]: https://img.shields.io/badge/License-GPL3.0-blue.svg
[license-url]: LICENSE
[travis-image]: https://img.shields.io/travis/michelaquinoe/golang_api_skeleton/master.svg
[travis-url]: https://travis-ci.org/michelaquino/golang_api_skeleton
