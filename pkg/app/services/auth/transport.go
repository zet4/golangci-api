// Code generated by genservices. DO NOT EDIT.
package auth

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

	hCheckAuth := httptransport.NewServer(
		makeCheckAuthEndpoint(svc, regCtx.Log),
		decodeCheckAuthRequest,
		encodeCheckAuthResponse,
		httptransport.ServerBefore(transportutil.StoreHTTPRequestToContext),
		httptransport.ServerAfter(transportutil.FinalizeSession),

		httptransport.ServerBefore(transportutil.MakeStoreAuthorizedRequestContext(regCtx.Log,
			regCtx.ErrTracker, regCtx.DB, regCtx.AuthSessFactory)),

		httptransport.ServerFinalizer(transportutil.FinalizeRequest),
		httptransport.ServerErrorEncoder(transportutil.EncodeError),
		httptransport.ServerErrorLogger(transportutil.AdaptErrorLogger(regCtx.Log)),
	)
	regCtx.Router.Methods("GET").Path("/v1/auth/check").Handler(hCheckAuth)

	hLogout := httptransport.NewServer(
		makeLogoutEndpoint(svc, regCtx.Log),
		decodeLogoutRequest,
		encodeLogoutResponse,
		httptransport.ServerBefore(transportutil.StoreHTTPRequestToContext),
		httptransport.ServerAfter(transportutil.FinalizeSession),

		httptransport.ServerBefore(transportutil.MakeStoreAuthorizedRequestContext(regCtx.Log,
			regCtx.ErrTracker, regCtx.DB, regCtx.AuthSessFactory)),

		httptransport.ServerFinalizer(transportutil.FinalizeRequest),
		httptransport.ServerErrorEncoder(transportutil.EncodeError),
		httptransport.ServerErrorLogger(transportutil.AdaptErrorLogger(regCtx.Log)),
	)
	regCtx.Router.Methods("GET").Path("/v1/auth/logout").Handler(hLogout)

	hUnlinkProvider := httptransport.NewServer(
		makeUnlinkProviderEndpoint(svc, regCtx.Log),
		decodeUnlinkProviderRequest,
		encodeUnlinkProviderResponse,
		httptransport.ServerBefore(transportutil.StoreHTTPRequestToContext),
		httptransport.ServerAfter(transportutil.FinalizeSession),

		httptransport.ServerBefore(transportutil.MakeStoreAuthorizedRequestContext(regCtx.Log,
			regCtx.ErrTracker, regCtx.DB, regCtx.AuthSessFactory)),

		httptransport.ServerFinalizer(transportutil.FinalizeRequest),
		httptransport.ServerErrorEncoder(transportutil.EncodeError),
		httptransport.ServerErrorLogger(transportutil.AdaptErrorLogger(regCtx.Log)),
	)
	regCtx.Router.Methods("PUT").Path("/v1/auth/unlink").Handler(hUnlinkProvider)

	hRelogin := httptransport.NewServer(
		makeReloginEndpoint(svc, regCtx.Log),
		decodeReloginRequest,
		encodeReloginResponse,
		httptransport.ServerBefore(transportutil.StoreHTTPRequestToContext),
		httptransport.ServerAfter(transportutil.FinalizeSession),

		httptransport.ServerBefore(transportutil.MakeStoreAuthorizedRequestContext(regCtx.Log,
			regCtx.ErrTracker, regCtx.DB, regCtx.AuthSessFactory)),

		httptransport.ServerFinalizer(transportutil.FinalizeRequest),
		httptransport.ServerErrorEncoder(transportutil.EncodeError),
		httptransport.ServerErrorLogger(transportutil.AdaptErrorLogger(regCtx.Log)),
	)
	regCtx.Router.Methods("GET").Path("/v1/auth/user/relogin").Handler(hRelogin)

	hLoginPublic := httptransport.NewServer(
		makeLoginPublicEndpoint(svc, regCtx.Log),
		decodeLoginPublicRequest,
		encodeLoginPublicResponse,
		httptransport.ServerBefore(transportutil.StoreHTTPRequestToContext),
		httptransport.ServerAfter(transportutil.FinalizeSession),

		httptransport.ServerBefore(transportutil.MakeStoreAnonymousRequestContext(
			regCtx.Log, regCtx.ErrTracker, regCtx.DB)),

		httptransport.ServerFinalizer(transportutil.FinalizeRequest),
		httptransport.ServerErrorEncoder(transportutil.EncodeError),
		httptransport.ServerErrorLogger(transportutil.AdaptErrorLogger(regCtx.Log)),
	)
	regCtx.Router.Methods("GET").Path("/v1/auth/{provider}").Handler(hLoginPublic)

	hLoginPrivate := httptransport.NewServer(
		makeLoginPrivateEndpoint(svc, regCtx.Log),
		decodeLoginPrivateRequest,
		encodeLoginPrivateResponse,
		httptransport.ServerBefore(transportutil.StoreHTTPRequestToContext),
		httptransport.ServerAfter(transportutil.FinalizeSession),

		httptransport.ServerBefore(transportutil.MakeStoreAuthorizedRequestContext(regCtx.Log,
			regCtx.ErrTracker, regCtx.DB, regCtx.AuthSessFactory)),

		httptransport.ServerFinalizer(transportutil.FinalizeRequest),
		httptransport.ServerErrorEncoder(transportutil.EncodeError),
		httptransport.ServerErrorLogger(transportutil.AdaptErrorLogger(regCtx.Log)),
	)
	regCtx.Router.Methods("GET").Path("/v1/auth/{provider}/private").Handler(hLoginPrivate)

	hLoginPublicOAuthCallback := httptransport.NewServer(
		makeLoginPublicOAuthCallbackEndpoint(svc, regCtx.Log),
		decodeLoginPublicOAuthCallbackRequest,
		encodeLoginPublicOAuthCallbackResponse,
		httptransport.ServerBefore(transportutil.StoreHTTPRequestToContext),
		httptransport.ServerAfter(transportutil.FinalizeSession),

		httptransport.ServerBefore(transportutil.MakeStoreAnonymousRequestContext(
			regCtx.Log, regCtx.ErrTracker, regCtx.DB)),

		httptransport.ServerFinalizer(transportutil.FinalizeRequest),
		httptransport.ServerErrorEncoder(transportutil.EncodeError),
		httptransport.ServerErrorLogger(transportutil.AdaptErrorLogger(regCtx.Log)),
	)
	regCtx.Router.Methods("GET").Path("/v1/auth/{provider}/callback/public").Handler(hLoginPublicOAuthCallback)

	hLoginPrivateOAuthCallback := httptransport.NewServer(
		makeLoginPrivateOAuthCallbackEndpoint(svc, regCtx.Log),
		decodeLoginPrivateOAuthCallbackRequest,
		encodeLoginPrivateOAuthCallbackResponse,
		httptransport.ServerBefore(transportutil.StoreHTTPRequestToContext),
		httptransport.ServerAfter(transportutil.FinalizeSession),

		httptransport.ServerBefore(transportutil.MakeStoreAuthorizedRequestContext(regCtx.Log,
			regCtx.ErrTracker, regCtx.DB, regCtx.AuthSessFactory)),

		httptransport.ServerFinalizer(transportutil.FinalizeRequest),
		httptransport.ServerErrorEncoder(transportutil.EncodeError),
		httptransport.ServerErrorLogger(transportutil.AdaptErrorLogger(regCtx.Log)),
	)
	regCtx.Router.Methods("GET").Path("/v1/auth/{provider}/callback/private").Handler(hLoginPrivateOAuthCallback)

}

func decodeCheckAuthRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request CheckAuthRequest
	if err := transportutil.DecodeRequest(&request, r); err != nil {
		return nil, errors.Wrap(err, "can't decode request")
	}

	return request, nil
}

func encodeCheckAuthResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
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

	resp := response.(CheckAuthResponse)
	wrappedResp := struct {
		transportutil.ErrorResponse
		CheckAuthResponse
	}{
		CheckAuthResponse: resp,
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

func decodeLogoutRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request LogoutRequest
	if err := transportutil.DecodeRequest(&request, r); err != nil {
		return nil, errors.Wrap(err, "can't decode request")
	}

	return request, nil
}

func encodeLogoutResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
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

	resp := response.(LogoutResponse)
	wrappedResp := struct {
		transportutil.ErrorResponse
		LogoutResponse
	}{
		LogoutResponse: resp,
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

func decodeUnlinkProviderRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request UnlinkProviderRequest
	if err := transportutil.DecodeRequest(&request, r); err != nil {
		return nil, errors.Wrap(err, "can't decode request")
	}

	return request, nil
}

func encodeUnlinkProviderResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
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

	resp := response.(UnlinkProviderResponse)
	wrappedResp := struct {
		transportutil.ErrorResponse
		UnlinkProviderResponse
	}{
		UnlinkProviderResponse: resp,
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

func decodeReloginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request ReloginRequest
	if err := transportutil.DecodeRequest(&request, r); err != nil {
		return nil, errors.Wrap(err, "can't decode request")
	}

	return request, nil
}

func encodeReloginResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
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

	resp := response.(ReloginResponse)
	wrappedResp := struct {
		transportutil.ErrorResponse
		ReloginResponse
	}{
		ReloginResponse: resp,
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

func decodeLoginPublicRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request LoginPublicRequest
	if err := transportutil.DecodeRequest(&request, r); err != nil {
		return nil, errors.Wrap(err, "can't decode request")
	}

	return request, nil
}

func encodeLoginPublicResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
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

	resp := response.(LoginPublicResponse)
	wrappedResp := struct {
		transportutil.ErrorResponse
		LoginPublicResponse
	}{
		LoginPublicResponse: resp,
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

func decodeLoginPrivateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request LoginPrivateRequest
	if err := transportutil.DecodeRequest(&request, r); err != nil {
		return nil, errors.Wrap(err, "can't decode request")
	}

	return request, nil
}

func encodeLoginPrivateResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
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

	resp := response.(LoginPrivateResponse)
	wrappedResp := struct {
		transportutil.ErrorResponse
		LoginPrivateResponse
	}{
		LoginPrivateResponse: resp,
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

func decodeLoginPublicOAuthCallbackRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request LoginPublicOAuthCallbackRequest
	if err := transportutil.DecodeRequest(&request, r); err != nil {
		return nil, errors.Wrap(err, "can't decode request")
	}

	return request, nil
}

func encodeLoginPublicOAuthCallbackResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
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

	resp := response.(LoginPublicOAuthCallbackResponse)
	wrappedResp := struct {
		transportutil.ErrorResponse
		LoginPublicOAuthCallbackResponse
	}{
		LoginPublicOAuthCallbackResponse: resp,
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

func decodeLoginPrivateOAuthCallbackRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request LoginPrivateOAuthCallbackRequest
	if err := transportutil.DecodeRequest(&request, r); err != nil {
		return nil, errors.Wrap(err, "can't decode request")
	}

	return request, nil
}

func encodeLoginPrivateOAuthCallbackResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
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

	resp := response.(LoginPrivateOAuthCallbackResponse)
	wrappedResp := struct {
		transportutil.ErrorResponse
		LoginPrivateOAuthCallbackResponse
	}{
		LoginPrivateOAuthCallbackResponse: resp,
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
