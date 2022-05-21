sensors:
	docker build \
	-f ./k8s/dockerfile \
	-t sensors:1.0 \
	.

postgres: 
	docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres  -d postgres
admin:
	go build -o app/admin ./...

