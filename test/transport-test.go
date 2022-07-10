package mian

import (
	"crypto/tls"
	"flag"
	"github.com/apache/thrift/lib/go/thrift"
)

func main() error {
	framed := flag.Bool("framed", false, "Use framed transport")
	buffered := flag.Bool("buffered", false, "Use buffered transport")
	addr := flag.String("addr", "localhost:9090", "Address to listen to")
	secure := flag.Bool("secure", false, "Use tls secure transport")

	// 声明 transportFactory
	var transportFactory thrift.TTransportFactory

	cfg := &thrift.TConfiguration{
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	//  transportFactory包装
	if *buffered {
		transportFactory = thrift.NewTBufferedTransportFactory(8192)
	} else {
		transportFactory = thrift.NewTTransportFactory()
	}
	if *framed {
		transportFactory = thrift.NewTFramedTransportFactoryConf(transportFactory, cfg)
	}

	// 声明transport
	var transport thrift.TTransport
	if *secure {
		transport = thrift.NewTSSLSocketConf(*addr, cfg)
	} else {
		transport = thrift.NewTSocketConf(*addr, cfg)
	}

	// 包装 功能增强
	transport, err := transportFactory.GetTransport(transport)
	if err != nil {
		return err
	}

	defer transport.Close()

	// 打开通信传输
	// 连接套接字，必要时创建新的套接字对象。
	if err := transport.Open(); err != nil {
		return err
	}

	return nil
}
