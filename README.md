# API skeleton
A simple API skeleton writen in Go

[![Build Status](https://api.travis-ci.org/michelaquino/golang_api_skeleton.svg?branch=master)](https://api.travis-ci.org/michelaquino/golang_api_skeleton.svg)

# Includes
## Libraries
- [Logrus](https://github.com/Sirupsen/logrus)
- [Echo Framework](https://github.com/labstack/echo)
- [mgo - MongoDB driver](https://github.com/go-mgo/mgo/tree/v2)

## Configuration
- Docker Compose
    - Nginx with `proxy_pass` configurated
    - API
    - MongoDB

# Dependencies
- Docker
- Docker compose

# Run
`make run`

# Usage
`curl http://localhost/healthcheck`
