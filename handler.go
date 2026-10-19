package main

import (
	"context"

	"github.com/kouleen/idl/kitex_gen/room"
	"github.com/kouleen/room-center/handle"
)

// RoomServiceImpl implements the last service interface defined in the IDL.
type RoomServiceImpl struct{}

// QueryRoomHeaderPage implements the RoomServiceImpl interface.
func (s *RoomServiceImpl) QueryRoomHeaderPage(ctx context.Context, req *room.RoomHeaderRequest) (resp *room.RoomHeaderPageResponse, err error) {
	return handle.QueryHeaderPage(ctx, req)
}

// QueryRoomHeaderList implements the RoomServiceImpl interface.
func (s *RoomServiceImpl) QueryRoomHeaderList(ctx context.Context, req *room.RoomHeaderRequest) (resp []*room.RoomHeaderResponse, err error) {
	return handle.QueryHeaderList(ctx, req)
}

// QueryRoomHeader implements the RoomServiceImpl interface.
func (s *RoomServiceImpl) QueryRoomHeader(ctx context.Context, req *room.RoomHeaderRequest) (resp *room.RoomHeaderResponse, err error) {
	return handle.QueryHeader(ctx, req)
}

// CreateRoomHeader implements the RoomServiceImpl interface.
func (s *RoomServiceImpl) CreateRoomHeader(ctx context.Context, req *room.RoomHeaderRequest) (resp bool, err error) {
	return handle.CreateHeader(ctx, req)
}

// UpdateRoomHeader implements the RoomServiceImpl interface.
func (s *RoomServiceImpl) UpdateRoomHeader(ctx context.Context, req *room.RoomHeaderRequest) (resp bool, err error) {
	return handle.UpdateHeader(ctx, req)
}

// DeleteRoomHeader implements the RoomServiceImpl interface.
func (s *RoomServiceImpl) DeleteRoomHeader(ctx context.Context, req *room.RoomHeaderRequest) (resp bool, err error) {
	return handle.DeleteHeader(ctx, req)
}
