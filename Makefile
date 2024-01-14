build:
	go build -o bin/notelify-logging-svc
	
serve: build
	./bin/notelify-logging-svc

test:
	go test -v -tags=myenv ./...
	Env=dev go test -v -tags=myenv ./...


