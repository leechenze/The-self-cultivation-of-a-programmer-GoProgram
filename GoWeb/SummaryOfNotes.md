博学之, 审问之, 慎思之, 明辨之, 笃行之;
零、壹、贰、叁、肆、伍、陆、柒、捌、玖、拾;




零.初识GoWeb

    main.go    
        package main
    
        import (
        "fmt"
        "io/ioutil"
        "net/http"
        )
        
        func sayHello(w http.ResponseWriter, r *http.Request) {
        
            file, _ := ioutil.ReadFile("./hello.htm")
        
            _, _ = fmt.Fprintf(w, string(file))
        }
        
        func main() {
        http.HandleFunc("/hello", sayHello)
        
            err := http.ListenAndServe(":9090", nil)
            if err != nil {
                fmt.Printf("http serve failed, err %v\n", err)
                return
            }
        
        }
    
    hello.htm
        <!doctype html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport"
                  content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
            <meta http-equiv="X-UA-Compatible" content="ie=edge">
            <title>Document</title>
        </head>
        <body>
        <h1 style="color: orange">hello golang</h1>
        </body>
        </html>











壹.初识Gin框架(GinFirstApp)
    
    官网地址：https://pkg.go.dev/github.com/gin-gonic/gin
    参考博客地址：https://liwenzhou.com/posts/Go/gin/

贰.Go模版引擎(GoTemplate)

叁.Gin模版渲染(GinTemplate)

肆.GinJson(GinJson)

伍.Gin参数交互(GinParams)

    get参数请求
    form参数请求（post）
    Uri路径参数
    参数绑定(常用){ShouldBind}

陆.Gin文件上传(GinFileUpload)

柒.Gin重定向(GinRedirect)

捌.Gin路由和路由组(GinRouter)

玖.Gin中间件(GinMiddleware)

    Gin中的中间件说白了就是 在路由函数处理之前，执行的一个函数就是中间件，一般将一些公共的逻辑都写到中间件中，比如：权限校验，登录认证，日志记录等...

    功能上相当于SpringBoot的拦截器，而且和Node的Koa的中间件非常相似，都是遵循洋葱模型。
    
    和KOA的洋葱模型同理，看下面例子就明白了（m1ware是中间件函数，handlerFunc是路由执行函数）：
        m1ware --- start
        m2ware --- start
        GET handlerFunc run
        m2ware --- end
        m1ware --- end
    
    Gin的中间件常用方式是结合路由和路由组一起使用，

    默认路由注意：
        gin.Default() 默认使用了 Logger 和 Recovery 中间件，其中：
            Logger中间件的作用就是日志打印
            Recovery中间件会 recover 任何panic，如果有panic的话，会自动写入500相应码
        如果不想使用默认提供的中间件，可以使用 gin.New() 新建一个没有任何默认中间件的路由。
    Gin中间件中使用 goroutine 注意：
        在Gin的 中间件 或者 HandlerFunc 中使用 goroutine 时，不能使用原始的上下文（ctx *gin.Context），必须使用其只读副本：ctx.Copy()。


拾.Gin实践(GinToDoList)
    
    
    
    





零、壹、贰、叁、肆、伍、陆、柒、捌、玖、拾;
