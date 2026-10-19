package handle

import (
	"context"

	"github.com/kouleen/idl/kitex_gen/room"
	"github.com/kouleen/room-center/service"
)

func QueryHeaderPage(ctx context.Context, req *room.RoomHeaderRequest) (resp *room.RoomHeaderPageResponse, err error) {
	return service.QueryHeaderPage(ctx, req)
}

func QueryHeaderList(ctx context.Context, req *room.RoomHeaderRequest) (resp []*room.RoomHeaderResponse, err error) {
	return service.QueryHeaderList(ctx, req)
}

func QueryHeader(ctx context.Context, req *room.RoomHeaderRequest) (resp *room.RoomHeaderResponse, err error) {
	return service.QueryHeader(ctx, req)
}

func CreateHeader(ctx context.Context, req *room.RoomHeaderRequest) (resp bool, err error) {
	return
}

func UpdateHeader(ctx context.Context, req *room.RoomHeaderRequest) (resp bool, err error) {
	return
}

func DeleteHeader(ctx context.Context, req *room.RoomHeaderRequest) (resp bool, err error) {
	return
}
