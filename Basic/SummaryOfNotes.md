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
        *：根据地址取值和声明为指针的作用
    
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
        




        
        
        
        
        



零、壹、贰、叁、肆、伍、陆、柒、捌、玖、拾;
        
        
        
        
        