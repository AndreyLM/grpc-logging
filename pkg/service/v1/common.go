package v1

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	apiVersion = "v1"
)

func checkAPI(api string) error {
	if len(api) > 0 && apiVersion == api {
		return nil
	}

	return status.Errorf(codes.Unimplemented, "unsupported API version: service implements API version '%s'", apiVersion)
}
