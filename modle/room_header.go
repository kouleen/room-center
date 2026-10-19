package modle

import (
	"time"

	"github.com/kouleen/idl/kitex_gen/room"
)

type RoomHeader struct {
	ID         int64      `json:"id,string" gorm:"column:id;primary_key;not null"`
	RoomNumber int64      `json:"roomNumber" gorm:"column:room_number;not null;default:0"`
	RoomName   string     `json:"type" gorm:"column:room_name;not null;default:''"`
	RoomType   *int8      `json:"roomType" gorm:"column:room_type;not null;default:1"`
	RoomImage  string     `json:"roomImage" gorm:"column:room_image;default:''"`
	Background string     `json:"background" gorm:"column:background;default:''"`
	Status     *int8      `json:"status" gorm:"column:status;not null;default:0"`
	Open       *int8      `json:"open" gorm:"column:open;not null;default:0"`
	IsDelete   *int8      `json:"isDelete" gorm:"column:is_delete;not null;default:0"`
	CreatedBy  int64      `json:"createdBy,string" gorm:"column:created_by;not null;default:-1"`
	UpdatedBy  int64      `json:"updatedBy,string" gorm:"column:updated_by;not null;default:-1"`
	CreateTime *time.Time `json:"createTime" gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	UpdateTime *time.Time `json:"updateTime" gorm:"column:update_time;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (p *RoomHeader) TableName() string {
	return "room_header"
}

func (p *RoomHeader) ConvertResp() *room.RoomHeaderResponse {
	return &room.RoomHeaderResponse{
		Id:         p.ID,
		RoomNumber: p.RoomNumber,
		RoomName:   p.RoomName,
		RoomType:   p.RoomType,
		RoomImage:  p.RoomImage,
		Background: p.Background,
		Status:     p.Status,
		Open:       p.Open,
		CreatedBy:  p.CreatedBy,
		UpdatedBy:  p.UpdatedBy,
		CreateTime: p.CreateTime.UnixMilli(),
	}
}
