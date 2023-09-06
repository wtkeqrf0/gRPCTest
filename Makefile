include .env

MYSQL_IMAGE_NAME=mysql

dep:
	@go mod download

init:
	@docker build --tag ${MYSQL_IMAGE_NAME} .
	@docker run --name mysql -p 33060:3306 --restart on-failure -d ${MYSQL_IMAGE_NAME}

run:
	@go run main.go

test:
	go test ./mysql
	go test ./grpc

test_coverage:
	@go test ./... -coverprofile=coverage.out

help:
	@echo Below are the commands provided by the Makefile file.
	@echo Note: The commands are already arranged in the recommended execution order.
	@echo `dep` - donwload go.mod dependicies. This is required command for a newly installed project.
	@echo `init` - download and start the MySQL server by Docker. This is required command for a newly installed project.
	@echo `run` - start the gRPC server.
	@echo `test` - run on the tests.
	@echo `test_coverage` - run on the tests and generate coverage file.