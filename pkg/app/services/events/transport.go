// Code generated by genservices. DO NOT EDIT.
package events

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/golangci/golangci-api/pkg/endpoint/apierrors"
	"github.com/golangci/golangci-api/pkg/transportutil"
	"github.com/pkg/errors"
)

func RegisterHandlers(svc Service, regCtx *transportutil.HandlerRegContext) {

	hTrackEvent := httptransport.NewServer(
		makeTrackEventEndpoint(svc, regCtx.Log),
		decodeTrackEventRequest,
		encodeTrackEventResponse,
		httptransport.ServerBefore(transportutil.StoreHTTPRequestToContext),
		httptransport.ServerAfter(transportutil.FinalizeSession),

		httptransport.ServerBefore(transportutil.MakeStoreAuthorizedRequestContext(regCtx.Log,
			regCtx.ErrTracker, regCtx.DB, regCtx.AuthSessFactory)),

		httptransport.ServerFinalizer(transportutil.FinalizeRequest),
		httptransport.ServerErrorEncoder(transportutil.EncodeError),
		httptransport.ServerErrorLogger(transportutil.AdaptErrorLogger(regCtx.Log)),
	)
	regCtx.Router.Methods("POST").Path("/v1/events/analytics").Handler(hTrackEvent)

}

func decodeTrackEventRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request TrackEventRequest
	if err := transportutil.DecodeRequest(&request, r); err != nil {
		return nil, errors.Wrap(err, "can't decode request")
	}

	return request, nil
}

func encodeTrackEventResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	if err := transportutil.GetContextError(ctx); err != nil {
		wrappedResp := struct {
			Error *transportutil.Error
		}{
			Error: transportutil.MakeError(err),
		}
		w.WriteHeader(wrappedResp.Error.HTTPCode)
		return json.NewEncoder(w).Encode(wrappedResp)
	}

	resp := response.(TrackEventResponse)
	wrappedResp := struct {
		transportutil.ErrorResponse
		TrackEventResponse
	}{
		TrackEventResponse: resp,
	}

	if resp.err != nil {
		if apierrors.IsErrorLikeResult(resp.err) {
			return transportutil.HandleErrorLikeResult(ctx, w, resp.err)
		}

		terr := transportutil.MakeError(resp.err)
		wrappedResp.Error = terr
		w.WriteHeader(terr.HTTPCode)
	}

	return json.NewEncoder(w).Encode(wrappedResp)
}
