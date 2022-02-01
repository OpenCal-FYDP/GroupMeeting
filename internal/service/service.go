package service

import (
	"context"
	"github.com/OpenCal-FYDP/GroupMeeting/rpc"
)

type GroupMeetingService struct{}

func (g *GroupMeetingService) GetGroupEvent(ctx context.Context, req *rpc.GetGroupEventReq) (*rpc.GetGroupEventRes, error) {
	panic("implement me")
}

func (g *GroupMeetingService) UpdateGroupEvent(ctx context.Context, req *rpc.UpdateGroupEventReq) (*rpc.UpdateGroupEventRes, error) {
	panic("implement me")
}

func New() *GroupMeetingService {
	return &GroupMeetingService{}
}
