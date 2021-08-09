protoc -I /usr/local/include -I ./ -I /Users/appointy/Desktop/Desk/protoc-gen-events/example/proto --go_out=plugins=grpc:./example  ./example/proto/example/example.proto

protoc -I /usr/local/include -I ./  --go_out=plugins=grpc:./example  ./example/proto/events/events.proto

protoc -I /usr/local/include -I ./  --go_out=plugins=grpc:./example  ./example/proto/eventspush/push.proto

packr && go build -o protoc-gen-events && protoc -I /usr/local/include -I  ./ -I /Users/appointy/Desktop/waqt-projects/protoc-gen-events/example --plugin=protoc-gen-events=protoc-gen-events  --events_out=:./example  ./example/test.proto && goimports -w ./example/pb