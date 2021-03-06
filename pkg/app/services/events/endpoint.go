// Code generated by genservices. DO NOT EDIT.
package events

import (
	"context"
	"errors"
	"runtime/debug"

	"github.com/go-kit/kit/endpoint"
	"github.com/golangci/golangci-api/pkg/endpoint/apierrors"
	"github.com/golangci/golangci-api/pkg/endpoint/endpointutil"
	"github.com/golangci/golangci-api/pkg/endpoint/request"
	"github.com/golangci/golangci-shared/pkg/logutil"
)

type TrackEventRequest struct {
	Req *Request
}

type TrackEventResponse struct {
	err error
}

func makeTrackEventEndpoint(svc Service, log logutil.Log) endpoint.Endpoint {
	return func(ctx context.Context, reqObj interface{}) (resp interface{}, err error) {

		req := reqObj.(TrackEventRequest)

		reqLogger := log
		defer func() {
			if rerr := recover(); rerr != nil {
				reqLogger.Errorf("Panic occured")
				reqLogger.Infof("%s", debug.Stack())
				resp = TrackEventResponse{
					err: errors.New("panic occured"),
				}
				err = nil
			}
		}()

		if err := endpointutil.Error(ctx); err != nil {
			log.Warnf("Error occurred during request context creation: %s", err)
			resp = TrackEventResponse{
				err: err,
			}
			return resp, nil
		}

		rc := endpointutil.RequestContext(ctx).(*request.AuthorizedContext)
		reqLogger = rc.Log

		req.Req.FillLogContext(rc.Lctx)

		err = svc.TrackEvent(rc, req.Req)
		if err != nil {
			if !apierrors.IsErrorLikeResult(err) {
				rc.Log.Errorf("events.Service.TrackEvent failed: %s", err)
			}
			return TrackEventResponse{err}, nil
		}

		return TrackEventResponse{nil}, nil

	}
}
