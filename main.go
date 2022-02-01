package main

import (
	"github.com/OpenCal-FYDP/GroupMeeting/internal/service"
	"github.com/OpenCal-FYDP/GroupMeeting/rpc"
	"log"
	"net/http"
)

func main() {
	service := service.New()
	server := rpc.NewGroupMeetingServiceServer(service)
	log.Fatal(http.ListenAndServe(":8080", server))
}