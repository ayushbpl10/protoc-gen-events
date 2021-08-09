# EventsPush 
Service wrapper to automatically trigger events on successful completion of client requests.

## Table of Contents
* [General Info](#general-info)
* [Motivation](#motivation)
* [Getting Started](#getting-started)
	* [Dependencies](#dependencies)
	* [Installation](#installation)
	* [Parameters](#parameters)
	* [Usage](#usage)
	* [Example](#example)
* [Available Options](#available-options)
* [Authors](#authors)

## General Info
EventsPush tool generates a wrapper over actual micro-service by implementing its methods. On successful completion of each request to service it triggers an event including the input and output data of request, which later can be used to send emails, sms, create calender events, run background jobs, etc. 

## Motivation
Earlier while implementing services if it demanded to trigger events to channels in background, extra code to push events to channel with data was needed to be written which was a redundant task.  
Also to handle such large number of events a central events type source was to be maintained, and to use data in background it needed to know what data was sent, due to this, the whole process became error prone. 
To avoid such situations and save time of developer, EventsPush was developed to handle events and push them automatically.

## Getting Started
These instructions with get you started with the tool on your local machines for development.

### Dependencies
* `go` toolchain
* `protoc` compiler in `$PATH`

### Installation
1. Clone repo to your local environment.  
	OR  
	You can get repo directly into your GOPATH using `go get github.com/ayushbpl10/protoc-gen-events/events`
2. Install crud-gen into $GOPATH using `go install .` inside the repo.

### Usage

For given protocol buffer, to generate events wrapper service run cmd   
```
protoc -I {includes} --events_out=:{output_folder_path} {input_path}/file.proto
```

### Example

For a given protocol buffer with Items service and CreateItem method: 
```Protocol Buffer
// CreateItem creates new item.
rpc CreateItem (CreateItemRequest) returns (Item) {
	option (eventspush.event) = {  
	    push: true  
	};
}
```

Events Wrapper service will look like
```go
const (   
   EVENT_Items_CreateItem pushglobal.NotificationEvent = ".package.name.Items.CreateItem"
)

// Events wrapper for Items service
type EventsItemsServer struct {  
   ItemsClient  
   eventsCli events.EventValidatorsClient  
}

func (s *EventsItemsServer) CreateItem(ctx context.Context, eventsvar *CreateItemRequest) (*Item, error) {  
  
   reqEventData, err := ptypes.MarshalAny(eventsvar)  
   if err != nil {  
      return nil, err  
   }  
  
   res, err := s.ItemsClient.CreateItem(ctx, eventsvar)  
   if err != nil {  
      return nil, err  
   }  
  
   resEventData, err := ptypes.MarshalAny(res)  
   if err != nil {  
      return nil, err  
   }  
  
   event := events.Event{  
      Type:     string(EVENT_Items_CreateItem),  
      Request:  reqEventData,  
      Response: resEventData,  
      RaisedOn: ptypes.TimestampNow(),  
      RaisedBy: userinfo.FromContext(ctx).Id,  
   }  
  
   _, pushErr := s.eventsCli.Push(ctx, &event)  
   if pushErr != nil {  
      return nil, err  
   }  
  
   return res, nil  
}
```

## Available Options
EventsPush retrives information from proto-rpcs to know for which successful rpc requests events needs to triggered. Like on creation of an appointment, a mail should be sent hence an event must be triggered, while on getting information about appointment from system should not trigger event.  

To specify if on an rpc call we must trigger event or not, we add a method option `eventspush.event` in proto corresponding to proto rpc.
```Protocol Buffer
rpc RPCName(RPCInput) returns (RPCOutput) {
	option (eventspush.event) = {  
	    push: true  
	};
}
```
Skip this option if triggering of event is not needed.

## Authors
* Ayush Gupta 