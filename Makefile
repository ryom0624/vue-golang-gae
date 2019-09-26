remove:
	docker image rm vue-compile:latest
	rm -rf ./frontend/dist

build:
	docker build -t vue-compile ./frontend
	docker container run --name npm_build vue-compile:latest npm run build
	docker container cp npm_build:/app/dist ./frontend


rebuild:
	docker container rm npm_build
	docker image rm vue-compile:latest
	rm -rf ./frontend/dist
	docker build -t vue-compile ./frontend
	docker container run --name npm_build vue-compile:latest npm run build
	docker container cp npm_build:/app/dist ./frontend
