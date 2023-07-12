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
    ... TODO ...








零、壹、贰、叁、肆、伍、陆、柒、捌、玖、拾;
