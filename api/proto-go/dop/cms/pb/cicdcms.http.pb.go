// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: cicdcms.proto

package pb

import (
	context "context"
	http1 "net/http"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	http "github.com/erda-project/erda-infra/pkg/transport/http"
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/transport/http" package it is being compiled against.
const _ = http.SupportPackageIsVersion1

// CICDCmsServiceHandler is the server API for CICDCmsService service.
type CICDCmsServiceHandler interface {
	// PUT /api/cicds/configs
	CICDCmsUpdate(context.Context, *CICDCmsUpdateRequest) (*CICDCmsUpdateResponse, error)
	// POST /api/cicds/configs
	CICDCmsCreate(context.Context, *CICDCmsCreateRequest) (*CICDCmsCreateResponse, error)
	// DELETE /api/cicds/configs
	CICDCmsDelete(context.Context, *CICDCmsDeleteRequest) (*CICDCmsDeleteResponse, error)
}

// RegisterCICDCmsServiceHandler register CICDCmsServiceHandler to http.Router.
func RegisterCICDCmsServiceHandler(r http.Router, srv CICDCmsServiceHandler, opts ...http.HandleOption) {
	h := http.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}
	encodeFunc := func(fn func(http1.ResponseWriter, *http1.Request) (interface{}, error)) http.HandlerFunc {
		handler := func(w http1.ResponseWriter, r *http1.Request) {
			out, err := fn(w, r)
			if err != nil {
				h.Error(w, r, err)
				return
			}
			if err := h.Encode(w, r, out); err != nil {
				h.Error(w, r, err)
			}
		}
		if h.HTTPInterceptor != nil {
			handler = h.HTTPInterceptor(handler)
		}
		return handler
	}

	add_CICDCmsUpdate := func(method, path string, fn func(context.Context, *CICDCmsUpdateRequest) (*CICDCmsUpdateResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CICDCmsUpdateRequest))
		}
		var CICDCmsUpdate_info transport.ServiceInfo
		if h.Interceptor != nil {
			CICDCmsUpdate_info = transport.NewServiceInfo("erda.dop.cms.CICDCmsService", "CICDCmsUpdate", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, CICDCmsUpdate_info)
				}
				r = r.WithContext(ctx)
				var in CICDCmsUpdateRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_CICDCmsCreate := func(method, path string, fn func(context.Context, *CICDCmsCreateRequest) (*CICDCmsCreateResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CICDCmsCreateRequest))
		}
		var CICDCmsCreate_info transport.ServiceInfo
		if h.Interceptor != nil {
			CICDCmsCreate_info = transport.NewServiceInfo("erda.dop.cms.CICDCmsService", "CICDCmsCreate", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, CICDCmsCreate_info)
				}
				r = r.WithContext(ctx)
				var in CICDCmsCreateRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_CICDCmsDelete := func(method, path string, fn func(context.Context, *CICDCmsDeleteRequest) (*CICDCmsDeleteResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CICDCmsDeleteRequest))
		}
		var CICDCmsDelete_info transport.ServiceInfo
		if h.Interceptor != nil {
			CICDCmsDelete_info = transport.NewServiceInfo("erda.dop.cms.CICDCmsService", "CICDCmsDelete", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, CICDCmsDelete_info)
				}
				r = r.WithContext(ctx)
				var in CICDCmsDeleteRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_CICDCmsUpdate("PUT", "/api/cicds/configs", srv.CICDCmsUpdate)
	add_CICDCmsCreate("POST", "/api/cicds/configs", srv.CICDCmsCreate)
	add_CICDCmsDelete("DELETE", "/api/cicds/configs", srv.CICDCmsDelete)
}