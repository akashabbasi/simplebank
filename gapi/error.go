package gapi

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func fieldViolation(field string, err error) *errdetails.BadRequest_FieldViolation {
	return &errdetails.BadRequest_FieldViolation{
		Field:       field,
		Description: err.Error(),
	}
}

func invalidArgErr(violations []*errdetails.BadRequest_FieldViolation) error {
	badReq := &errdetails.BadRequest{
		FieldViolations: violations,
	}
	statusInvalid := status.New(codes.InvalidArgument, "invalid parameters")
	statusDetails, err := statusInvalid.WithDetails(badReq)
	if err != nil {
		return statusInvalid.Err()
	}
	return statusDetails.Err()
}

func unauthenticatedError(err error) error {
	return status.Errorf(codes.Unauthenticated, "unauthorized: %s", err)
}
