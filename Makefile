build:
	go build -o bin/notelify-logging-service
	
serve-dev: build
	ENV=development ./bin/notelify-logging-service

serve-dev-test: build
	ENV=development_test go test -v ./...

docker-push:
	docker build -t antonyinjila/notelify-logging-service:latest --build-arg ENV=docker .
	docker push antonyinjila/notelify-logging-service:latest

docker-run:
	docker run -p 8002:8002 ENV=docker antonyinjila/notelify-logging-service:latest

docker-test:
	ENV=docker_test go test -v ./...
	

