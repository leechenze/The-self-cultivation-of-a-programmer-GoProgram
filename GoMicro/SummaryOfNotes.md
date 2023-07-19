博学之, 审问之, 慎思之, 明辨之, 笃行之;
零、壹、贰、叁、肆、伍、陆、柒、捌、玖、拾;



零.微服务架构演进之路
    
    单体架构
    垂直架构（将大的功能模块拆分）
    微服务架构（将搜索功能最细粒度的拆分）
        独立性
        技术多样性（不同服务可以用不同语言开发，完后相互调用即可）
        开发者更加容易理解
        
        微服务如何通信？
            RPC
        微服务如何发现彼此？
            传统服务发现：IP:端口号
            客户端发现：
                微服务启动后，将自己的IP和端口进行注册，客户端查询注册，得到提供服务的IP和端口，通过负载均衡访问微服务。
            服务端发现：
                客户端访问时，不去注册中心了，通过服务发现代理直接访问。
        微服务如何部署，更新和扩容：
            微服务部署到Docker容器上，一个大型项目会拥有很多很多服务，会用到很多很多的容器，那么这么多的容器的统一管理就需要用到服务编排工具
            服务编排工具比如：K8S，swarm等，用来统一管理所有的Docker容器。




壹.RPC(NativeRPC)

    简介：
        全称为远程过程调用（Remote Procedure Call），是一个计算机通信协议
        该协议允许运行于一台计算机的程序调用另一台计算机的程序
        
    案例：
        服务端通过RPC求矩形面积和周长（rpcClient&rpcServer）
        注意要将 NativeRPC.rpcClient 和 NativeRPC.rpcServer 的 package 指定为 main 才可以单独运行
        运行 rpcServer.server.go
        再运行 rpcClient.client.go 查看结果
        
        踩坑：
            JsonRPC编码方式中的 net.Listen 方法返回了错误( listen tcp 127.0.0.1:1001: bind: permission denied )
            这是因为端口号小于1024导致的。尽量避免使用 <1024 的端口。

    调用流程：
        微服务架构下数据交互一般都是对内RPC，对外REST。
        
    会话中读写数据测试：(sessionRWData.session.go)
        运行测试文件 session_test.go
        里面的代码仔细读一下，就是一个写入的方法和一个读取的方法，看最终读取的数据和写入时的是否一致即可。
        
    在sessionRWData的示例中，只是用 my_data := "hello" 来演示如何在RPC中通信，但是实际上RPC通信的数据格式应该是
    一个网络字节流，即一个对象： key（header） 为 uint32 类型，value 为 []byte 类型
    那么这样一种格式的data就要涉及到编码和解码，所以需要用到官方的encoding/gob 进行编解码。
    数据格式和编解码定义：(sessionRWData.codec.go)
        codec.go中写好两个函数，为实现RPC的服务端和客户端做铺垫。

    实现RPC服务端：(achieveRPC)

    实现RPC客户端：(achieveRPC)

    实现RPC通信测试：(achieveRPC.rpc_test.go)
        







贰.gRPC
    
    gRPC是由Google开发，是一款语言中立，平台中立，开源的远程过程调用系统。
    gRPC是对RPC封装的一个框架，可以在多种环境中运行和交互，例如用Java写一个服务端，可以用go写的客户端进行调用。
    protobuf是一种传输协议 或 叫网络传输的数据格式，grpc就是通过protobuf这个协议进行传输的。
    
    grpc 和 protobuf 安装
    ... TODO ...

