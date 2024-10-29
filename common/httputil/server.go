package httputil

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync/atomic"

	"github.com/pkg/errors"
)

type HTTPServer struct {
	listener net.Listener
	srv      *http.Server
	closed   atomic.Bool
}

type HTTPOption func(svr *HTTPServer) error

func StarHttpServer(addr string, handler http.Handler, opts ...HTTPOption) (*HTTPServer, error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("listen errorr=", err)
		return nil, errors.New("Init listener fail")
	}
	srvCtx, srvCancel := context.WithCancel(context.Background())
	srv := &http.Server{
		Handler:           handler,
		ReadTimeout:       timeouts.ReadTimeout,
		ReadHeaderTimeout: timeouts.ReadHeaderTimeout,
		WriteTimeout:      timeouts.WriteTimeout,
		IdleTimeout:       timeouts.IdleTimeout,
		BaseContext: func(listener net.Listener) context.Context {
			return srvCtx
		},
	}
	out := &HTTPServer{listener: listener, srv: srv}

	for _, opt := range opts {
		if err := opt(out); err != nil {
			srvCancel()
			fmt.Println("apply err:", err)
			return nil, errors.New("One of http op fail")
		}
	}
	go func() {
		err := out.srv.Serve(listener)
		srvCancel()
		if errors.Is(err, http.ErrServerClosed) {
			out.closed.Store(true)
		} else {
			fmt.Println("unknow err:", err)
			panic("unknow error")
		}
	}()
	return out, nil
}

func (hs *HTTPServer) Closed() bool {
	return hs.closed.Load()
}

func (hs *HTTPServer) Stop(ctx context.Context) error {
	if err := hs.Shutdown(ctx); err != nil {
		if errors.Is(err, ctx.Err()) {
			return hs.Close()
		}
		return err
	}
	return nil
}

func (hs *HTTPServer) Shutdown(ctx context.Context) error {
	return hs.srv.Shutdown(ctx)
}

func (hs *HTTPServer) Close() error {
	return hs.srv.Close()
}

func (hs *HTTPServer) Addr() net.Addr {
	return hs.listener.Addr()
}

func WithMaxHeaderBytes(max int) HTTPOption {
	return func(srv *HTTPServer) error {
		srv.srv.MaxHeaderBytes = max
		return nil
	}
}
