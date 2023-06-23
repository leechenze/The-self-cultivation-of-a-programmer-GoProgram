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
        代码组织：
            go应用使用包和模块来组织代码，包对应到文件系统就是文件夹，模块就是 **.go** 的源文件，一个包中会有多个模块或者多个子包。
        go项目管理工具：
            早起的go项目使用gopath来管理项目，不方便而且容易出错，从golang1.1开始，使用gomod管理项目，当然还有第三方模块例如govender。
        实现步骤：
            创建项目
                mkdir hellogo
            初始化项目
                cd hellogo
                go mod init hellogo
                    执行完后就会生成一个go.mod文件。
            创建包
                mkdir user
                cd user
                touch user.go
                    package user
                    func Hello() string {
                        return "hello"
                    }
            创建模块&相互调用
                cd ..
                touch main.go
                    // main是go应用的入口程序。
                    package main
                    import (
                        "fmt"
                        "hellogo/user"
                    )
                    func main() {
                        s: = user.Hello()
                        fmt.Printf("s: %v\n",s)
                    }
            
        










贰.标识符&关键字&命名规则（02keywordsAndIdentifiers）
    
    标识符：
        identifiers通俗地讲，就是给变量，常量，函数，方法，结构体，数组，切片，接口起名字。
        组成规则：
            标识符由于数字，字母和下划线组成。
            只能以下划线和字母开头。
            标识符区分大小写。
    关键字：
        go语言的关键字大概有25个左右，如下所示：
            程序声明：import、package

            程序实体声明和定义：chan、const、func、interface、map、struct、type、var
        
            程序流程控制：go、select、break、case、continue、default、defer、else、fallthrough、for、goto、if、range、return、switch

        除了以上介绍的这些关键字，Go语言还有36个预定义的标识符，其中包含了基本类型的名称和一些基本的内置函数，如下所示：
            append	    bool	    byte	    cap	    close	    complex
            complex64	complex128	uint16	    copy	false	    float32
            float64	    imag	    int	        int8	int16	    uint32
            int32	    int64	    iota	    len	    make	    new
            nil	        panic	    unit64	    print	println	    real
            recover	    string	    true	    uint	uint8	    uintprt
    命名规范：
        命名规则涉及变量，常量，全局函数，结构，接口，方法等命名。Go语言从语法层面进行了以下限定：
            任何需要对外暴露的名称以大写字母开头，不需要对外暴露的则应该以小写字母开头。
            当大写命名的标识符可以被外部包的代码所使用，类似面向对象语言中的 public
            命名如果以小写字母开头，则对包外是不可见的，但是它们在整个包的内部是可见并可用的，类似面向对象语言中的 private
        包名称：
            保持package的名字和目录名一致，包名应该为小写单词，不要使用下划线或者混合大小写，命名也不要和标准库冲突。
            如：
                package dao
                package service
        文件命名：
            应该为小写单词，使用下划线分隔各个具有意义的单词。
            如：
                customer_dao.go
        结构体命名：
            应该采用驼峰命名，首字母根据访问控制大写或小写。
        接口命名：
            命名规则和结构体类似，单个函数的结构名以 er 作为后缀，例如：Reader，Writer，例如 Java 中的接口以 I 作为前缀类似。
        变量命名：
            变量名一般遵循驼峰法。首字母同样根据访问控制大写或小写
        常量命名：
            常量一般全部使用大写字母组成，并使用下划线分隔有意义的单词。











叁.变量&常量（03VariableAndConstant）
    
    声明变量：(var)
        Go语言中的变量需要声明后才能使用，同一作用域内不支持重复声明。并且Go语言的变量声明后必须使用。
        声明一个后续不用的变量就会报错，可以将变量名改为 _ 表示匿名变量，以示后续没有用到。
        语法：
            var identifier type
        例如：
            var name string
            var age int
            var isOk bool
        
    声明常量：(const)
        常量就是在程序编译阶段就确定下来的值，而程序在运行时则无法改变该值。
        与变量定义语法相同，不同的是常量在声明时就必须赋值，因为常量定义后就不能再修改了。
        
        特殊常量值：iota
            iota比较特殊，可以被认为是一个可被编译器修改的常量，它默认的初始值是0，每调用一次加1，遇到const关键字时被重置为0。
                例如：
                    const (
                        I1 = iota
                        I2 = iota
                        I3 = iota
                    )
                    println(I1, I2, I3) // 0 1 2
            使用 _ 跳过某些值
                例如：
                    const (
                        I11 = iota
                        _
                        I22 = iota
                    )
                    println(I11, I22) // 0 2
            iota中间插队
                例如：
                    const (
                        I111 = iota
                        I222 = 999
                        I333 = iota
                    )
                    println(I111, I222, I333) // 0 999 2













