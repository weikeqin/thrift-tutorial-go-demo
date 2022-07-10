# thrift-tutorial-go-demo

  github.com/apache/thrift/tutorial/go  0.16.0分支代码
  增加了通过idl生成gen-go的代码，可以直接运行

```shell
cd src 
go mod tidy
 
# 运行server  
go run main.go server.go client.go handler.go  -server true -P binary -framed false -buffered false -addr localhost:9090 -secure false

# 运行client
go run main.go server.go client.go handler.go  -server false -P binary -framed false -buffered false -addr localhost:9090 -secure false
```
