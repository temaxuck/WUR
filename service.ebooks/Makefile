PROJECT_NAME ?= service-ebooks

build:
	go build -o ../build/service.ebooks/main cmd/main.go

devrun:
	go run cmd/main.go

postgres:
	docker stop $(PROJECT_NAME)-postgres || true
	docker run --rm --detach --name=$(PROJECT_NAME)-postgres \
		--env POSTGRES_USER=admin \
		--env POSTGRES_PASSWORD=admin \
		--env POSTGRES_DB=ebooks \
		--publish 5430:5432 postgres
