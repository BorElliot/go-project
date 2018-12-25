## Golang RPC 相关知识

### 如何根据protoc生成代码 

指定文件目录的方式
`protoc B.proto --go_out=. --proto_path=$GOPATH/src/B `

已经在文件夹中
`protoc B.proto --go_out=.`

生成RPC服务
`protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld`

### rpc protobuf 入门

http://lameleg.com/tech/rpc-protobuf.html

### 问题

安装protoc-gen-go
https://github.com/golang/protobuf/issues/365#issuecomment-305240878

https://github.com/golang/protobuf#installation

终端代理解决go依赖被和谐的问题

https://github.com/Quinton/blog/issues/2
https://github.com/mrdulin/blog/issues/18
https://droidyue.com/blog/2016/04/04/set-shadowsocks-proxy-for-terminal/index.html