肆.数据类型（04DataType）
    
    类型和描述：
        布尔型：bool
            注意点：go中不能使用0或非0表示布尔值，在C和JS中是可以这样表示的。
        数字类型：
            整型 int 和 浮点型 float32，flost64。
            除此之外Golang还有基于架构的类型，例如：int，unit 和 uintptr等等。
        字符串类型：string
            Go语言的字符串的字节使用 UTF-8编码标识Unicode文本。
        派生类型：
            指针类型：Pointer
            数组类型：Array
            结构化类型：struct
            通道类型：Channel
            函数类型：func
            切片类型：slice
            接口类型：interface
            字典类型：Map
    Go格式化输出
        详见：formatOutput.go













伍.运算符&流程控制(05OperatorsAndProcessControl)

    加减乘除：
        a := 100
        b := 20
        println("(a + b)", (a + b))
        println("(a - b)", (a - b))
        println("(a * b)", (a * b))
        println("(a / b)", (a / b))
        println("(a % b)", (a % b))

    注意：++（自增）和 --（自减）在go语言中是单独的语句，并不是运算符。
        如下会报错：
            c := 1
            d := c++
        如下则可以，只要不参与运算：
            c := 1
            c++
    关系运算符：
        println("(a > b)",(a > b))
        println("(a < b)",(a < b))
        println("(a >= b)",(a >= b))
        println("(a <= b)",(a <= b))
        println("(a == b)",(a == b))
        println("(a != b)",(a != b))
    逻辑运算符：
        d:=true
        e:=false
        println((d && e))
        println((d || e))
        println((!d))
        println((!e))
    位运算符：
        f:=4 // 二进制 100
        fmt.Printf("a: %b\n", f)
        g:=8 // 二进制 1000
        fmt.Printf("a: %b\n", g)
        fmt.Printf("(f & g): %v %b \n",(f&g),(f&g))
        fmt.Printf("(f | g): %v %b \n",(f|g),(f|g))
        fmt.Printf("(f ^ g): %v %b \n",(f^g),(f^g))
        fmt.Printf("(f << g): %v %b \n",(f<<g),(f<<g))
        fmt.Printf("(f >> g): %v %b \n",(f>>g),(f>>g))
    赋值运算符：
        h:=100
        h = 200
        println("h",h)
        h += 1
        println("h += 1",h)
        h -= 1
        println("h -= 1",h)
        h *= 1
        println("h *= 1",h)
        h /= 1
        println("h /= 1",h)
    流程控制：
        go语言中只for 和 for range循环，去除了while，do while。
            1。for循环
                for i := 0; i < 10; i++ {
                    println(i)
                }
            2。for range循环
                x := [...]string{"lincoln", "trump", "obama"}
                for ind, val := range x {
                    println(ind, val)
                }
        初始变量可以声明在布尔表达式里面，但是作用域只在当年if else的花括号之内。
            if age := 20; age > 18 {
                println(age)
            } else {
                println("18")
            }
        不能使用0或非0表示真假，如下是错误示例：
            a := 100
            if a {
                println(a)
            }
    流程关键字：
        break：单独在switch中使用break和不使用没有什么区别，它在switch中是默认项。
        continue：在go中只能用在for循环中，意义和其他语言相同，终止本地循环，进行下次循环。break是终止循环，continue只是跳出一次循环。
        fallthrough：和break情况正好相反，如果想在switch中执行多个case条件，就在每个case中添加这个关键字。
        goto：可以跳出双层循环和多层循环，break只能跳出当前循环。同时和break相同都能跳转到指定标签（这就是跳出多层循环的原理,通过goto指定标签跳转多层循环）











