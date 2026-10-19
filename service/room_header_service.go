package service

import (
	"context"

	"github.com/kouleen/idl/kitex_gen/room"
	"github.com/kouleen/room-center/repository"
)

func QueryHeaderPage(ctx context.Context, req *room.RoomHeaderRequest) (resp *room.RoomHeaderPageResponse, err error) {
	list, total, err := repository.QueryHeaderPage(ctx, req)
	if err != nil {
		return
	}
	respList := make([]*room.RoomHeaderResponse, len(list))
	for index, item := range list {
		respList[index] = item.ConvertResp()
	}
	return &room.RoomHeaderPageResponse{
		Total:   total,
		Records: respList,
	}, err
}

func QueryHeaderList(ctx context.Context, req *room.RoomHeaderRequest) (resp []*room.RoomHeaderResponse, err error) {
	list, err := repository.QueryHeaderList(ctx, req)
	if err != nil {
		return
	}
	respList := make([]*room.RoomHeaderResponse, len(list))
	for index, item := range list {
		respList[index] = item.ConvertResp()
	}
	return respList, nil
}

func QueryHeader(ctx context.Context, req *room.RoomHeaderRequest) (resp *room.RoomHeaderResponse, err error) {
	item, err := repository.QueryHeaderById(ctx, req.GetId())
	if err != nil {
		return
	}
	return item.ConvertResp(), err
}
