博学之, 审问之, 慎思之, 明辨之, 笃行之;
零、壹、贰、叁、肆、伍、陆、柒、捌、玖、拾;







零.Go微服务实践
    
    本章节的所有博客文档参考请看：https://www.mszlu.com
    
    在微服务领域Java有很好的时间如Spring cloud，Golang也有自己不同的实践方式
    
    标准库/自研派系：
        得益于go标准库的性能，稳定性。很多团队基于标准库来灵活的搭建微服务程序。
        这个是想来自于一个认知：不要让框架束缚开发。而Go标准库的强大，能够让开发者做到这一点，即使不用框架，也能开发出高效的应用。
        微服务的基础是通信，也就是RPC框架的选择，这一点上，可以选择grpc或基于grpc进行自研rpc框架。
        至于其他用到的组件，有需要的话进行集成即可。
        如果部署采用k8s，并且使用服务网格，比如lstio来处理，那么只需要关注业务逻辑即可，不需要再关心服务发现，熔断，流量控制，负载均衡等等。
    Web框架派系：
        比如 Gin + grpc 为核心，将其他组件集成进来的微服务架构，生产稳定，同样可以使用k8s + istio。
        最终构建的仍是现代化的云原生微服务架构。
    大一统框架：
        并不是所有的团队都有能力执行上述的两种方式，go同样也有像其他语言那样的大一统框架，可以达到快速开发的目的，比如：go-zero
        go-zero由国人开发，文档齐全，国内使用人数多，相对是一个不错的选择。
    
    Github地址：https://github.com/zeromicro/go-zero
    官网地址：https://go-zero.dev/
    
    go版本：1.18.1
    goctl安装：
        地址：https://go-zero.dev/docs/tasks/installation/goctl
        goctl（go control）是go-zero微服务框架下的代码生成工具，可以提升开发效率，使得开发将时间专注于业务开发上，其功能有：
            api服务生成
            rpc服务生成
            model代码生成
            模版管理
        go install github.com/zeromicro/go-zero/tools/goctl@1.4.4
        
        protoc & protoc-gen-go 安装：
            goctl提供了便捷的安装方式：
                goctl env check -i -f --verbose
    
    名次解释：
        mono：单体服务（Monolithic application）
        micro：微服务（Microservices）
    
    单体应用Demo：(monolithic)
        单体应用构建：
            goctl api new monolithic
                通过gotcl构建项目
            cd monolithic
            go mod tidy
                下载依赖
        项目结构：
            https://go-zero.dev/docs/concepts/layout
        服务启动：
            命令启动：
                go run ./hello.go -f ./etc/monolithic-api.yaml
            编辑器集成环境启动：
                右键debug运行
    
    微服务应用Demo：(microservice)
        假设有两个服务：
            order service
            user service
        用户在查询订单时，同时需要获取用户信息。
        微服务应用构建：
            mkdir mall
                创建商城的项目（目录）
            cd mall
            goctl api new order
                构建order服务
            goctl api new user
                构建user服务
            go work init
                用workspace模式初始化项目
            go work use order
            go work use user
                将order和user服务加到工作区中
            user服务整改：
                创建rpc目录
                新建proto文件：user.proto
                通过 gotcl rpc 指令生成proto的go代码
                    goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
                然后将rpc目录下生成的目录和文件放置到服务根目录下，将根目录下的原目录和user.go入口文件都删除，因为rpc已经生成了新的user.go入口文件。
                go mod tidy
                    重新加载user.go入口文件的依赖
                编辑 logic.getuserlogic.go.GetUser 的逻辑编写文件的GetUser方法
                    return &user.UserResponse{
                        Id:     in.Id,
                        Name:   "leechenze",
                        Gender: "man",
                    }, nil
            order服务整改:(order服务是一个对外提供的服务，不需要user那么多rpc的操作了)
                编辑order.api自动生成代码
                    详见order.api
                goctl api go -api order.api -dir ./gen
                    将生成的代码放在gen目录下。也可直接 goctl api go -api order.api -dir ./ 生成
                    如果有同名文件不会覆盖，只会将要生成的代码片段生成到文件中。
                修改 etc.order-api.yaml
                修改 config.config.go
                修改 svc.servicecontext.go
                修改 logic.getorderlogic.go
                    以上修改执行跟代码吧，挺好理解，但是关联性比较杂乱。
            为order-api.yaml中的Etcd配置启动docker
                根目录创建：docker-compose.yaml
                在项目根目录创建.env环境变量文件
                启动docker：
                    docker-compose up -d
                启动user服务8080
                启动order服务8888
                访问：http://localhost:8888/api/order/get/1
                    {
                        "id": "1",
                        "name": "hello order name",
                        "userName": "hello user name"
                    }
                








壹.集成mysql数据库
    
    在user模块中集成mysql，数据库调用，对order模块暴露服务
    新建 internal/model/user.sql 文件
        CREATE TABLE `user`
        (
            `id`     bigint(0) NOT NULL AUTO_INCREMENT,
            `name`   varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
            `gender` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
            PRIMARY KEY (`id`) USING BTREE
        ) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;
    在model目录下打开终端根据 user.sql 文件生成一些CURD的相关代码
        goctl model mysql ddl --src user.sql --dir .
        此时根据user.sql 就生成了 usermodel.go, usermodel_gen.go, vars.go三个文件
    新建model/user.go     存放user的模型代码
    新建repo/user.go      存放user的接口文件
    新建dao/user.go       存放user的接口实现类的文件
    新建rpc/gen_rpc.sh    经常执行的命令可以创建 .sh 后缀的批处理文件，直接在rpc目录下 ./gen_rpc.sh 即可执行其中的命令
        goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
    chmod +x ./gen_rpc.sh   注意还需要给处理文件加权限才能执行，否则权限被拒绝，sudo模式也不行。
    
    由于user的grpc服务没法直接进行调用，所以新建一个userapi服务对user服务进行调用
        goctl api new userapi
            新建 userapi 服务
        go work use ./userapi/
            将 userapi 服务加到工作目录中
        cd userapi
        new Go Modules File
            新建go.mod文件 或 go mod init
        go mod tidy
    再创建一个rpc-common服务，这个服务什么也没有，只是存放一些rpc相关的公共引用包。
        步骤如上，处理哪些地方详见代码。
    其实 userapi 这个服务和 order服务是大同的，只不过用的是post请求，抽离出来有些杂乱无章。请详细跟踪代码吧。
    最后在postman访问：localhost:8888/userapi
    post请求，body体中 raw 传参如下：
        {
            "name": "test1111",
            "gender": "man"
        }
    成功返回如下信息：
        {
            "message": "success",
            "data": {
                "id":"",
                "name": "test1111",
                "gender": "man"
            }
        }
    
    
    







零、壹、贰、叁、肆、伍、陆、柒、捌、玖、拾;