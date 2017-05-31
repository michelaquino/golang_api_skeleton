# Golang API Skeleton
> A simple API skeleton writen in Go

[![Build Status][travis-image]][travis-url]
[![License][license-image]][license-url]


## Includes
  - [Zap - Uber Log library](https://github.com/uber-go/zap)
  - [Echo Framework](https://github.com/labstack/echo)
  - [mgo - MongoDB driver](https://github.com/go-mgo/mgo/tree/v2)
  - [Go-Redis](github.com/go-redis/redis)

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

## Release History

* 0.0.1
    * Work in progress

## Meta

Michel Aquino – [@michelaquino](https://github.com/michelaquino)
Vinicius Souza – [@vsouza](https://github.com/vsouza)


[license-image]: https://img.shields.io/badge/License-GPL3.0-blue.svg
[license-url]: LICENSE
[travis-image]: https://img.shields.io/travis/michelaquinoe/golang_api_skeleton/master.svg
[travis-url]: https://travis-ci.org/michelaquino/golang_api_skeleton
