protoc -I /usr/local/include -I ./ -I /Users/appointy/Desktop/Desk/protoc-gen-events/example/proto --go_out=plugins=grpc:./example  ./example/proto/example/example.proto

protoc -I /usr/local/include -I ./  --go_out=plugins=grpc:./example  ./example/proto/events/events.proto

protoc -I /usr/local/include -I ./  --go_out=plugins=grpc:./example  ./example/proto/eventspush/push.proto

packr && go build && protoc -I /usr/local/include -I  ./ -I /Users/appointy/Desktop/Desk/protoc-gen-events/example/proto --plugin=protoc-gen-events=protoc-gen-events  --events_out=:./example  ./example/proto/example/example.proto && goimports -w ./example/pb