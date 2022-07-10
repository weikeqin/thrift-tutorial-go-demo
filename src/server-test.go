package main

import (
	"flag"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/weikeqin/thrift-tutorial-go-demo/gen-go/tutorial"
	"testing"
)

func main() {

	addr := flag.String("addr", "localhost:9090", "Address to listen to")

	// 传输协议 TProtocol
	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryConf(nil)

	// 声明 transportFactory
	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTTransportFactory()

	// 声明传输方式 transport
	var transport thrift.TServerTransport
	var err error
	transport, err = thrift.NewTServerSocket(*addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	// CalculatorHandler 实现了 Calculator接口
	handler := NewCalculatorHandler()
	// CalculatorProcessor 组合(继承) SharedServiceProcessor  ，  SharedServiceProcessor 实现了 TProcessor接口
	// CalculatorProcessor 实现了 TProcessorFunction接口
	processor := tutorial.NewCalculatorProcessor(handler)
	// 创建一个TSimpleServer
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	fmt.Println("Starting the simple server... on ", addr)

	// server开始监听处理请求
	server.Serve()

}

func TestClientTest(t *testing.T) {

	t.Run("", func(t *testing.T) {
		fmt.Println("q23")
		t.Log("abc123")
	})

}