陆.数组&切片(06ArrayAndSliceAndMap)
    
    数组
        go数组一旦定义数组长度就不可修改
        详见：array/array.go
    切片
        切片就是一个可变长度的数组，它底层就是使用数组实现的，增加了自动扩容的功能。
        详见：slice/slice.go
        切片初始化
            切片的初始化方法有很多，可以直接初始化，也可以使用数组初始化等。
        切片切割
            切片切割和遍历都和数组相同。
            切片切割，功能就是JS数组的slice方法，数组和切片都可以进行如下的切割操作
                // var a5 = [...]int{1, 2, 3, 4, 5, 6, 7, 8}
                var a5 = []int{1, 2, 3, 4, 5, 6, 7, 8}
                a6 := a5[:3]
                a66 := a5[0:3]
                fmt.Printf("a6 %v\n", a6)
                fmt.Printf("a66 %v\n", a66)
                a7 := a5[3:]
                fmt.Printf("a7 %v\n", a7)
                a8 := a5[2:5]
                fmt.Printf("a8 %v\n", a8)
                a9 := a5[:]
                fmt.Printf("a9 %v\n", a9)
        切片的CRUD操作
            切片删除的公式：
                index就是要删除的那个下标，这个 ... 就是展开操作符：将数组/切片中的每一个元素展开为一个个的个体。
                slice = append(slice[:index],slice[index+1:]...)
            切片拷贝，如果直接将一个切片赋值给另一个切片，那么这两个切片指向同一个地址，修改一个切片后另一个切片也会被修改。
            这个问题可以用拷贝函数来处理，拷贝值要用make创建
                a14 := []int{1, 2, 3, 4, 5}
                a15 := make([]int, len(a14))
                copy(a15, a14)
                fmt.Printf("a14修改前 %v\n", a14)
                fmt.Printf("a15修改前 %v\n", a15)
                a14[0] = 100
                fmt.Printf("a14修改后 %v\n", a14)
                fmt.Printf("a15修改后 %v\n", a15)
    map
        详见：gomap/gomap.go











柒.函数（07Function）
    
    C++和Java中类是一级公民，go和js一样函数是一级公民。
    函数特性：
        三种函数：普通函数，匿名函数，方法（定义在struct上的函数）
        go中不允许函数重载，也就是不能函数同名。
        go中的函数不能嵌套函数，可以嵌套匿名函数
        函数可以作为参数传递给另一个函数
        函数的返回值可以是一个函数
        函数的传参是一个拷贝的过程
        函数参数可以没有名称，即匿名函数
    函数传参需要注意：
        map，slice，interface，channel这些指针类型的，拷贝传值也是拷贝的指针地址
        所以在函数中修改这种指针类型的参数会影响外部的值。
    
	函数类型声明
        type fn func(int, int) int
        var foo fn
        foo = sum
	    fmt.Printf("foo(2,3) %v\n", foo(2, 3))
    
    高阶函数
        高阶函数就是将一个函数作为另一个函数的参数进行传递，并且返回一个函数。
        
    匿名函数
        匿名函数既可以没有参数也可以没有返回值
        
    闭包
        闭包是将函数内部和函数外部连接起来的桥梁，闭包就是一个函数内部有另外一个函数并将其作为结果返回
        
    递归
        递归实现5的阶乘，正确结果为120
            func fiveFactorial(a int) int {
                if a == 1 {
                    return 1
                } else {
                    return a * fiveFactorial(a-1)
                }
            }
    defer
        defer关键字会将后续的语句执行逆序处理，也就是说同一个作用域下，最先被defer定义的最后被执行，最后被defer定义的先执行。
        没有被defer定义的都正常执行
    
    init函数
        go中又一个特殊的函数 init函数，先于main函数执行，实现包级别的一些初始化操作
        init特点：
            init函数先与main自动执行，且和main一样是主函数，不能被其他函数调用
            init函数没有参数和返回值
            每个包可以有多个init函数
            包的每个源文件也可以有多个init函数
            不同包的init函数按照包导入的依赖关系决定执行顺序的
        初始化顺序：
            初始化变量 > init > main
            详见：initFunc/initFunc.go








