// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: action.proto

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

// ActionServiceHandler is the server API for ActionService service.
type ActionServiceHandler interface {
	// GET /api/pipeline-actions
	List(context.Context, *PipelineActionListRequest) (*PipelineActionListResponse, error)
	// POST /api/pipeline-actions/action/save
	Save(context.Context, *PipelineActionSaveRequest) (*PipelineActionSaveResponse, error)
	// DELETE /api/pipeline-actions
	Delete(context.Context, *PipelineActionDeleteRequest) (*PipelineActionDeleteResponse, error)
}

// RegisterActionServiceHandler register ActionServiceHandler to http.Router.
func RegisterActionServiceHandler(r http.Router, srv ActionServiceHandler, opts ...http.HandleOption) {
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

	add_List := func(method, path string, fn func(context.Context, *PipelineActionListRequest) (*PipelineActionListResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*PipelineActionListRequest))
		}
		var List_info transport.ServiceInfo
		if h.Interceptor != nil {
			List_info = transport.NewServiceInfo("erda.core.pipeline.action.ActionService", "List", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, List_info)
				}
				r = r.WithContext(ctx)
				var in PipelineActionListRequest
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

	add_Save := func(method, path string, fn func(context.Context, *PipelineActionSaveRequest) (*PipelineActionSaveResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*PipelineActionSaveRequest))
		}
		var Save_info transport.ServiceInfo
		if h.Interceptor != nil {
			Save_info = transport.NewServiceInfo("erda.core.pipeline.action.ActionService", "Save", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, Save_info)
				}
				r = r.WithContext(ctx)
				var in PipelineActionSaveRequest
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

	add_Delete := func(method, path string, fn func(context.Context, *PipelineActionDeleteRequest) (*PipelineActionDeleteResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*PipelineActionDeleteRequest))
		}
		var Delete_info transport.ServiceInfo
		if h.Interceptor != nil {
			Delete_info = transport.NewServiceInfo("erda.core.pipeline.action.ActionService", "Delete", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, Delete_info)
				}
				r = r.WithContext(ctx)
				var in PipelineActionDeleteRequest
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

	add_List("GET", "/api/pipeline-actions", srv.List)
	add_Save("POST", "/api/pipeline-actions/action/save", srv.Save)
	add_Delete("DELETE", "/api/pipeline-actions", srv.Delete)
}