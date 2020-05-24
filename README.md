# rpc
rpc是远程过程调用的简称，是分布式系统中不同节点间流行的通讯方式。通俗来来说就是调用远程一个函数，这个函数可能是在别的机器的某个服务，函数实现可以是不同的语言，然后通过Protobuf来实现他们之间的翻译，从而实现不同语言无障碍交流。

# protobuf命令
$ protoc --go-netrpc_out=plugins:. hello.proto

