protoc -I /usr/local/include -I ./  --go_out=plugins=grpc:./example  ./example/example.proto
protoc -I /usr/local/include -I ./  --go_out=plugins=grpc:./  ./event/event.proto

packr && go build && protoc -I /usr/local/include -I  ./  --plugin=protoc-gen-events=protoc-gen-events  --events_out=:./example  ./example/example.proto
