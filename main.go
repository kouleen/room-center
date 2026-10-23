package main

import (
	"github.com/cloudwego/kitex/server"
	"github.com/kouleen/common/bootstrap"
	"github.com/kouleen/common/middleware"
	"github.com/kouleen/idl/kitex_gen/room/roomservice"
	"github.com/kouleen/idl/kitex_gen/rpc"
)

func main() {
	bootstrap.Run(rpc.ROOM_RPC_SERVER, func(option ...server.Option) server.Server {
		return roomservice.NewServer(new(RoomServiceImpl), option...)
	}, bootstrap.WithServerMiddleware(middleware.RpcServerMiddleware))
}
