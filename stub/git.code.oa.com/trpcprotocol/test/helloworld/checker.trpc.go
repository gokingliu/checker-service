// Code generated by trpc-go/trpc-go-cmdline. DO NOT EDIT.
// source: checker.proto

package helloworld

import (
	"context"
	"fmt"

	_ "git.code.oa.com/trpc-go/trpc-go"
	_ "git.code.oa.com/trpc-go/trpc-go/http"

	"git.code.oa.com/trpc-go/trpc-go/client"
	"git.code.oa.com/trpc-go/trpc-go/codec"
	"git.code.oa.com/trpc-go/trpc-go/server"
)

/* ************************************ Service Definition ************************************ */

// GreeterService defines service
type GreeterService interface {
	SayHello(ctx context.Context, req *HelloRequest, rsp *HelloReply) (err error)
}

func GreeterService_SayHello_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {

	req := &HelloRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}, rspbody interface{}) error {
		return svr.(GreeterService).SayHello(ctx, reqbody.(*HelloRequest), rspbody.(*HelloReply))
	}

	rsp := &HelloReply{}
	err = filters.Handle(ctx, req, rsp, handleFunc)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

// GreeterServer_ServiceDesc descriptor for server.RegisterService
var GreeterServer_ServiceDesc = server.ServiceDesc{
	ServiceName: "trpc.test.helloworld.Greeter",
	HandlerType: ((*GreeterService)(nil)),
	Methods: []server.Method{
		{
			Name: "/trpc.test.helloworld.Greeter/SayHello",
			Func: GreeterService_SayHello_Handler,
		},
	},
}

// RegisterGreeterService register service
func RegisterGreeterService(s server.Service, svr GreeterService) {
	if err := s.Register(&GreeterServer_ServiceDesc, svr); err != nil {
		panic(fmt.Sprintf("Greeter register error:%v", err))
	}

}

/* ************************************ Client Definition ************************************ */

// GreeterClientProxy defines service client proxy
type GreeterClientProxy interface {
	SayHello(ctx context.Context, req *HelloRequest, opts ...client.Option) (rsp *HelloReply, err error)
}

type GreeterClientProxyImpl struct {
	client client.Client
	opts   []client.Option
}

var NewGreeterClientProxy = func(opts ...client.Option) GreeterClientProxy {
	return &GreeterClientProxyImpl{client: client.DefaultClient, opts: opts}
}

func (c *GreeterClientProxyImpl) SayHello(ctx context.Context, req *HelloRequest, opts ...client.Option) (*HelloReply, error) {

	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)

	msg.WithClientRPCName("/trpc.test.helloworld.Greeter/SayHello")
	msg.WithCalleeServiceName(GreeterServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("test")
	msg.WithCalleeServer("helloworld")
	msg.WithCalleeService("Greeter")
	msg.WithCalleeMethod("SayHello")
	msg.WithSerializationType(codec.SerializationTypePB)

	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)

	rsp := &HelloReply{}

	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}

	return rsp, nil
}
