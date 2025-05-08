# 修改内容

因项目过于老旧，使用的quic库接口已经更新，导致高版本go无法编译执行代码

修改了库github.com/lucas-clemente/quic-go为github.com/quic-go/quic-go，并将对应接口修改，以适应高版本go语言

# qsocks

A socks5 proxy over quic.

[![Travis](https://travis-ci.com/net-byte/qsocks.svg?branch=main)](https://github.com/net-byte/qsocks)
[![Go Report Card](https://goreportcard.com/badge/github.com/net-byte/qsocks)](https://goreportcard.com/report/github.com/net-byte/qsocks)
![image](https://img.shields.io/badge/License-MIT-orange)
![image](https://img.shields.io/badge/License-Anti--996-red)

# Usage
```
Usage of /main:
  -S    server mode
  -bypass
        bypass private ip
  -l string
        local address (default "127.0.0.1:1080")
  -s string
        server address (default ":8443")
  -ck string
        client key file path (default "../certs/client.key")
  -cp string
        client pem file path (default "../certs/client.pem")
  -sk string
        server key file path (default "../certs/server.key")
  -sp string
        server pem file path (default "../certs/server.pem")
```

# Docker

## Run client
```
docker run -d --restart=always --name qsocks-client -p 1083:1083 -p 1083:1083/udp netbyte/qsocks -l :1083 -s SERVER_IP:8443 -ck=/app/certs/client.key -cp=/app/certs/client.pem -sk=/app/certs/server.key -sp=/app/certs/server.pem

```

## Run server
```
docker run -d --restart=always --name qsocks-server -p 8443:8443/udp netbyte/qsocks -S -s :8443 -ck=/app/certs/client.key -cp=/app/certs/client.pem -sk=/app/certs/server.key -sp=/app/certs/server.pem
```

## Setting on linux
It is recommended to increase the maximum buffer size by running:
```
sysctl -w net.core.rmem_max=2500000
```

# License
[The MIT License (MIT)](https://raw.githubusercontent.com/net-byte/qsocks/main/LICENSE)


