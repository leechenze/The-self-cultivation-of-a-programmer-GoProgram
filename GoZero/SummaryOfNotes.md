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
    最后在postman访问：localhost:8888/register
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









贰.集成Redis缓存
    
    go-zero框架使用的是go-redis, 在go-zero包中本来就集成了go-redis, 所以可以直接用，不用再导包了。
    user模块:
        添加配置：
            user/internal/cofnig/config.go
                CacheRedis cache.CacheConf
            user/internal/etc/user.yaml
                CacheRedis:
                    - Host: 127.0.0.1:6379
                    - Type: node
        配置缓存连接：
            user/database/sqlx.go
        接口中定义FindById方法
            internal/repo/user.go
        实现FindById方法
            internal/dao/user.go
        启动redis服务端口为默认的6379
        启动user服务
    userapi模块：
        添加路由： 
            internal/handler/routes.go
                {
                    Method:  http.MethodGet,
                    Path:    "/user/get/:id",
                    Handler: userapiHandler.GetUser,
                },
        声明并实现GetUser方法和逻辑：
            internal/handler/userapihandler.go
            internal/logic/userapilogic.go
        启动userapi服务并访问测试：
            localhost:8888/user/get/4
        查看redis缓存，每访问一次都将在Redis中成功缓存。










叁.集成JWT中间件
    
    流程：
        先登录，获取token，前端会存储token，下一次请求在header中携带token
    
    userapi模块：
        添加配置：
            etc/userapi-api.yaml
                Auth:
                    AccessSecret: "~!@#$%^&*()"
                    AccessExpire: 604800
                解读：AccessSecret是一个自定义的验证密钥,AccessExpire是过期时间为七天
            internal/config/config.go
                Auth    struct {
                    AccessSecret string
                    AccessExpire int64
                }
        添加登录路由和登录逻辑：
            internal/handler/routers.go
            internal/handler/userapihandler.go
            internal/logic/userapilogic.go
        修改userapi.api并生成go代码
            userapi.api
                goctl api go -api userapi.api -dir ./gen
            注意：这里只是生成的代码只是看下goctl生成的jwt代码是如何写的和写在哪的，最好不要拿来即用。
        然后通过userapi.api模版生成的代码参考，jwt的配置需要写到routes.go中
            []rest.Route{
                {
                    Method:  http.MethodGet,
                    Path:    "/user/get/:id",
                    Handler: userapiHandler.GetUser,
                },
            },
            rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
        启动user和userapi服务进行测试：
            postman测试：localhost:8888/user/get/3，如果返回状态码为401表示已经初步成功认证了。
            然后访问login接口，返回jwt字符串密钥并复制
            再次访问/user/get/3，在Headers请求头中携带这段token，添加：Authorization: login接口生成的jwt密钥。
            此时状态码不再是401，且得到的正确的返回结果。











肆.GoZero中间件集成

    userapi模块：
        修改路由，添加中间件：
            internal/handler/routes.go
        新建middlewares中间件目录：
            internal/middlewares/user.go
        添加中间件的资源池（service层）的依赖
            internal/svc/servicecontext.go







 

伍.自定义错误

    userapi模块：
        在GetUser逻辑中模拟一个错误
            internal/logic/userapilogic.go
        自定义错误包
            internal/customError/error.go
        在userapi.go入口文件中使用自定义错误
            httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (httpStatus int, errorData any) 
        重启服务测试：
            postman访问：localhost:8888/user/get/1，因为在GetUser的逻辑处理中只对id为1的数据做了自定义错误测试
            返回结果即customError.BizError的格式
                {
                    "code": 500,
                    "msg": "参数错误"
                }
    （额外知识点）生成的模版格式修改：
        go-zero生成代码是基于模版生成的，如果需要自定义或者生成的代码格式不符合期望，也可以修改模版代码
        操作步骤：
            goctl template init
                可以返回你本地模版的位置在哪里，然后打开这个位置执行修改即可。
        详细请参考官网地址：
            https://go-zero.dev/docs/tutorials/cli/template










陆.Goctl详解
    
    Goctl命令详见官网：https://go-zero.dev









