proto: ## generate all protobufs in api/
	@cd api/driver;protoc -I . cflag.proto --go_out=plugins=grpc:.
	@cd api/driver;protoc -I . broker.proto --go_out=plugins=grpc:.
	@cd api/driver;protoc -I . controller.proto --go_out=plugins=grpc:.

install:
	@cd cflag; go install .

help:   ## show this help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'
