package rpccall

import (
	"github.com/xjianfeng/gocomm/logger"
	"net/rpc"
)

var (
	log           = logger.New("rpccall.log")
	rpcClient     *rpc.Client
	rpcServerAddr = ""
)

func InitRpcClient(serverAddr string) error {
	client, err := rpc.Dial("tcp", serverAddr)
	if err != nil {
		log.Error("dialing:", err)
		return err
	}
	rpcServerAddr = serverAddr
	rpcClient = client
	return nil
}

func ReConnect() {
	for i := 0; i < 3; i++ {
		err := InitRpcClient(rpcServerAddr)
		if err != nil {
			continue
		}
		break
	}
}

func RpcCall(fname string, args interface{}, reply interface{}) error {
ReTry:
	err := rpcClient.Call(fname, args, reply)
	if err != nil && err == rpc.ErrShutdown {
		ReConnect()
		goto ReTry
	}
	if err != nil {
		log.Error("RpcCall fname:%s, args:%v, reply:%v Error:%s", fname, args, reply, err.Error())
	}
	return err
}