柒.集成日志
    
    在Gozero中使用最多的第三方日志库：
        zap：    https://github.com/zeromicro/zero-contrib/blob/main/logx/zapx/zap.go
        logrus： https://github.com/zeromicro/zero-contrib/blob/main/logx/logrusx/logrus.go
    这两个日志库相比zap用的更多一些，这里在userapi模块集成zap日志
    userapi模块：
        新建日志包：
            userapi/zapx/zap.go
                将上面给出的zap.go的github地址的代码直接拷贝粘贴。
        启动类中设置日志的writer
            userapi/userapi.go
                logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
                writer := NewZeroLogWriter(logger)
                logx.SetWriter(writer)
        在userapilogic.go中测试日志输出
            userapi/internal/logic/userapilogic.go
                查看之前就使用过的 logx.Infof 的日志输出，启动user和userapi服务
                访问登录接口和获取用户接口在控制台查看日志
                localhost:8888/user/get/1
                可以在zapx/zap.go的Info方法打个断点，以证明 logx.Infof 走的是zap包中的方法。









捌.集成监控
    
    在微服务开发中，监控是一件非常重要的事情，很多线上问题都需要通过监控来出发告警，从而进行即使处理。
    Prometheus是目前应用最广泛，使用最多的监控中间件。
    
    docker-compose.yaml添加配置
          prometheus:
            container_name: prometheus
            image: bitnami/prometheus:2.40.7
            environment:
              - TZ=Asia/Shanghai
            privileged: true
            volumes:
              - ${PRO_DIR}/prometheus.yml:/opt/bitnami/prometheus/conf/prometheus.yml  # 将 prometheus 配置文件挂载到容器里
              - ${PRO_DIR}/targets.json:/opt/bitnami/prometheus/conf/targets.json  # 将 prometheus 配置文件挂载到容器里
            ports:
              - "9090:9090"                     # 设置容器9090端口映射指定宿主机端口，用于宿主机访问可视化web
            restart: always
    prometheus.yml和targets.json都在prometheus目录下
    targets.json是对当前进行监控的服务的描述
    docker-compose重新启动以加载prometheus的配置
        docker-compose up -d
    在userapi/etc/userapi-api.yaml添加配置
        Prometheus:
            Host: 127.0.0.1
            Port: 9081
            Path: /metrics
    在user/etc/user.yaml添加配置
        Prometheus:
            Host: 127.0.0.1
            Port: 9091
            Path: /metrics
    访问prometheus的web可视化界面：(选择Status/Targets)
        http://localhost:9090/targets?search=
        注意：docker容器内部访问外部的地址就使用 host.docker.internal:8080，host.docker.internal用以指代外部的IP地址.
        选择graph，输入Expression: {app="user-api"}
        即可查看app名为 user-api 的服务请求。user-api就是在targets.json中配置的服务名称











玖.集成Jeager

    Jeager是一个链路追踪的中间件，说白了就是方便查看某个服务在某个时刻具体调用了某个服务，
    
    docker-compose.yaml添加配置(jeager要依赖ElasticSearch，所以也需要集成下es的配置)
        jaeger:
            container_name: jaeger
            image: jaegertracing/all-in-one
            environment:
              - TZ=Asia/Shanghai
              - SPAN_STORAGE_TYPE=elasticsearch
              - ES_SERVER_URLS=http://elasticsearch:9200
              - LOG_LEVEL=debug
            privileged: true
            ports:
              - "6831:6831/udp"
              - "6832:6832/udp"
              - "5778:5778"
              - "16686:16686"
              - "4317:4317"
              - "4318:4318"
              - "14250:14250"
              - "14268:14268"
              - "14269:14269"
              - "9411:9411"
            restart: always
        elasticsearch:
            container_name: elasticsearch
            image: elasticsearch:7.13.1
            environment:
                - TZ=Asia/Shanghai
                - discovery.type=single-node
                - "ES_JAVA_OPTS=-Xms512m -Xmx512m"
            privileged: true
            ports:
                - "9200:9200"
            restart: always
    在userapi/etc/userapi-api.yaml添加配置
        Telemetry:
            Name: user-api
            Endpoint: http://localhost:14268/api/traces
            Sampler: 1.0
            Batcher: jaeger
    在user/etc/user.yaml添加配置
        Telemetry:
            Name: user-rpc
            Endpoint: http://localhost:14268/api/traces
            Sampler: 1.0
            Batcher: jaeger
    重启userapi和user服务
    访问jeager的web可视化界面：http://localhost:16686
        首次打开是没有service可以监控的，需要先请求两次，然后在刷新jeager查看service就有了，
        然后选择user-rpc，点击find traces 即可查看请求的访问信息和信息的跟踪链路











拾.分布式事务
    
    ... TODO ...





零、壹、贰、叁、肆、伍、陆、柒、捌、玖、拾;