捌.指针&结构体(08PointerAndStruct)
    
    golang中的指针不能进行偏移和运算
    golang中的指针操作非常简单：
        &：取地址
        *：根据地址取值 和 声明表示为指针的作用
    
    数组指针
        c语言中的数组指针是指向数组的，而因为go的指针不能运算，所以go语言的指针是指向数组中的每一个元素的。
        声明语法如下：
            var ap [3]*int
            为什么是 [3]*int，而不是 *[3]int ？
                就是因为go语言中的指针不能运算，是给数组的每一个具体值赋予指针，所以要写成 [3]*int
                而 *[3]int 就成了给数组赋予指针了，用心思考领悟一下这种思想。
    
    在了解结构体指针之前先了解一下类型定义的知识
    类型定义与类型别名
        类型定义：type MyInt int
        类型别名：type MyInt = int
        区别：
            类型定义相当于定义了一个全新的类型
            类型别名并没有定义一个新的类型，而是把一个类型重新换个名，实际类型还是原来的类型而不是类型的别名
            类型别名只会在代码中存在，在编译完后不会存在该别名
            类型别名和原来的类型是一致的，所以原来类型所拥有的方法，类型的别名中也可以调用。
        
    结构体指针
        go语言中没有面向对象和类的概念了，取而代之的是C语言中的结构体，通过结构体来描述一些类特性：继承，多态，封装，组合等。
        结构体也要用type关键字来定义。
        
    结构体作为函数参数
        结构体作为函数参数（值传递，并不会改变外部Person的值。）
            func changePerson(person Person) {
                person.id = 105
                person.name = "nixon"
                fmt.Printf("nixon, %v\n", person)
            }
        
        结构体作为函数参数（引用传递/指针传递，将会改变外部Person的值。）
            func changeperson2(person *Person) {
                person.id = 107
                person.name = "nixon"
                fmt.Printf("nixon, %v\n", person)
            }
    golang嵌套结构体
    golang结构体中的方法
        结构体中的方法是一个特殊的函数，定义与结构体之上，被称为结构体的接收者（reciver），通俗的讲：方法就是有接收者的函数
        相比Java中的类，最特别的地方就是golang结构体的属性和方法分开写了，方法是定义在结构体花括号之外的。
        语法如下：
            func (recv struct) method_type(param) return_type {}
        注意事项：
            方法的接收类型并非一定是结构体，也可以是type定义的类型别名，slice，map，channel，func类型等都可以。
            struct结合它的方法就等价面向对象中的类，只不过方法是声明到结构体的花括号之外了，并非一定要属于同一个文件，但必须属于同一个包。
            结构体方法的接收类型既可以是结构体，类型，slice，map等第一条提到过的值类型之外，还可以接收指针类型。
            方法就是函数，所以go中没有方法重载（overload）的说法，也就是说类型中的所有方法名都必须是唯一的。
    golang中方法的接收者类型
        值类型：不会改变结构体花括号外部的值,因为是值传递
        指针类型：会改变结构体花括号外部的值,因为是指针传递/引用传递。











玖.接口(09Interface)
    
    接口和结构体一样应该以type定义，当然用var也可以声明。
    按规定，接口通常以er结尾。
    
    接口和类型的关系：
        一个类型可以实现多个接口
        多个类型可以实现同一个接口（多态）
    接口嵌套
    通过接口实现OCP设计原则（Open-Close-Principle）
        面向对象的可复用设计的第一块基石，便是所谓的"开-闭"原则，虽然，go不是面向对象语言，但也可以模拟实现这个原则
        原则即：对扩展是开放的，对修改是关闭的。














拾.继承(10inherit)
    
    继承
        golang中的继承是通过结构体嵌套实现的。
    构造函数
        golang中的构造函数的功能是可以通过方法模拟的。












拾壹.包(11packet)

    包用以区分命令空间（一个文件夹中不能有两个同名的包），所以可以更好的管理项目
    go中使用 package 关键字声明包，通常文件夹名称和包名称相同，并且同一个文件下只有一个包。
    包管理工具：go modules
        go modules是golang 1.11 版本新加的特性，用来管理模块中包的依赖关系
        使用方法：
            初始化模块：
                go mod init <项目模块名称>
                项目模块名称 一般规则是 域名+项目名
            依赖关系处理，根据go.mod文件
                go mod tidy
            将依赖包复制到项目下的vendor目录
                go mod vendor
            显示依赖关系
                go list -m all
                go list -m --json all
            下载依赖
                go mod download [path@version]
                [path@version]是非必写的
    远程包引用：
        https://go.dev/
            进入导航栏package
        这里就以gin框架为例进行演示：（详细步骤会在官网中有）
            安装Gin
                go get -u github.com/gin-gonic/gin
            下载Gin程序所需要的依赖：
                将一下代码片段放到main程序中。
                    // gin application
                    r := gin.Default()
                    r.GET("/ping", func(c *gin.Context) {
                        c.JSON(http.StatusOK, gin.H{
                            "message": "pong",
                        })
                    })
                    r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
                导入Gin包（"github.com/gin-gonic/gin"）
                最后执行 tidy 以下载gin程序执行所需要的依赖：（tidy后编辑器中的错误会自动消失）
                    go mod tidy












