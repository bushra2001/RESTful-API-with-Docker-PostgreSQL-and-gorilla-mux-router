# RESTful API with Docker, PostgreSQL and gorilla-mux-router
In this project, I have built a fully-fledged REST API that exposes GET, POST, DELETE and PUT endpoints that will subsequently allow us to perform the full range of CRUD operations. Furthermore the values of A and B are stored in postgres sql database so the values persist even if the API service is shutdown or restarted. The database as well as REST API service are deployed each in their own docker container.

## Run main.go script:
- Download golang from here.
- `go get github.com/gorilla/mux github.com/lib/pq`

## Connect postgresql:
- Download postgresql from here.
- Download pgAdmin 4 from here.
- Configure server from here.

## Deploy Docker container:
### Install Docker and Docker Compose:
`apt-get install apt-transport-https ca-certificates curl software-properties-common -y

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -

add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu focal stable"

apt-get install docker-ce docker-compose -y

docker --version
`

### Deploy a PostgreSQL with Docker

`docker search postgres

docker pull postgres:latest

docker run --name postgres-container -e POSTGRES_PASSWORD=password -d postgres`

Now, verify the Postgres container with the following command:

`docker ps`

### Deploy main.go with Docker

`docker build --tag docker-gs-ping .
docker run -d -p 10000:10000 docker-gs-ping
docker ps
`
