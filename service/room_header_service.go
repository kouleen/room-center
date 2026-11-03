package service

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/kouleen/common/client"
	"github.com/kouleen/common/pkg/alloc"
	"github.com/kouleen/common/pkg/ctxutil"
	"github.com/kouleen/idl/kitex_gen/common"
	"github.com/kouleen/idl/kitex_gen/room"
	"github.com/kouleen/idl/kitex_gen/rpc"
	"github.com/kouleen/idl/kitex_gen/workflow"
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
		Remark:     req.GetRemark(),
		CreatedBy:  ctxutil.GetUserId(ctx),
	}
	if err = repository.CreateHeader(ctx, entity); err != nil {
		return false, err
	}
	return true, nil
}

func UpdateHeader(ctx context.Context, req *room.RoomHeaderRequest) (resp bool, err error) {
	now := time.Now()
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
	item.Remark = req.GetRemark()
	item.UpdatedBy = ctxutil.GetUserId(ctx)
	item.UpdateTime = &now
	if err = repository.UpdateHeader(ctx, item); err != nil {
		return false, err
	}
	return true, nil
}

func CommitHeader(ctx context.Context, req *room.RoomHeaderRequest) (resp bool, err error) {
	now := time.Now()
	item, err := repository.QueryHeaderById(ctx, req.GetId())
	if err != nil {
		return
	}
	process, err := client.GetWorkflowRpc(ctx).SaveWorkflowProcess(ctx, &workflow.WorkflowProcessRequest{
		Module:   rpc.ROOM_RPC_SERVER,
		Business: item.TableName(),
		TradeNo:  strconv.FormatInt(item.ID, 10),
		Remark:   req.GetRemark(),
	})
	if err != nil {
		return
	}
	item.FlowNo = process.FlowNo
	item.RoomName = req.GetRoomName()
	item.RoomType = req.RoomType
	item.RoomImage = req.GetRoomImage()
	item.Background = req.GetBackground()
	status := req.GetStatus() + 1
	item.Status = &status
	item.Open = req.Open
	item.Remark = req.GetRemark()
	item.UpdatedBy = ctxutil.GetUserId(ctx)
	item.UpdateTime = &now
	if err = repository.UpdateHeader(ctx, item); err != nil {
		return false, err
	}
	return true, nil
}

func DeleteHeader(ctx context.Context, req *room.RoomHeaderRequest) (resp bool, err error) {
	itemList, err := repository.QueryHeaderByIdList(ctx, req.GetIdList())
	if err != nil {
		return
	}
	closeStatus := modle.CLOSED
	for _, item := range itemList {
		// 不是待提交和驳回状态
		if modle.PREPARE != *item.Status && modle.REJECTED != *item.Status {
			return false, errors.New("deleting an enable room is prohibited")
		}
		if common.BaseStatus_ENABLE == common.BaseStatus(*item.Open) {
			return false, errors.New("deleting an open room is prohibited")
		}
		item.Status = &closeStatus
		item.UpdatedBy = ctxutil.GetUserId(ctx)
	}

	if err = repository.BatchUpdateHeader(ctx, itemList); err != nil {
		return false, err
	}
	return true, nil
}

func ReviewHeader(ctx context.Context, req *room.RoomHeaderRequest) (err error) {
	now := time.Now()
	item, err := repository.QueryHeaderByFlowNo(ctx, req.GetFlowNo())
	if err != nil {
		return
	}
	if item == nil {
		return errors.New("no such flow no")
	}
	// 状态不是审批中
	if *item.Status != 1 {
		return errors.New("incorrect status")
	}
	item.Status = req.Status
	item.UpdatedBy = ctxutil.GetUserId(ctx)
	item.UpdateTime = &now
	if err = repository.UpdateHeader(ctx, item); err != nil {
		return err
	}
	return nil
}
