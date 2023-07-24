博学之, 审问之, 慎思之, 明辨之, 笃行之;
零、壹、贰、叁、肆、伍、陆、柒、捌、玖、拾;







零.Go微服务实践
    
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
    
    微服务应用Demo：(microservices)
        ... TODO ...
        
    
    
    
    
    
        
