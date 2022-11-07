package web

import (
	"context"
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"time"
)

type (
	Server struct {
		http.Server
	}

	ServerOption func(server *Server)
)

func NewServer(options ...ServerOption) *Server {
	s := &Server{}
	for _, option := range options {
		option(s)
	}
	return s
}

// WithServerAddr sets the server address.
func WithServerAddr(addr string) ServerOption {
	return func(server *Server) {
		server.Addr = addr
	}
}

// WithServerReadTimeout sets the server read timeout.
func WithServerReadTimeout(timeout time.Duration) ServerOption {
	return func(server *Server) {
		server.ReadTimeout = timeout
	}
}

// WithServerWriteTimeout sets the server write timeout.
func WithServerWriteTimeout(timeout time.Duration) ServerOption {
	return func(server *Server) {
		server.WriteTimeout = timeout
	}
}

// WithServerMaxHeaderBytes sets the server max header bytes.
func WithServerMaxHeaderBytes(size int) ServerOption {
	return func(server *Server) {
		server.MaxHeaderBytes = size
	}
}

// WithServerIdleTimeout sets the server idle timeout.
func WithServerIdleTimeout(timeout time.Duration) ServerOption {
	return func(server *Server) {
		server.IdleTimeout = timeout
	}
}

// WithServerReadHeaderTimeout sets the server read header timeout.
func WithServerReadHeaderTimeout(timeout time.Duration) ServerOption {
	return func(server *Server) {
		server.ReadHeaderTimeout = timeout
	}
}

// WithServerTLSConfig sets the server TLS config.
func WithServerTLSConfig(config *tls.Config) ServerOption {
	return func(server *Server) {
		server.TLSConfig = config
	}
}

// WithServerTLSNextProto sets the server TLS next proto.
func WithServerTLSNextProto(proto map[string]func(*http.Server, *tls.Conn, http.Handler)) ServerOption {
	return func(server *Server) {
		server.TLSNextProto = proto
	}
}

// ConnState sets the server connection state.
func ConnState(f func(net.Conn, http.ConnState)) ServerOption {
	return func(server *Server) {
		server.ConnState = f
	}
}

// ErrorLog sets the server error log.
func ErrorLog(logger *log.Logger) ServerOption {
	return func(server *Server) {
		server.ErrorLog = logger
	}
}

// BaseContext sets the server base context.
func BaseContext(f func(net.Listener) context.Context) ServerOption {
	return func(server *Server) {
		server.BaseContext = f
	}
}

// ConnContext sets the server connection context.
func ConnContext(f func(context.Context, net.Conn) context.Context) ServerOption {
	return func(server *Server) {
		server.ConnContext = f
	}
}
