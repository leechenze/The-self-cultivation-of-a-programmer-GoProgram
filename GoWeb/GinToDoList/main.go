package GinToDoList

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/http"
)

type Todo struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

// 初始化数据库
func InitDataBase() (dataBase *gorm.DB, err error) {
	dsn := "root:lcz19930316@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(
		mysql.New(
			mysql.Config{
				DSN:               dsn,
				DefaultStringSize: 256, // string字段类型的默认长度
			},
		),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info), // 打印Sql日志
		},
	)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}
	return db, nil
}

func Main() {
	// 连接数据库
	db, err := InitDataBase()
	if err != nil {
		panic(err)
	} else {
		// 创建数据表(模型绑定)
		db.AutoMigrate(&Todo{})
	}

	// 创建Gin路由
	engine := gin.Default()
	// <link href=/static/css/app.8eeeaf31.css rel=preload as=style>
	// 就是指定如上 href的地址的 /static 目录对应目录是 ./GinToDoList/static
	engine.Static("/static", "./GinToDoList/static")
	// 加载模版文件
	engine.LoadHTMLGlob("./GinToDoList/templates/*")
	// 建立 / 路由请求
	engine.GET("/", func(ctx *gin.Context) {
		// index.html 就是加载的模版文件的下的 index.html
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	// 请求接口分组
	v1Group := engine.Group("v1")
	{
		/** 待办事项 */
		// 添加
		v1Group.POST("/todo", func(ctx *gin.Context) {
			var todo Todo
			// 绑定参数（将前段传过来的参数绑定到 todo对象中）
			ctx.ShouldBind(&todo)
			// 转换前段传过来的参数数据进行打印
			marshal, _ := json.Marshal(todo)
			fmt.Printf("POST return result:%v\n", string(marshal))
			// 给数据库插入该数据
			tx := db.Create(&todo)
			// 根据影响的行数进行判断
			if tx.RowsAffected == 0 {
				ctx.JSON(http.StatusOK, gin.H{
					"code":    500,
					"message": "操作失败",
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"code":    200,
					"message": "添加成功",
				})
			}
		})
		// 查询所有
		v1Group.GET("/todo", func(ctx *gin.Context) {
			var todoList []Todo
			// 查询所有的待办事项
			tx := db.Find(&todoList)
			// 转换前段传过来的参数数据进行打印
			marshal, _ := json.Marshal(todoList)
			fmt.Printf("GET return result:%v\n", string(marshal))
			// 根据影响的行数进行判断
			if tx.RowsAffected == 0 {
				ctx.JSON(http.StatusOK, gin.H{
					"code":    500,
					"message": "操作失败",
				})
			} else {
				ctx.JSON(http.StatusOK, todoList)
			}
		})
		// 查询一个
		v1Group.GET("/todo/:id", func(ctx *gin.Context) {

		})
		// 修改
		v1Group.PUT("/todo/:id", func(ctx *gin.Context) {
			var todo Todo
			// 获取路径参数 /:id
			id := ctx.Param("id")
			// 查询对应ID的数据
			err := db.Where("id = ?", id).First(&todo).Error
			// 可以根据影响的行数(RowsAffected)进行判断，也可以通过返回错误进行判断
			if err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"code":      500,
					"message":   "操作失败",
					"errorInfo": err.Error(),
				})
				return
			} else {
				// 绑定参数（将前段传过来的参数绑定到 todo对象中）
				ctx.ShouldBind(&todo)
				// 更新数据,返回错误判断的简写方式
				if err := db.Save(&todo).Error; err != nil {
					ctx.JSON(http.StatusOK, gin.H{
						"code":      500,
						"message":   "操作失败",
						"errorInfo": err.Error(),
					})
				} else {
					ctx.JSON(http.StatusOK, todo)
				}
			}
		})
		// 删除
		v1Group.DELETE("/todo/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			tx := db.Where("id = ?", id).Delete(Todo{})
			// 根据影响的行数进行判断
			if tx.RowsAffected == 0 {
				ctx.JSON(http.StatusOK, gin.H{
					"code":    500,
					"message": "操作失败",
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"code":    200,
					"message": "删除成功",
				})
			}
		})
	}

	engine.Run(":9090")

}
