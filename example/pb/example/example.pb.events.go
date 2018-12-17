// Code generated by protoc-gen-defaults. DO NOT EDIT.

package pb

import "context"
import "gitlab.com/appointy/services/protoc-gen-events/example/events"
import "gitlab.com/appointy/services/protoc-gen-events/example/users"
import "github.com/golang/protobuf/ptypes"
import "github.com/golang/protobuf/ptypes/empty"
import "go.uber.org/fx"

import "github.com/golang/protobuf/ptypes/empty"

import "github.com/golang/protobuf/ptypes/empty"

import "gitlab.com/appointy/services/protoc-gen-events/example/pb/example"

type NotificationEvent string

const (
	EVENT_Users_AddUser NotificationEvent = "a0c105ec-01c9-11e9-9c0c-acbc32d0d9a9"

	EVENT_Users_GetUser NotificationEvent = "a0c1151e-01c9-11e9-9c0c-acbc32d0d9a9"
)

type EventsUsersServer struct {
	example.UsersServer
	eventsCli events.EventValidatorsClient
	user      users.UserIDer
}

func init() {
	options = append(options, fx.Provide(NewEventsUsersClient))
}

type EventsUsersClientResult struct {
	fx.Out
	UsersClient example.UsersClient `name:"r"`
}

func NewEventsUsersClient(c events.EventValidatorsClient, s example.UsersServer) EventsUsersClientResult {
	return EventsUsersClientResult{UsersClient: example.NewLocalUsersClient(NewEventsUsersServer(c, s))}
}
func NewEventsUsersServer(c events.EventValidatorsClient, s example.UsersServer, u users.UserIDer) example.UsersServer {
	return &EventsUsersServer{
		s,
		c,
		u,
	}
}

func (s *EventsUsersServer) AddUser(ctx context.Context, eventsvar *example.User) (*empty.Empty, error) {

	res, err := s.UsersServer.AddUser(ctx, eventsvar)
	if err != nil {
		return nil, err
	}

	reqEventData, err := ptypes.MarshalAny(&eventsvar)
	if err != nil {
		return nil, err
	}
	resEventData, err := ptypes.MarshalAny(&res)
	if err != nil {
		return nil, err
	}

	event := events.PushNotificationEventDataReq{
		EventType:    EVENT_Users_AddUser,
		RequestData:  reqEventData,
		ResponseData: resEventData,
		CreatedOn:    ptypes.TimestampNow(),
		UserId:       s.user.UserID(ctx),
	}

	res, err := s.eventsCli.Push(ctx, &event)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (s *EventsUsersServer) GetUser(ctx context.Context, eventsvar *example.GetUserReq) (*example.User, error) {

	res, err := s.UsersServer.GetUser(ctx, eventsvar)
	if err != nil {
		return nil, err
	}

	reqEventData, err := ptypes.MarshalAny(&eventsvar)
	if err != nil {
		return nil, err
	}
	resEventData, err := ptypes.MarshalAny(&res)
	if err != nil {
		return nil, err
	}

	event := events.PushNotificationEventDataReq{
		EventType:    EVENT_Users_GetUser,
		RequestData:  reqEventData,
		ResponseData: resEventData,
		CreatedOn:    ptypes.TimestampNow(),
		UserId:       s.user.UserID(ctx),
	}

	res, err := s.eventsCli.Push(ctx, &event)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (s *EventsUsersServer) UpdateUser(ctx context.Context, eventsvar *example.UpdateUserReq) (*empty.Empty, error) {

	res, err := s.UsersServer.UpdateUser(ctx, eventsvar)
	if err != nil {
		return nil, err
	}

	return res, nil

}
