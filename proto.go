package unmarshal_benchmark

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jhump/protoreflect/dynamic"
	m "github.com/mbier/unmarshal_benchmark/gen/go/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
	"io/ioutil"

	"os"
	"path/filepath"
)

func NewProto() *m.Model {
	return &m.Model{
		Text:    "string",
		Float:   3.33,
		Integer: 3,
		//Date:    timestamppb.Now(),
		Boolean: true,
	}
}

func marshalProto() ([]byte, error) {

	marshal, err := proto.Marshal(NewProto())
	if err != nil {
		return nil, err
	}

	return marshal, nil
}

func unmarshalProto(model []byte) (*m.Model, error) {

	var m m.Model

	err := proto.Unmarshal(model, &m)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func unmarshalProtoDynamicJhump(md *desc.MessageDescriptor, model []byte) (*m.Model, error) {

	msg := dynamic.NewMessage(md)
	err := proto.Unmarshal(model, msg)
	if err != nil {
		return nil, err
	}

	return &m.Model{
		Text:    msg.GetFieldByName("text").(string),
		Float:   msg.GetFieldByName("float").(float32),
		Integer: msg.GetFieldByName("integer").(int64),
		//Date:    msg.GetFieldByName("date").(*timestamppb.Timestamp),
		Boolean: msg.GetFieldByName("boolean").(bool),
	}, nil
}

func loadDescriptor(message string, importPaths []string) (*desc.MessageDescriptor, error) {
	if message == "" {
		return nil, errors.New("message field must not be empty")
	}

	var parser protoparse.Parser
	if len(importPaths) == 0 {
		importPaths = []string{"."}
	} else {
		parser.ImportPaths = importPaths
	}

	var files []string
	for _, importPath := range importPaths {
		if err := filepath.Walk(importPath, func(path string, info os.FileInfo, ferr error) error {
			if ferr != nil || info.IsDir() {
				return ferr
			}
			if filepath.Ext(info.Name()) == ".proto" {
				rPath, ferr := filepath.Rel(importPath, path)
				if ferr != nil {
					return fmt.Errorf("failed to get relative path: %v", ferr)
				}
				files = append(files, rPath)
			}
			return nil
		}); err != nil {
			return nil, err
		}
	}

	fds, err := parser.ParseFiles(files...)
	if err != nil {
		return nil, fmt.Errorf("failed to parse .proto file: %v", err)
	}
	if len(fds) == 0 {
		return nil, fmt.Errorf("no .proto files were found in the paths '%v'", importPaths)
	}

	var msg *desc.MessageDescriptor
	for _, d := range fds {
		if msg = d.FindMessage(message); msg != nil {
			break
		}
	}
	if msg == nil {
		err = fmt.Errorf("unable to find message '%v' definition within '%v'", message, importPaths)
	}
	return msg, err
}

func unmarshalProtoDynamicOfficial(md protoreflect.MessageDescriptor, model []byte) (*m.Model, error) {

	message := dynamicpb.NewMessage(md)

	err := proto.Unmarshal(model, message)
	if err != nil {
		return nil, err
	}

	return &m.Model{
		Text:    message.Get(md.Fields().ByName("text")).String(),
		Float:   float32(message.Get(md.Fields().ByName("float")).Float()),
		Integer: message.Get(md.Fields().ByName("integer")).Int(),
		//Date:    message.Get(md.Fields().ByName("date")).Interface().(timestamppb.Timestamp),
		Boolean: message.Get(md.Fields().ByName("boolean")).Bool(),
	}, nil
}

func loadDescriptorOfficial() (protoreflect.MessageDescriptor, error) {
	file, err := ioutil.ReadFile("gen/descriptor_set/proto/proto.bin")
	if err != nil {
		return nil, err
	}

	fds := new(descriptorpb.FileDescriptorSet)
	if err := proto.Unmarshal(file, fds); err != nil {
		return nil, err
	}

	files, err := protodesc.NewFiles(fds)
	if err != nil {
		return nil, err
	}

	desc, err := files.FindFileByPath("proto/model.proto")
	if err != nil {
		return nil, err
	}

	md := desc.Messages().ByName("Model")

	return md, nil
}
