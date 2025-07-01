package httpserver

import "github.com/kangyueyue/template/internal/httpserver/hello"

// Server 服务
type Server struct {
	*hello.HelloSvr
}

// NewServer 创建服务
func NewServer(h *hello.HelloSvr) *Server {
	return &Server{
		HelloSvr: h,
	}
}
