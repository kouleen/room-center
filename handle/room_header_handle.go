package handle

import (
	"context"
	"errors"

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
	return service.CreateHeader(ctx, req)
}

func UpdateHeader(ctx context.Context, req *room.RoomHeaderRequest) (resp bool, err error) {
	if err = checkCreateHeader(req); err != nil {
		return
	}
	return service.UpdateHeader(ctx, req)
}

func CommitHeader(ctx context.Context, req *room.RoomHeaderRequest) (resp bool, err error) {
	if err = checkUpdateHeader(req); err != nil {
		return
	}
	return service.CommitHeader(ctx, req)
}

func DeleteHeader(ctx context.Context, req *room.RoomHeaderRequest) (resp bool, err error) {
	if len(req.GetIdList()) == 0 {
		return false, errors.New("id is required")
	}
	return service.DeleteHeader(ctx, req)
}

func checkCreateHeader(req *room.RoomHeaderRequest) (err error) {
	if req.GetRoomName() == "" {
		return errors.New("room name is required")
	}
	if req.RoomType == nil {
		return errors.New("room type is required")
	}
	return nil
}

func checkUpdateHeader(req *room.RoomHeaderRequest) (err error) {
	if req.Id == nil {
		return errors.New("id is required")
	}
	return checkCreateHeader(req)
}
