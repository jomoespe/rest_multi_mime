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

Se ha implementado un **balanceador de capa 4** (aka. *layer 4 load balancer* o *TCP load balancer*) con [HAProxy](http://www.haproxy.org/). Para el ejemplo se ha configurado un balanceador que distribuye las peticiones entre tres instancias del servicio REST anterior.

Está montado con *docker compose*. Para ello se ha definido un [Dockerfile](./Dockerfile) para la creación de la imagen del contenedor del servicio rest. Para el balanceador haproxy se utiliza la [imagen docker oficial](https://hub.docker.com/_/haproxy/) en la versión 1.6.2. La configuración del balanceador está en [haproxy/haproxy.cfg](haproxy/haproxy.cfg).

La composición crea tres contenedores del servicio rest y un contenedor de haproxy que balancea entre estas tres instancias


### Prerequisitos

  - Docker y Docker compose


### Compilación del servicio rest

Para poder crearse una imagen minimalista del contenedor del servico rest este ha de ser compilado estáticamente. Para ello:

    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server server.go


### Arranque del balanceador

Una vez compilado el servicio rest se puede levantar la composición de contenedores. Para ello:

    docker-compose up


### Acceso

Se accederá igual que en el caso del servico REST, pero accediendo al puerto 80, que es el puerto del balanceador

    curl http://localhost:8080
    curl -H "Content-Type:application/json" http://localhost
    curl -H "Content-Type:text/xml" http://localhost


### Parada del balanceador

    docker-compose kill
