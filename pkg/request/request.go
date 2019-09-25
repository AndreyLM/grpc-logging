package request

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// ReqInfo - interface
type ReqInfo interface {
	LogRequest()
	LogDuration()
	LogError(error)
	WrapError(codes.Code, error) error
	GetRequestUUID() string
	GetServiceName() string
	GetRequestTime() time.Time
	GetMethodName() string
}

// NewRequestInfo - gathers info from metadat and makes request info
func NewRequestInfo(ctx context.Context, serviceName, methodName string) ReqInfo {
	sender := &info{}
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if val, ok := md["request_uuid"]; ok {
			sender.requestUUID = val[0]
		}

		if val, ok := md["request_time"]; ok {
			i, err := strconv.ParseInt(val[0], 10, 64)
			if err == nil {
				sender.requestTime = time.Unix(i, 0)
			}
		}

		if val, ok := md["service_name"]; ok {
			sender.serviceName = val[0]
		}
	}
	var infoUUID string

	if sender.GetRequestUUID() != "" {
		infoUUID = sender.GetRequestUUID()
	} else {
		infoUUID = (uuid.NewV4()).String()
	}

	return &info{
		requestUUID: infoUUID,
		requestTime: time.Now(),
		serviceName: serviceName,
		senderInfo:  sender,
		methodName:  methodName,
	}
}

// info - request info
type info struct {
	requestUUID string
	requestTime time.Time
	serviceName string
	senderInfo  ReqInfo
	methodName  string
}

// LogRequest - logging - incoming request
func (i *info) LogRequest() {
	log.Printf("<<%s:%s, RequestUUID:'%s', Time: '%s'>>\n SenderName:'%s', SenderRequestUUID:'%s', SenderTime:'%s'\n",
		i.GetServiceName(),
		i.GetMethodName(),
		i.GetRequestUUID(),
		i.GetRequestTime(),
		i.senderInfo.GetServiceName(),
		i.senderInfo.GetRequestUUID(),
		i.senderInfo.GetRequestTime(),
	)
}

// LogError - log error
func (i *info) LogError(err error) {
	log.Printf("<<%s:%s, RequestUUID:'%s'>>\n Error: %s\n",
		i.GetServiceName(),
		i.GetMethodName(),
		i.GetRequestUUID(),
		err,
	)
}

func (i *info) WrapError(code codes.Code, err error) error {
	return status.Error(codes.Unknown,
		fmt.Sprintf("<<%s:%s, RequestUUID:%s>>Error: %s",
			i.GetServiceName(),
			i.GetMethodName(),
			i.GetRequestUUID(),
			err.Error(),
		),
	)
}

func (i *info) LogDuration() {
	log.Printf("<<%s:%s, RequestUUID:'%s'>>Duration: %s\n",
		i.GetServiceName(),
		i.GetMethodName(),
		i.GetRequestUUID(),
		time.Since(i.GetRequestTime()),
	)
}

func (i *info) GetRequestUUID() string {
	return i.requestUUID
}

func (i *info) GetServiceName() string {
	return i.serviceName
}

func (i *info) GetRequestTime() time.Time {
	return i.requestTime
}

func (i *info) GetMethodName() string {
	return i.methodName
}
