package message

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/kouleen/common/message"
	"github.com/kouleen/common/pkg/ctxutil"
	"github.com/kouleen/common/pkg/logutil"
	"github.com/kouleen/idl/kitex_gen/common"
	"github.com/kouleen/idl/kitex_gen/room"
	"github.com/kouleen/idl/kitex_gen/workflow"
	"github.com/kouleen/room-center/service"
)

type WorkflowReview struct{}

func (WorkflowReview) Queue() string {
	return common.ROOM_FLOW_REVIEW_QUEUE
}

func (WorkflowReview) HandleMessage(ctx context.Context, msg *message.Message) error {
	ctx = ctxutil.SetMeta(ctx, msg.TraceID, -1)
	defer func() {
		if err := recover(); err != nil {
			logutil.CtxError(ctx, "recover panic err: %#v", err)
		}
	}()
	var req workflow.WorkflowProcessRequest
	if err := json.Unmarshal(msg.Body, &req); err != nil {
		return err
	}
	if err := checkWorkflowReviewReq(&req); err != nil {
		return err
	}
	ctx = ctxutil.SetMeta(ctx, msg.TraceID, req.GetCreatedBy())
	status := req.GetStatus() - 1
	if err := service.ReviewHeader(ctx, &room.RoomHeaderRequest{
		FlowNo: req.FlowNo,
		Status: &status,
	}); err != nil {
		logutil.CtxError(ctx, "review header err: %#v", err)
	}
	return nil
}

func init() {
	message.MustRegisterConsumer(new(WorkflowReview))
}

func checkWorkflowReviewReq(req *workflow.WorkflowProcessRequest) error {
	if req.GetTradeNo() == "" {
		return errors.New("empty trade number is not allowed")
	}
	if req.GetBusiness() == "" {
		return errors.New("empty business is not allowed")
	}
	return nil
}