拾贰.并发编程(12ConcurrentProgramming)

    协程：（goroutines）
        golang中的并发是函数相互独立运行的能力，goroutines是并发运行的函数，Golang提供了Goroutines作为并发处理操作的一种方式。
        创建一个协程非常简单，就是在一个任务函数前添加一个go关键字。

    通道：（goChannel）
        go提供了一种称为通道（channel）的机制，用于协程之间（goroutines）共享数据，当作为goroutines执行并发活动时，需要在goroutines之间共享资源或数据
        通道充当了goroutines之间的管道，并提供了一种机制来保证同步交换
        根据数据交换行为，通道会分为两种类型：
            无缓冲通道和缓冲通道。
            无缓冲通道用于执行goroutines之间的同步通信，而缓冲通道用于执行异步通信。
            无缓冲通道保证在发送和接收的瞬间执行两个goroutines之间的交换，缓冲通道没有这样的保证。
        通道的发送和接收特性：
            对于同一个通道，发送操作之间是互斥的，接收操作之间也是互斥的。
            发送和接收的操作对元素值的处理都是不可分割的
            发送操作在完全完成之前会被阻塞，接收操作也是如此。（即加锁，放置其他的协程切入）
        语法：
            通道声明：
                var gochannel = make(chan int)
            通道发送：
                gochannel <- value
            通道接收：
                value := <-gochannel
    
    WaitGroup同步：（goWaitGroup）
        WaitGroup在golang中用于线程同步。
        语法：
            WaitGroup声明：
                var wg sync.WaitGroup
            开启一个要等待的子协程：（进入当前子协程的同步等待）
                在 go 协程关键字的上下声明都可以
                    wg.Add(1)
                    go showMsg(i)
            关闭一个要等待的子协程：（退出当前子协程的同步等待）
                wg.Done() // or wg.Add(-1)
            执行等待：
                在需要阻塞执行的地方很执行
                    wg.Wait()
        
    RunTime：(goRunTime)
        runtime里面定义了一些协程管理相关的API。
            GoSched：
                让出CPU时间片，重新等待安排任务（让其他协程先运行）
            Goexit:
                退出协程
            GOMAXPROCS：
                设置最大的CPU核心数
            NumCPU：
                查看CPU的核心数
    
    Mutex互斥锁同步：（goMutex）
        语法：
            声明Metux锁：
                var lock sync.Mutex
            加锁：
                lock.Lock()
            解锁：
                lock.Unlock()
        解释：
            加锁和解锁放在要执行的协程中，在锁之间的程序会被锁定，其他协程不会此段程序执行完之前切进来执行。

    通道遍历：（goChanIterate）
        在通道完成之后一定要调用 close(chan) 关闭通道，否则就会产生死锁的问题而报错
    
    Select：（goSelect）
        select是go中的一个控制结构，类似于switch语句，用于处理异步IO操作。select会监听case语句中channel的读写操作
        当case中的channel读写操作为非阻塞状态（能读写）时，将会触发相应的动作。
        注意：
            select中的case语句必须是一个channel操作。
            select中的default子句总是可运行的。
        如果有多个case都可以运行，select会随机选择一个case执行，其他不会执行。
        如果没有可运行的case语句，且有default语句时，就执行default的动作
        如果没有可运行的case语句，且没有default语句时，select将被阻塞，直到某个case通信可以运行。
    
    Timer：（goTimer）
        timer顾名思义就是定时器的意思，可以实现一些定时操作，内部是通过channel实现的。
        
    Ticker：（goTicker）
        timer只执行一次，ticker可以周期的执行。
        
    原子操作：（goAtomic）
        atomic提供的原子操作能够确保任意时刻只有一个goroutines对变量进行操作，善用atomic能够避免程序中出现大量的锁操作。
        atomic常见的操作有：
            增减
            载入
            存储
            比较并交换cas
            交换










拾叁.标准库OS（13OSLibrary）
    
    ... TODO ...
    

























零、壹、贰、叁、肆、伍、陆、柒、捌、玖、拾;
        
        
        
        
        