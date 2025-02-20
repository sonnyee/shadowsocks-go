package direct

import (
	"net"

	"github.com/database64128/shadowsocks-go/conn"
	"github.com/database64128/shadowsocks-go/zerocopy"
)

// TCPTransparentServer is a transparent proxy server.
//
// TCPTransparentServer implements the zerocopy TCPServer interface.
type TCPTransparentServer struct{}

func NewTCPTransparentServer() (zerocopy.TCPServer, error) {
	return &TCPTransparentServer{}, nil
}

// Accept implements the zerocopy.TCPServer Accept method.
func (s *TCPTransparentServer) Accept(tc *net.TCPConn) (rw zerocopy.ReadWriter, targetAddr conn.Addr, payload []byte, err error) {
	return &DirectStreamReadWriter{rw: tc}, conn.AddrFromIPPort(tc.LocalAddr().(*net.TCPAddr).AddrPort()), nil, nil
}

// NativeInitialPayload implements the zerocopy.TCPServer NativeInitialPayload method.
func (s *TCPTransparentServer) NativeInitialPayload() bool {
	return false
}

// DefaultTCPConnCloser implements the zerocopy.TCPServer DefaultTCPConnCloser method.
func (s *TCPTransparentServer) DefaultTCPConnCloser() zerocopy.TCPConnCloser {
	return zerocopy.JustClose
}
