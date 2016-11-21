package g

import (
	"github.com/smartcaas/common/rpc"
	"time"
)

var (
	RpcClient *rpc.SingleConnRpcClient
)

func InitRpcClient() {
	RpcClient = &rpc.SingleConnRpcClient{
		RpcServers: Config().Servers,
		Timeout:    time.Duration(Config().Timeout) * time.Millisecond,
	}
}
