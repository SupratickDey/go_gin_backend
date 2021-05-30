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
  * `PORT`=ANY_PORT_VALUE (Eg: **4000**)
  * `DB_CONFIG`=POSTGRES_CONNECTION_STRING
  * `ELASTIC_APM_SERVICE_NAME`=APM_SERVICE_NAME (can be left empty)
  * `ELASTIC_APM_SERVER_URL`=APM_SERVER_URL (Eg: http://localhost:8200)
  * `ELASTIC_APM_SECRET_TOKEN`=APM_SERVER_TOKEN (use the token from apm server, Eg: xxVpmQB2HMzCL9PgBHVrnxjNXXw5J7bd79DFm6sjBJR5HPXDhcF8MSb3vv4bpg44)
  * `ELASTIC_APM_ENVIRONMENT`=APM_SERVER_ENVIRONMENT (Eg: development)


*command to start server*
```
go run server.go
```


*command to start server using docker-compose*
```
docker-compose up
```

*command to start server using docker-compose in detach mode*
```
docker-compose up -d
```



## Requirements to run PostgreSQL server:
* docker >= 17.12.0+
* docker-compose

### Quick Start
* Go inside of root directory
* Run this command `docker-compose -f docker-compose.postgres.yml up -d`


### Environments
This Compose file contains the following environment variables:

* `POSTGRES_USER` the default value is **postgres**
* `POSTGRES_PASSWORD` the default value is **changeme**
* `PGADMIN_PORT` the default value is **5050**
* `PGADMIN_DEFAULT_EMAIL` the default value is **pgadmin4@pgadmin.org**
* `PGADMIN_DEFAULT_PASSWORD` the default value is **admin**

### Access to postgres: 
* `localhost:5432`
* **Username:** postgres (as a default)
* **Password:** changeme (as a default)

### Access to PgAdmin: 
* **URL:** `http://localhost:5050`
* **Username:** pgadmin4@pgadmin.org (as a default)
* **Password:** admin (as a default)

### Add a new server in PgAdmin:
* **Host name/address** `postgres`
* **Port** `5432`
* **Username** as `POSTGRES_USER`, by default: `postgres`
* **Password** as `POSTGRES_PASSWORD`, by default `changeme`


## Run Elasticsearch-Kibana-APM server


### Quick Start
* Go inside the elasticsearch directory
* Run this command `docker-compose up -d`


### Access to Elasticsearch:
* `localhost:9200`
* **Username:** elastic (as a default)
* **Password:** PleaseChangeMe@1! (as a default)


### Access to Kibana:
* `localhost:5601`
* **Username:** elastic (as a default)
* **Password:** PleaseChangeMe@1! (as a default)

### Access to APM:
* `localhost:8200`
* **APM Server Token:** xxVpmQB2HMzCL9PgBHVrnxjNXXw5J7bd79DFm6sjBJR5HPXDhcF8MSb3vv4bpg44 (as a default)