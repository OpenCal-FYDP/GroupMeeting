package service

import (
	"context"
	"github.com/OpenCal-FYDP/GroupMeeting/internal/storage"
	"github.com/OpenCal-FYDP/GroupMeeting/rpc"
)

type GroupMeetingService struct{
	storage *storage.Storage
}

func (g *GroupMeetingService) GetGroupEvent(ctx context.Context, req *rpc.GetGroupEventReq) (*rpc.GetGroupEventRes, error) {
	res := new(rpc.GetGroupEventRes)
	err := g.storage.GetGroupEvent(req, res)
	if err != nil {
		return nil, err
	}
	for key, attendeeAvaVal := range res.GetAvailabilities() {
		println(key)
		println(attendeeAvaVal.DateRanges)
		println(attendeeAvaVal.GetAvailabilityID())
	}
	return res, nil
}

func (g *GroupMeetingService) UpdateGroupEvent(ctx context.Context, req *rpc.UpdateGroupEventReq) (*rpc.UpdateGroupEventRes, error) {
	res := new(rpc.UpdateGroupEventRes)
	err := g.storage.UpdateGroupEvent(req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func New(s *storage.Storage) *GroupMeetingService {
	return &GroupMeetingService{
		s,
	}
}
