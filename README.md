# go_gin_backend

## The setup for go gin-gonic backend server

* Framework: **gin-gonic**
* Authentication: **JWT**
* Database: **PostgreSQL**
* ORM: **gorm**
* Go Version: **1.16**

## Setup process
* Create an **_.env_** file in the root directory
* Add the following key and values
  * PORT=ANY_PORT_VALUE (Eg: **4000**)
  * DB_CONFIG=POSTGRES_CONNECTION_STRING


*command to start server*
```
go run server.go
```




## Requirements to spin up PostgreSQL:
* docker >= 17.12.0+
* docker-compose

## Quick Start
* Go inside of root directory
* Run this command `docker-compose -f docker-compose.postgres.yml up -d`


## Environments
This Compose file contains the following environment variables:

* `POSTGRES_USER` the default value is **postgres**
* `POSTGRES_PASSWORD` the default value is **changeme**
* `PGADMIN_PORT` the default value is **5050**
* `PGADMIN_DEFAULT_EMAIL` the default value is **pgadmin4@pgadmin.org**
* `PGADMIN_DEFAULT_PASSWORD` the default value is **admin**

## Access to postgres: 
* `localhost:5432`
* **Username:** postgres (as a default)
* **Password:** changeme (as a default)

## Access to PgAdmin: 
* **URL:** `http://localhost:5050`
* **Username:** pgadmin4@pgadmin.org (as a default)
* **Password:** admin (as a default)

## Add a new server in PgAdmin:
* **Host name/address** `postgres`
* **Port** `5432`
* **Username** as `POSTGRES_USER`, by default: `postgres`
* **Password** as `POSTGRES_PASSWORD`, by default `changeme`
