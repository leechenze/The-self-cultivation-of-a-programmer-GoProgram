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
        







贰.gRPC(GRPC)


    gRPC是由Google开发，是一款语言中立，平台中立，开源的远程过程调用系统。
    gRPC是对RPC封装的一个框架，可以在多种环境中运行和交互，例如用Java写一个服务端，可以用go写的客户端进行调用。
    protobuf是一种传输协议 或 叫网络传输的数据格式，grpc就是通过protobuf这个协议进行传输的。
    
    grpc 和 protobuf 安装
        请参考：
            https://blog.csdn.net/oJinTian1234567/article/details/129305376?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522168985870416800182795636%2522%252C%2522scm%2522%253A%252220140713.130102334..%2522%257D&request_id=168985870416800182795636&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~baidu_landing_v2~default-1-129305376-null-null.142^v90^control_2,239^v3^control&utm_term=mac%20grpc%20%E5%AE%89%E8%A3%85&spm=1018.2226.3001.4187
    protobuf语法：
        ProtoBuf官网：
            https://protobuf.dev/
        文件以.proto作为文件后缀，除结构定义外的语句以分号结尾
        结构定义可以包含：message，service，enum
        RPC方法定义结尾的分号可有可无
        类型命名采用驼峰命名方式，字段命名采用小写字母加下划线分隔
            限定修饰符 数据类型 字段名称 = 字段编码值
            required string  song_name = 1
        限定修饰符：
            required: 表示一个必须字段
            optional: 表示一个可选字段
            repeated: 表示该字段可以包含 0-n 个元素，其特性和optional一样
        数据类型：
            和Go中的数据类型差不多，更多请查文档。
        字段名称：
            明明规则和Java，C的命名规则几乎相同
        字段编码值：
            1 - 15 这个区间的编码时间和空间效率是最高的，编码值越大，其编码时间和编码效率就越低。
            1900 - 2000 这个区间编码值不要用，因为是Google Protobuf 系统内部保留值
        .proto文件编译
            通过定义好的.proto文件生成Java,Python,C++,Go等各平台的代码。
            Go生成的文件后缀：xxx.pb.go 文件，每个消息类型对应一个结构体。
        更多详细语法可以参看：
            https://blog.csdn.net/weixin_42905141/article/details/125272803?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522168986112416800222837211%2522%252C%2522scm%2522%253A%252220140713.130102334..%2522%257D&request_id=168986112416800222837211&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~top_positive~default-2-125272803-null-null.142^v90^control_2,239^v3^control&utm_term=protobuf&spm=1018.2226.3001.4187
        
    使用gRPC构建微服务（GRPC）
        做一个处理用户信息的微服务
        客户端通过用户名，可以从服务端查询用户的基本信息
            目录结构及文件含义
                gRPC
                    proto
                        user.proto      定义客户端请求，服务端响应的数据格式
                        user.pb.go      命令生成的，为数据交换提供的函数
                    server.go       微服务服务端
                    client.go       微服务客户端
            根据 proto 文件生成 .go 文件
            在user.proto文件右键，open in terminal，生成命令如下：
                protoc -I . --go_out=plugins=grpc:. ./user.proto
                命令注意：
                    因为用到了protoc的grpc的plugin，所以对应插件需要确保安装。
                        go get google.golang.org/protobuf/cmd/protoc-gen-go
                        go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
                    关于protoc-gen-go插件的配置请看：
                        https://blog.csdn.net/qq_39938666/article/details/124370388?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522168986592216800185870678%2522%252C%2522scm%2522%253A%252220140713.130102334.pc%255Fall.%2522%257D&request_id=168986592216800185870678&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~first_rank_ecpm_v1~rank_v31_ecpm-2-124370388-null-null.142^v90^control_2,239^v3^control&utm_term=%20protoc-gen-go%20%E6%8A%A5%E9%94%99&spm=1018.2226.3001.4187
                最终在proto目录下生成user.pb.go文件
            gRPC编写服务端和客户端
                代码详见: GRPC.server.go 和 GRPC.client.go
                运行 GRPC.server.go
                再运行 GRPC.client.go 查看结果










叁.go-micro
    
    简介：
        Go Micro是一个插件化的基础框架，基于此可以构建微服务，Micro的设计哲学是可插拔的插件化架构，
        它默认使用consul作为服务发现（2019年源码修改了默认使用mdns），通过http通信，通过protobuf和json进行编解码
        
    主要功能：
        服务发现：
        负载均衡：
        消息编码：
        请求响应：
        异步通信：
        可插拔接口：

    go-micro源码库：
        https://github.com/micro
    
    go-micro安装：
        在user.proto文件右键，open in terminal，生成命令如下：
            protoc -I . --micro_out=. --go_out=. ./hello.proto
            命令注意：
                需要安装 protoc-gen-micro 插件。
                通过 echo $GOPATH 查看go在本地的路径。
                一般命令文件都放在 $GOPATH/bin目录下，比如前面的 protoc-gen-go
                protoc-gen-micro 文件同样也会安装在这个路径下。
            安装命令：
                go get -u -v github.com/micro/micro
                go get -u -v github.com/micro/go-micro
                go get -u -v github.com/micro/protoc-gen-micro
            install是 安装到 GOPATH 的bin目录下
                go install github.com/micro/protoc-gen-micro

        最终在proto目录下生成hello.pb.go 和 hello.pb.micro.go 文件
        
    micro编写服务端和访问
        运行 MicroGo.server.go 文件
        运行成功后
        在命令行通过如下命令调用服务（因为是在命令行所以 引号需要转译）
            call 后面的 hello 就是 微服务的示例参数中给定的服务名：micro.Name("hello")
            micro call hello Hello.Info {\"username\":\"trump\"}
        最后命令行返回响应结果就是Info方法的返回值。
        
    go-micro案例：
        开启如下命令才可以通过postman进行http的接口访问。
            micro api --handler=rpc
        遇到环境问题没有写这个案例。
        
        
    


完结撒花!...