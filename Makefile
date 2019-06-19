generate-proto:
	protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:pkg/api/v1 logging-service.proto \
	--php_out=plugins=grpc:pkg/api/v1/php

generate-proto-php:
	protoc --proto_path=api/proto/v1 --proto_path=third_party \
	--php_out=pkg/api/v1/php \
	--grpc_out=pkg/api/v1/php \
	--plugin=protoc-gen-grpc=/home/andrew/programs/grpc/bins/opt/grpc_php_plugin \
	logging-service.proto

DBHOST=localhost
DBPORT=5432
DBUSER=admin
DBPASSWORD=admin_1234
DBSCHEMA=logdb
GRPCPORT=9090
GRPCARGS= -grpc-port=${GRPCPORT} -db-host=${DBHOST} -db-port=${DBPORT} -db-schema=${DBSCHEMA} -db-user=${DBUSER} -db-password=${DBPASSWORD}
run-server:
	# ./cmd/server/service ${GRPCARGS}
	go run ./cmd/server/main.go ${GRPCARGS}

GRPC_PROXY_PORT=9091
GRPC_ADRESS=localhost:${GRPCPORT}
run-proxy:
	go run ./cmd/proxy-server/main.go -grpc-proxy-port=${GRPC_PROXY_PORT} -grpc-server-address=${GRPC_ADRESS}

run-client:
	go run ./cmd/client/main.go -server=localhost:${GRPC_PROXY_PORT}

run-client-add-data:
	go run ./cmd/client-add-data/main.go -server=localhost:${GRPC_PROXY_PORT}