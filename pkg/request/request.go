package request

import (
	"context"
	"log"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/metadata"
)

// ReqInfo - interface
type ReqInfo interface {
	LogRequest()
	LogDuration()
	LogError(error)
	GetRequestUUID() string
	GetServiceName() string
	GetRequestTime() time.Time
}

// NewRequestInfo - gathers info from metadat and makes request info
func NewRequestInfo(ctx context.Context, serviceName string) ReqInfo {
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

	return &info{
		requestUUID: (uuid.NewV4()).String(),
		requestTime: time.Now(),
		serviceName: serviceName,
		senderInfo:  sender,
	}
}

// info - request info
type info struct {
	requestUUID string
	requestTime time.Time
	serviceName string
	senderInfo  ReqInfo
}

// LogRequest - logging - incoming request
func (i *info) LogRequest() {
	log.Printf("<<ServiceName:'%s', RequestUUID:'%s', Time: '%s'>>\n SenderName:'%s', SenderRequestUUID:'%s', SenderTime:'%s'\n",
		i.serviceName,
		i.requestUUID,
		i.requestTime,
		i.senderInfo.GetServiceName(),
		i.senderInfo.GetRequestUUID(),
		i.senderInfo.GetRequestTime(),
	)
}

// LogError - log error
func (i *info) LogError(err error) {
	log.Printf("<<ServiceName:'%s', RequestUUID:'%s'>>\n Error: %s\n",
		i.GetServiceName(),
		i.GetRequestUUID(),
		err,
	)
}

func (i *info) LogDuration() {
	log.Printf("<<ServiceName:'%s', RequestUUID:'%s'>>Duration: \n Error: %s\n",
		i.GetServiceName(),
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
