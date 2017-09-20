package libnet

import (
	"net"
	"time"

	"github.com/wuqifei/server_lib/libio"
	"github.com/wuqifei/server_lib/logs"
)

type ServerOptions struct {
	// 类型
	Network string

	// 地址
	Address string

	// cpu大小头设置
	IsLittleIndian bool

	// 接收和发送的队列大小
	SendQueueBuf int
	RecvQueueBuf int

	//接收和发送的超时时间
	SendTimeOut time.Duration
	RecvTimeOut time.Duration

	// 允许超时次数
	ReadTimeOutTimes int

	// 最大的接收字节数
	MaxRecvBufferSize int

	// 最大的发送字节数
	MaxSendBufferSize int
}

func Serve(options *ServerOptions) *Server {
	listener, err := net.Listen(options.Network, options.Address)
	if err != nil {
		panic(err)
	}
	proto := libio.New(options.IsLittleIndian, options.MaxRecvBufferSize, options.MaxSendBufferSize)
	server := NewServer(listener, proto)
	server.Options = options
	logs.Informational("libnet:Server Start")
	//打印输出的logger
	logs.Informational("libnet:network(%s),listento(%s),", options.Network, options.Address)
	return server
}
