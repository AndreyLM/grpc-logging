generate-proto:
	protoc --proto_path=api/proto/v2 --proto_path=third_party --go_out=plugins=grpc:pkg/api/v2 logging-service.proto 

generate-proto-php:
	protoc --proto_path=api/proto/v2 --proto_path=third_party \
	--php_out=pkg/api/v2/php \
	--grpc_out=pkg/api/v2/php \
	--plugin=protoc-gen-grpc=/home/andrew/programs/grpc/bins/opt/grpc_php_plugin \
	logging-service.proto

DEBUG=true
DBHOST=localhost
DBPORT=5432
DBUSER=admin
DBPASSWORD=admin_1234
DBSCHEMA=logdb
GRPCPORT=9090
LOGPATH=./logs/server/log.txt
LOGPATHPROXY=./logs/proxy/log.txt
GRPCARGS= -debug=${DEBUG} -log-path=${LOGPATH} -grpc-port=${GRPCPORT} -db-host=${DBHOST} -db-port=${DBPORT} -db-schema=${DBSCHEMA} -db-user=${DBUSER} -db-password=${DBPASSWORD}
run-server:
	go run ./cmd/v2/server/main.go ${GRPCARGS}

GRPC_PROXY_PORT=9091
GRPC_ADDRESS=localhost:${GRPCPORT}
PROXYARGS= -debug=${DEBUG} -log-path=${LOGPATHPROXY} -grpc-proxy-port=${GRPC_PROXY_PORT} -grpc-server-address=${GRPC_ADDRESS}
run-proxy:
	go run ./cmd/v2/proxy/main.go ${PROXYARGS}

run-client:
	go run ./cmd/client/main.go -server=localhost:${GRPC_PROXY_PORT}

run-client-add-data:
	go run ./cmd/client-add-data/main.go -server=localhost:${GRPC_PROXY_PORT}

build:
	go build -o ./bin/v2/service ./cmd/v2/server/main.go
	go build -o ./bin/v2/proxy ./cmd/v2/proxy/main.go
