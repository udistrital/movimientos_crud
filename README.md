# movimientos_crud

API CRUD para movimientos

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [Docker](https://docs.docker.com/engine/install/ubuntu/)
* [Docker Compose](https://docs.docker.com/compose/)

### Variables de Entorno
```shell
# Ejemplo que se debe actualizar acorde al proyecto
MOVIMIENTOS_CRUD_DB_USER = [descripción]
MOVIMIENTOS_CRUD_DB_PASS = [descripción]
MOVIMIENTOS_CRUD_DB_HOST = [descripción]
```
**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y están identificadas con MOVIMIENTOS_CRUD_...

### Ejecución del Proyecto
```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/movimientos_crud

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/movimientos_crud

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
MOVIMIENTOS_CRUD_HTTP_PORT=8080 MOVIMIENTOS_CRUD_DB_HOST=127.0.0.1:27017 MOVIMIENTOS_CRUD_SOME_VARIABLE=some_value bee run
```

### Ejecución Dockerfile
```shell
# docker build --tag=movimientos_crud . --no-cache
# docker run -p 80:80 movimientos_crud
```

### Ejecución docker-compose
```shell
#1. Clonar el repositorio
git clone -b develop https://github.com/udistrital/movimientos_crud

#2. Moverse a la carpeta del repositorio
cd movimientos_crud

#3. Crear un fichero con el nombre **custom.env**
touch custom.env

#4. Crear la network **back_end** para los contenedores
docker network create back_end

#5. Ejecutar el compose del contenedor
docker-compose up --build

#6. Comprobar que los contenedores estén en ejecución
docker ps
```
### Ejecución Pruebas

Pruebas unitarias
```shell
# Not Data
```

## Modelo de datos
[Modelo de Datos Relacional](https://user-images.githubusercontent.com/15944053/59788345-1dd62f00-9291-11e9-9261-1eb06d1d1454.png)  


## Estado CI
| Develop | Relese 0.0.1 | Master |
| -- | -- | -- |
|1 |2 |3 |

## Licencia
This file is part of [movimientos_crud](LICENSE).

movimientos_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

movimientos_crud is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with necesidades_crud. If not, see https://www.gnu.org/licenses/.
