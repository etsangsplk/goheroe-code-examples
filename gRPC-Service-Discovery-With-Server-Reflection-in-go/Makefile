.PHONY: help

check_defined = \
    $(strip $(foreach 1,$1, $(call __check_defined,$1,$(strip $(value 2)))))
__check_defined = \
    $(if $(value $1),, $(error Usage : make $@ $1=$(if $2,$2)))

UNAME:=$(shell uname -s)

help: ## - Show help message
	@echo "usage: make [target]\n";
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

clean:	## - Clean api directory
	@rm -rf proto/goheroe.pb.go
	@rm -rf proto/endpoints.pb
	@printf "\033[32m\xE2\x9c\x93 Clean\n\033[0m"

api:clean	## - Generate Go gRPC API
	@protoc -I ./proto/ ./proto/goheroe.proto --go_out=plugins=grpc:proto
	@printf "\033[32m\xE2\x9c\x93 Build API\n\033[0m"

api_endpoint:clean	## - Generate pb api endpoint for gRPC service
	@protoc -I ./proto/ \
	--proto_path=${GOPATH}/src \
	--proto_path=. \
	--include_imports \
	--include_source_info ./proto/goheroe.proto \
	--descriptor_set_out ./proto/endpoints.pb
	@printf "\033[32m\xE2\x9c\x93 Generate Backend API endpoint\n\033[0m"

api_endpoint_list: ## - List deployed Endpoint
	@printf "\033[32m=> Listing api deployed\n\033[0m"
	@gcloud service-management list --produced

api_endpoint_deploy:	## - Deploy Google Endpoint
	gcloud service-management deploy proto/endpoints.pb config/endpoints/api_config.yaml

docker:	## - Build docker image and push to google cloud
	@:$(call check_defined, V, docker_image_version)
	gcloud container builds submit . \
	--config=config/docker/cloudbuild.yaml \
	--substitutions _VERSION=${V}
	@printf "\033[32m\xE2\x9c\x93 Build Docker Image\n\033[0m"