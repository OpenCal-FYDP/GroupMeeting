package main

import (
	"fmt"
	"github.com/OpenCal-FYDP/Authorization"
	"github.com/OpenCal-FYDP/GroupMeeting/internal/service"
	"github.com/OpenCal-FYDP/GroupMeeting/internal/storage"
	rpc "github.com/OpenCal-FYDP/GroupMeeting/rpc"
	"github.com/twitchtv/twirp"
	"log"
	"math/rand"
	"net/http"
	"os"
)

var authorizeMethods = []string{
	"GetGroupEvent",
	"UpdateGroupEvent",
}

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func main() {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = randomString(128)
		fmt.Printf("Randomly Generated Secret: %s\n", secret)
	}
	svc := service.New(storage.New())
	server := rpc.NewGroupMeetingServiceServer(svc, twirp.WithServerInterceptors(Authorization.NewAuthorizationInterceptor([]byte(secret), authorizeMethods...)))
	jwtServer := Authorization.WithJWT(server)
	log.Fatal(http.ListenAndServe(":8080", jwtServer))
}
