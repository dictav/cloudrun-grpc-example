PROJECT := YOUR-GCP-PROJECT-NAME
SERVICE := grpc-example 
CONTAINER := gcr.io/$(PROJECT)/$(SERVICE)

PORT := 8080
ADDR := localhost:$(PORT)

GRPC_ADDR := grpc-example-REPLACE_YOUR_URL-uc.a.run.app:443
TOKEN := $(shell gcloud auth print-identity-token)

YOURNAME := REPLACE_YOUR_NAME

export PORT
export TOKEN

run: proto/service.pb.go
	go run main.go

deploy: proto/service.pb.go
	docker build -t $(CONTAINER) .
	docker push $(CONTAINER)
	gcloud run deploy $(SERVICE) \
	  --project=$(PROJECT) \
	  --image=$(CONTAINER) \
	  --region=us-central1 \
	  --platform=managed \
	  --no-allow-unauthenticated

proto/service.pb.go: proto/service.proto
	protoc --go_out=plugins=grpc:. $<

test-local:
	go run ./cmd/client/main.go -addr=$(ADDR) -insecure $(YOURNAME)

test-grpc:
	go run ./cmd/client/main.go -addr=$(GRPC_ADDR) $(YOURNAME)
