package repository

import (
	"context"

	"github.com/kouleen/common/pkg/mysql"
	"github.com/kouleen/idl/kitex_gen/room"
	"github.com/kouleen/room-center/modle"
	"github.com/kouleen/room-center/utils"
	"gorm.io/gorm"
)

func QueryHeaderPage(ctx context.Context, req *room.RoomHeaderRequest) (resp []*modle.RoomHeader, total int64, err error) {
	query := getRoomHeaderQuery(ctx, req)
	if req.Params != nil {
		if req.GetParams().GetBeginTime() != "" && req.GetParams().GetEndTime() != "" {
			startUTC, endUTC, err := utils.DateToLocalRange(req.GetParams().GetBeginTime(), req.GetParams().GetEndTime())
			if err != nil {
				return nil, 0, err
			}
			query = query.Where("create_time between ? and ?", startUTC, endUTC)
		}
	}
	if err = query.Count(&total).Error; err != nil || total == 0 {
		return
	}
	query = query.Order("create_time desc,id desc,status desc,open desc")
	i := (req.GetCurrent() - 1) * req.GetSize()
	if err = query.Offset(int(i)).Limit(int(req.GetSize())).Find(&resp).Error; err != nil {
		return
	}
	return
}

func getRoomHeaderQuery(ctx context.Context, req *room.RoomHeaderRequest) (query *gorm.DB) {
	query = mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.RoomHeader{}).Where("is_delete = 0")
	if req.RoomNumber != nil {
		query = query.Where("room_number = ?", req.GetRoomNumber())
	}
	if req.RoomType != nil {
		query = query.Where("room_type = ?", req.GetRoomType())
	}
	if req.Status != nil {
		query = query.Where("status = ?", req.GetStatus())
	}
	if req.Open != nil {
		query = query.Where("open = ?", req.GetOpen())
	}
	if req.GetRoomName() != "" {
		query = query.Where("room_name like ?", "%"+req.GetRoomName()+"%")
	}
	return query
}

func QueryHeaderList(ctx context.Context, req *room.RoomHeaderRequest) (resp []*modle.RoomHeader, err error) {
	query := getRoomHeaderQuery(ctx, req)
	if req.Params != nil {
		if req.GetParams().GetBeginTime() != "" && req.GetParams().GetEndTime() != "" {
			startUTC, endUTC, err := utils.DateToLocalRange(req.GetParams().GetBeginTime(), req.GetParams().GetEndTime())
			if err != nil {
				return nil, err
			}
			query = query.Where("create_time between ? and ?", startUTC, endUTC)
		}
	}
	query = query.Order("create_time desc,id desc,status desc,open desc")
	if err = query.Find(&resp).Error; err != nil {
		return
	}
	return
}

func QueryHeaderById(ctx context.Context, id int64) (resp *modle.RoomHeader, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.RoomHeader{}).First(&resp, id).Error; err != nil {
		return nil, err
	}
	return
}
