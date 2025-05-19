package grpc

import (
	"github.com/sorawaslocked/ap2final_api_gateway/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func wrapError(err error) error {
	st, ok := status.FromError(err)
	if ok {
		switch st.Code() {
		case codes.NotFound:
			return model.ErrNotFound
		}
	}

	return err
}
