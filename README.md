REST con diferentes representaciones + TCP load balancing
=========================================================

Este proyecto contiene dos casos:

  - Servicio REST con diferentes representaciones
  - TCP load balancer

**ESTE PROYECTO ESTÁ EN DESARROLLO Y PRO LOS QUE TODO LO DESCRITO EN ESTE README NO FUNCIONA. LEER LOG DE LOS COMMITS PARA VER LO IMPLEMENTADO.**


## Servicio REST

Implementación de un servicio REST en Go con diferentes representaciones:

  - **text/html**, que será la representación por defecto
  - **application/json**
  - **text/xml**


### Arranque del servidor

    go run server.go

### Acceso

    curl http://localhost:8080
    curl -H "Content-Type:application/json" http://localhost:8080
    curl -H "Content-Type:text/xml" http://localhost:8080


## TCP load balancer

Se configurará un balanceador TCP. Se usará *HA Proxy*.


### Arranque del balanceador

    docker run --name my-running-haproxy \
           --rm \
           -ti \
           -p 80:80 \
           -v `pwd`/haproxy/haproxy.conf:/usr/local/etc/haproxy/haproxy.cfg:ro \
        haproxy:1.6.2

### Parada del balanceador

    docker stop my-running-haproxy
