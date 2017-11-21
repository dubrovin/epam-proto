default: model db watcher reserver

model:
	protoc -I api/ api/model.proto --go_out=plugins=grpc:api

db:
	protoc -I api/ api/db.proto --go_out=plugins=grpc:api

watcher:
	protoc -I api/ api/watcher.proto --go_out=plugins=grpc:api

reserver:
	protoc -I api/ api/reserver.proto --go_out=plugins=grpc:api

