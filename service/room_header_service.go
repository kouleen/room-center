package service

import (
	"context"
	"errors"

	"github.com/bwmarrin/snowflake"
	"github.com/kouleen/common/pkg/alloc"
	"github.com/kouleen/common/pkg/ctxutil"
	"github.com/kouleen/idl/kitex_gen/common"
	"github.com/kouleen/idl/kitex_gen/room"
	"github.com/kouleen/idl/kitex_gen/rpc"
	"github.com/kouleen/room-center/modle"
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

func CreateHeader(ctx context.Context, req *room.RoomHeaderRequest) (resp bool, err error) {
	segmentAlloc, err := alloc.NewSegmentAlloc(ctx, rpc.ROOM_RPC_SERVER)
	if err != nil {
		return
	}
	obfuscate, err := segmentAlloc.NextObfuscate(ctx)
	if err != nil {
		return
	}
	node, err := snowflake.NewNode(1)
	if err != nil {
		return false, err
	}
	id := node.Generate().Int64()
	entity := &modle.RoomHeader{
		ID:         id,
		RoomNumber: obfuscate,
		RoomName:   req.GetRoomName(),
		RoomType:   req.RoomType,
		RoomImage:  req.GetRoomImage(),
		Background: req.GetBackground(),
		CreatedBy:  ctxutil.GetUserId(ctx),
	}
	if err = repository.CreateHeader(ctx, entity); err != nil {
		return false, err
	}
	return true, nil
}

func UpdateHeader(ctx context.Context, req *room.RoomHeaderRequest) (resp bool, err error) {
	item, err := repository.QueryHeaderById(ctx, req.GetId())
	if err != nil {
		return
	}
	item.RoomName = req.GetRoomName()
	item.RoomType = req.RoomType
	item.RoomImage = req.GetRoomImage()
	item.Background = req.GetBackground()
	item.Status = req.Status
	item.Open = req.Open
	item.UpdatedBy = ctxutil.GetUserId(ctx)
	if err = repository.UpdateHeader(ctx, item); err != nil {
		return false, err
	}
	return true, nil
}

func DeleteHeader(ctx context.Context, req *room.RoomHeaderRequest) (resp bool, err error) {
	item, err := repository.QueryHeaderById(ctx, req.GetId())
	if err != nil {
		return
	}
	if common.BaseStatus_ENABLE == common.BaseStatus(*item.Status) {
		return false, errors.New("deleting an enable room is prohibited")
	}
	if common.BaseStatus_ENABLE == common.BaseStatus(*item.Open) {
		return false, errors.New("deleting an open room is prohibited")
	}
	isDelete := int8(common.BaseStatus_ENABLE)
	item.IsDelete = &isDelete
	item.UpdatedBy = ctxutil.GetUserId(ctx)
	if err = repository.UpdateHeader(ctx, item); err != nil {
		return false, err
	}
	return true, nil
}
