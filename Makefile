proto: ## generate all protobufs in api/
	@cd driver;protoc -I . cflag.proto --go_out=plugins=grpc:.
	@cd driver;protoc -I . broker.proto --go_out=plugins=grpc:.
	@cd driver;protoc -I . controller.proto --go_out=plugins=grpc:.


help:   ## show this help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'
