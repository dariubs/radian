# [![Radian](/public/img/logo.small.png?raw=true "radian")](https://github.com/dariubs/radian)

[![Build Status](https://travis-ci.org/dariubs/radian.svg?branch=master)](https://travis-ci.org/dariubs/radian)   [![Hound](https://img.shields.io/badge/houndci-golint-ff69b4.svg)](https://houndci.com)  [![Apache License](https://img.shields.io/badge/license-Apache-blue.svg)](https://github.com/dariubs/radian/blob/master/license)

Simple image server in golang.

## Docker quick start

One of the quickest ways to get Radian up and running on your machine is by using Docker.

1. Edit user info in config.toml file:

```toml
[user]
accesskey = "admin"
privatekey = "123456"
```

2. Build dockerfile and run your container :

```sh
docker build -t radian .
docker run -d -p 2112:2112 --volume /path/to/your/storage:/data --name radian-server radian
```

## routes

### Admin routes
(*authentication needed*)

Upload from postfile :
- /upload/sendfile

Upload from url:
- /upload/byurl

Simple gui for manual upload:
- /upload

### Public routes

Show file:
- /show/:filename

Show thumbnail:
- /resize/thumbnail/:width/:height/:filename

and more ...
