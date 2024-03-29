# 第006节：middleware的编写与使用

[Golang中国](https://www.qfgolang.com/?author=3) * 2019-11-11 * 阅读 2*7

### 中间件

在web应用服务中，完整的一个业务处理在技术上包含客户端操作、服务器端处理、返回处理结果给客户端三个步骤。

在实际的业务开发和处理中，会有更负责的业务和需求场景。一个完整的系统可能要包含鉴权认证、权限管理、安全检查、日志记录等多维度的系统支持。

鉴权认证、权限管理、安全检查、日志记录等这些保障和支持系统业务属于全系统的业务，和具体的系统业务没有关联，对于系统中的所有业务都适用。

由此，在业务开发过程中，为了更好的梳理系统架构，可以将上述描述所涉及的一些通用业务单独抽离并进行开发，然后以插件化的形式进行对接。这种方式既保证了系统功能的完整，同时又有效的将具体业务和系统功能进行解耦，并且，还可以达到灵活配置的目的。

这种通用业务独立开发并灵活配置使用的组件，一般称之为"中间件"，因为其位于服务器和实际业务处理程序之间。其含义就是相当于在请求和具体的业务逻辑处理之间增加某些操作，这种以额外添加的方式不会影响编码效率，也不会侵入到框架中。中间件的位置和角色示意图如下图所示：

![img](https://www.qfgolang.com/wp-content/uploads/2019/11/WX20191012-120438@2x-1.png)

### Gin的中间件

在gin中，中间件称之为middleware,中间件的类型定义如下所示：

```go
// HandlerFunc defines the handler used by gin middleware as return value.
type HandlerFunc func(*Context)
```

HandlerFunc是一个函数类型,接收一个Context参数。用于编写程序处理函数并返回HandleFunc类型，作为中间件的定义。

### 中间件Use用法

在之前学习的课程中，均使用gin.Default创建了gin引擎engins变量，其中，就使用了中间件。如下图所示：

```go
func Default() *Engine {
    debugPrintWARNINGDefault()
    engine := New()
    engine.Use(Logger(), Recovery())
    return engine
}
//Log中间件
func Logger() HandlerFunc {
    return LoggerWithConfig(LoggerConfig{})
}
//Recovery中间件
func Recovery() HandlerFunc {
    return RecoveryWithWriter(DefaultErrorWriter)
}
```

在Default函数中，engine调用Use方法设置了Logger中间件和Recovery中间件。Use函数接收一个可变参数，类型为HandlerFunc，恰为中间件的类型。Use方法定义如下：

```go
func (engine *Engine) Use(middleware ...HandlerFunc) IRoutes {
    engine.RouterGroup.Use(middleware...)
    engine.rebuild404Handlers()
    engine.rebuild405Handlers()
    return engine
}
```

### 自定义中间件

根据上文的介绍，可以自己定义实现一个特殊需求的中间件，中间件的类型是函数，有两条标准：

- func函数
- 返回值类型为HandlerFunc

比如，我们自定义一个自己的中间件。在前面所学的内容中，我们在处理请求时，为了方便代码调试，通常都将请求的一些信息打印出来。有了中间件以后，为了避免代码多次重复编写，使用统一的中间件来完成。定义一个名为RequestInfos的中间件，在该中间件中打印请求的path和method。具体代码实现如下所示：

```go
func RequestInfos() gin.HandlerFunc {
    return func(context *gin.Context) {
        path := context.FullPath()
        method := context.Request.Method
        fmt.Println("请求Path：", path)
        fmt.Println("请求Method：", method)
    }
}

func main() {

    engine := gin.Default()
    engine.Use(RequestInfos())

    engine.GET("/query", func(context *gin.Context) {
        context.JSON(200, map[string]interface{}{
            "code": 1,
            "msg":  context.FullPath(),
        })
    })
    engine.Run(":9000")
}

```

通过两条fmt.Println打印出了请求的信息，并通过engine.Use使用RequestInfos中间件。

运行程序，能够得到正确的返回JSON格式的数据：

```go
{
    "code": 1,
    "msg": "/query"
}
```

### context.Next函数

通过一个例子和使用场景来说明Next函数的作用。

在上文自定义的中间件RequestInfos中，打印了请求了请求的path和method，接着去执行了正常的业务处理函数。如果我们想输出业务处理结果的信息，该如何实现呢。答案是使用context.Next函数。

context.Next函数可以将中间件代码的执行顺序一分为二，Next函数调用之前的代码在请求处理之前之前，当程序执行到context.Next时，会中断向下执行，转而先去执行具体的业务逻辑，执行完业务逻辑处理函数之后，程序会再次回到context.Next处，继续执行中间件后续的代码。具体用法如下：

```go
func main() {
    engine := gin.Default()
    engine.Use(RequestInfos())
    engine.GET("/query", func(context *gin.Context) {
        fmt.Println(" 中间件的使用方法  ")
        context.JSON(404, map[string]interface{}{
            "code": 1,
            "msg":  context.FullPath(),
        })
    })
    engine.Run(":9000")
}

func RequestInfos() gin.HandlerFunc {
    return func(context *gin.Context) {
        path := context.FullPath()
        method := context.Request.Method
        fmt.Println("请求Path：", path)
        fmt.Println("请求Method：", method)
        context.Next()
        fmt.Println(context.Writer.Status())
    }
}
```

执行程序，输出结果如下：

```go
请求Path： /query
请求Method： GET
 中间件的使用方法  
404
```

通过打印的顺序可以看到，Next函数将中间件程序的代码执行分为了前后连个部分，Next之前的按照顺序执行，Next之后会在业务逻辑处理完毕后再执行。

Next函数的作用及代码执行流程示意图如下图所示：

![img](https://www.qfgolang.com/wp-content/uploads/2019/11/WX20191012-163117@2x.png)

- 1、程序先执行①和②。
- 2、执行到③时，转而去执行业务处理程序。
- 3、返回到中间件中，执行④。