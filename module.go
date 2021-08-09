package main

import (
	"encoding/json"
	"github.com/ayushbpl10/protoc-gen-events/eventspush"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"strings"
)

var IMPORTS = []string{
	"context",
	"fmt",
	"go.uber.org/fx",
	"google.golang.org/grpc/codes",
	"google.golang.org/grpc/status",
	"github.com/golang/protobuf/ptypes/empty",
}

type rightsGen struct {
	pgs.ModuleBase
	pgsgo.Context
}

func (*rightsGen) Name() string {
	return "event"
}

func (m *rightsGen) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.Context = pgsgo.InitContext(c.Parameters())
}

func (m *rightsGen) Execute(targets map[string]pgs.File, packages map[string]pgs.Package) []pgs.Artifact {

	for _, f := range targets {

		name := m.Context.OutputPath(f).SetExt(".events.go").String()
		fm := fileModel{PackageName: m.Context.PackageName(f).String()}

		for _, im := range f.Imports() {
			found := false
			for _, i := range IMPORTS {
				if i == im.Descriptor().Options.GetGoPackage() {
					found = true
				}
			}
			if !found {
				x := im.Descriptor().Options.GetGoPackage()
				if strings.Contains(x, ";") {
					x = strings.Split(x, ";")[0]
				}
				fm.Imports = append(fm.Imports, x)
			}
		}

		for _, srv := range f.Services() {

			if srv.Name().UpperCamelCase().String() == "ParentService" {
				continue
			}

			service := serviceModel{}
			service.ServiceName = srv.Name().String()
			service.PackageName = m.Context.PackageName(f).String()

			for _, rpc := range srv.Methods() {

				rpcModel := rpcModel{RpcName: rpc.Name().UpperCamelCase().String(), Input: rpc.Input().Name().UpperCamelCase().String(), Output: rpc.Output().Name().UpperCamelCase().String(), PackageName: m.Context.PackageName(f).String(), Missing: true, Protopath: rpc.FullyQualifiedName()}
				rpcModel.ConstantValue = uuid.Must(uuid.NewUUID()).String()

				if rpc.Descriptor() == nil || rpc.Descriptor().Options == nil {
					service.Rpcs = append(service.Rpcs, rpcModel)
					continue
				}

				missing := false

				opt := rpc.Descriptor().GetOptions()
				option, err := proto.GetExtension(opt, eventspush.E_Event)
				if err != nil {
					if err == proto.ErrMissingExtension {
						missing = true
					} else {
						panic(err)
					}
				}

				rpcModel.Missing = missing

				if !missing {
					byteData, err := json.Marshal(option)
					if err != nil {
						panic(err)
					}
					event := eventspush.MyEvents{}
					err = json.Unmarshal(byteData, &event)
					if err != nil {
						panic(err)
					}
					rpcModel.Push = event.Push
				}

				service.Rpcs = append(service.Rpcs, rpcModel)
			}

			if len(service.Rpcs) != 0 {
				fm.Services = append(fm.Services, service)
			}
		}

		m.OverwriteGeneratorTemplateFile(
			name,
			T.Lookup("File"),
			&fm,
		)
	}

	return m.Artifacts()
}

type rpcModel struct {
	PackageName   string
	RpcName       string
	Input         string
	Output        string
	Push          bool
	Missing       bool
	ConstantValue string
	Protopath     string
}

type serviceModel struct {
	ServiceName string
	PackageName string
	Rpcs        []rpcModel
}

type fileModel struct {
	PackageName string
	Imports     []string
	Services    []serviceModel
}
