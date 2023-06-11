博学之, 审问之, 慎思之, 明辨之, 笃行之;
零、壹、贰、叁、肆、伍、陆、柒、捌、玖、拾;




零.GoLang简介
    
    Go是Google的Ken Thompson 开发的一种静态强类型，编译型语言，Go语言语法与C相近，但功能上有内存安全，GC（垃圾回收），结构形态及CSP-Style并发计算。
    go语言特点：
        背靠大厂，Google背书，可靠
        天生支持并发
        语法简单，容易上手
        内置runtime，至此垃圾回收
        可直接编译成机器码，不依赖其他库
        丰富的标准库
        跨平台编译
    Go语言的应用领域：
        服务器开发
        云平台开发
        区块链
        分布式系统
        网络编程
    Go官网：
        https://go.dev/dl/
        https://golang.google.cn/
    环境变量：
        自行在官网下载，并配置环境变量。
        查看配置：
            go env
                GOROOT：指go安装的根目录
                GOPATH：指go项目所在的路径，高版本的go项目已经不再依赖GOPATH来管理项目了，使用go mod来管理项目。
        修改配置：
            $env:GO111MODULE = "on"
            $env:GOPROXY = "http://goproxy.cn"
            或者：
            go env -w GO111MODULE=on
            go env -w GOPROXY=http://goproxy.cn













壹.Go常用命令（01CommonCommand）
    
    在控制台通过 go help 就可以看到go的所有命令。
    常用命令如下：
        install: 编译并安装包和依赖
            go install github.com/go-sql-driver/mysql
            在官网选择package会跳转到 https://pkg.go.dev/，这是go的包管理平台，各种库都可以在里面搜索。 
        get: 下载并安装包和依赖
            go get github.com/go-sql-driver/mysql
            在官网选择package会跳转到 https://pkg.go.dev/，这是go的包管理平台，各种库都可以在里面搜索。 
        build: 编译包和依赖
            go build main.go
            此时会生成main文件，在win下就是一个.exe的可执行文件，在控制台 ./main 即可运行结果。
        run: 编译并运行go程序
            go run main.go
        mod：模块管理
            go mod init commoncommand
            commoncommand是包名称
        clean: 移除对象文件
        list: 列出包
        env: 打印go的环境信息
        bug: 启动错误报告
        doc: 显示包或者符号的文档
        fix: 运行go tool fix来更新包
        fmt: 运行gofmt进行格式化
        generate: 从processing source生成go文件
        test: 运行测试
        tool: 运行go提供的工具
        version: 显示go的版本
        vet: 运行go tool vet
    编写go代码：
        ... here ...











贰.标识符&关键字&命名规则（）
    
    
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        