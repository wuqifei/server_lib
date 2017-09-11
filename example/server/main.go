package main

import (
	"flag"
	"time"

	"github.com/wuqifei/server_lib/perf"

	"github.com/wuqifei/server_lib/libnet"
	"github.com/wuqifei/server_lib/logs"
	"github.com/wuqifei/server_lib/signal"
	"os"
	"runtime/pprof"
)

func main() {
	initLogger()
	libServer()
	ips := []string{":8080"}
	perf.Init(ips)

	f, err := os.Create("cpu.prof")
    if err != nil {
        logs.Error(err)
    }
    pprof.StartCPUProfile(f)
    go func () {
		time.Sleep(time.Duration(60*20) * time.Second)
		pprof.StopCPUProfile()
    	}()
	signal.InitSignal()
}

func initLogger() {

	logger := logs.GetLibLogger()
	logger.SetLogger("console", `{"color":false}`)
	logger.EnableFuncCallDepth(true)

	
}

func libServer() {
	//解析命令行
	flag.Parse()

	options := &libnet.ServerOptions{}
	options.Network = "tcp"
	options.Address = ":6868"
	options.IsLittleIndian = false
	options.SendQueueBuf = 10
	options.RecvQueueBuf = 10

	options.SendTimeOut = time.Duration(180) * time.Second
	options.RecvTimeOut = time.Duration(180) * time.Second //5s 超时间
	options.ReadTimeOutTimes = 3
	options.MaxRecvBufferSize = 8 * 1024
	options.MaxSendBufferSize = 8 * 1024

	server := libnet.Serve(options)
	server.RegistRoute(100, func(content []byte, sessionID uint64) (args []interface{}) {
		args = make([]interface{}, 0)
		args = append(args, uint16(1000))
		args = append(args, []byte{11, 22, 33, 44})
		return args
	})

	server.OnClose(OnClose)

	go server.Run()
}
func OnClose (sessID uint64) {
	logs.Info("sessID:%l",sessID)
}
