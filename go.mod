module github.com/ayushbpl10/protoc-gen-events

require (
	github.com/gobuffalo/packr v1.21.5
	github.com/golang/protobuf v1.2.0
	github.com/google/uuid v1.1.0
	github.com/lyft/protoc-gen-star v0.4.2
	github.com/spf13/pflag v1.0.3 // indirect
	go.appointy.com/google/pb v0.0.0-20181127094400-2858c30392cd
	go.appointy.com/google/userinfo v0.0.0-20181031211205-aea83d863b6e
	go.uber.org/atomic v1.3.2 // indirect
	go.uber.org/dig v1.6.0 // indirect
	go.uber.org/fx v1.8.0
	go.uber.org/multierr v1.1.0 // indirect
	golang.org/x/net v0.0.0-20181114220301-adae6a3d119a
	google.golang.org/appengine v1.2.0 // indirect
	google.golang.org/grpc v1.16.0
	pb/eventspush v0.0.0
)

replace pb/eventspush => /Users/appointy/go/src/gitlab.com/appointy/services/protoc-gen-events/example/pb/eventspush
