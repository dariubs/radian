# [![Radian](/public/img/logo.small.png?raw=true "radian")](https://github.com/dariubs/radian)

[![Apache License](https://img.shields.io/badge/license-Apache-blue.svg)](https://github.com/dariubs/radian/blob/master/license)

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

admin routes
------------
(*authentication needed*)

**/upload/sendfile**: Upload from postfile 
**/upload/byurl**: Upload from url
**/upload**: Simple gui for manual upload

Public routes
-------------

**/show/:filename**: Show file
**/resize/thumbnail/:width/:height/:filename**: Show thumbnail

and more ...


License
--------

Published under [Apache License 2.0](license)