# gRPC

## protobuf
> 虽然可以自己定义golang struct, 但是比较繁杂. 自动生成的go package 可以完美兼容JSON; 方便互相转换

- 下载安装protoc工具 https://github.com/google/protobuf/releases

- 安装`protoc-gen-go`插件

- 解析protobuf文件:
```
protoc --go_out=plugins=grpc:. *.proto
```