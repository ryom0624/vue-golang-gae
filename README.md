# Installation
```
$ git clone https://github.com/ryom0624/vue-golang-docker.git
```

# Getting-Started

```
$ docker-compose up -d
```

※ Please install docker before getting-started
https://docs.docker.com/


# Development

```
$ docker-compose exec frontend npm install
$ docker-compose exec frontend npm install axios
```

http://localhost:8081

# Deployment

frontend
```
$ cd frontend
$ docker build -t vue-compile .
$ docker container run --name npm_build vue-compile:latest npm run build
$ docker container cp npm_build:/app/dist ./
$ gcloud app deploy
```

or makefile
```
$ make build
$ cd frontend
$ gcloud app deploy
```


backend
```
$ cd ./backenrd-api
$ gcloud app deploy
```


## Deprecated
add frontend && backend-api/secret.yaml

```
$ cd ./backenrd-api
$ gcloud app deploy

$ cd ..

$ cd ./frontend
$ docker-compose exec frontend npm run build -- --mode production
$ gcloud app deploy
```

