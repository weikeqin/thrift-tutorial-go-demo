module github.com/weikeqin/thrift-tutorial-go-demo

go 1.18

require (
	github.com/apache/thrift v0.16.0
	github.com/apache/thrift/tutorial/go v0.16.0 // 替换成本地的文件了
)

// 把github.com/apache/thrift/tutorial/go/gen-go 替换成本地的 ./gen-go
replace github.com/apache/thrift/tutorial/go v0.16.0 => ./
