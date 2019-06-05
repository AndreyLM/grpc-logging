generate-proto:
	protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:pkg/api/v1 logging-service.proto \
	--php_out=plugins=grpc:pkg/api/v1/php

DBHOST=localhost
DBPORT=5432
DBUSER=admin
DBPASSWORD=admin_1234
DBSCHEMA=tutorial
GRPCPORT=9090
GRPCARGS= -grpc-port=${GRPCPORT} -db-host=${DBHOST} -db-port=${DBPORT} -db-schema=${DBSCHEMA} -db-user=${DBUSER} -db-password=${DBPASSWORD}
run-server:
	go run ./cmd/server/main.go ${GRPCARGS}

GRPC_PROXY_PORT=9091
GRPC_ADRESS=localhost:${GRPCPORT}
run-proxy:
	go run ./cmd/proxy-server/main.go -grpc-proxy-port=${GRPC_PROXY_PORT} -grpc-server-address=${GRPC_ADRESS}

run-client:
	go run ./cmd/client/main.go -server=localhost:${GRPC_PROXY_PORT}

run-client-add-data:
	go run ./cmd/client-add-data/main.go -server=localhost:${GRPC_PROXY_PORT}