package main

import (
	"gcasbin/lib"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(lib.Middlewares()...)
	r.GET("/:domain/depts", func(context *gin.Context) {
		context.JSON(200, gin .H{"result": "部门列表--"+context.Param("domain")})
	})

	//r.GET("/depts/:id", func(context *gin.Context) {
	//	context.JSON(200, gin.H{"result": "部门详细"})
	//})

	r.POST("/:domain/depts", func(context *gin.Context) {
		context.JSON(200, gin.H{"result": "批量修改部门列表--"+context.Param("domain")})
	})
	r.Run(":8080")
}




//
//func main() {
//
//	sub:= "shenyi" // 想要访问资源的用户。
//	obj:= "/depts" // 将被访问的资源。
//	act:= "GET" // 用户对资源执行的操作。
//	e,_:= casbin.NewEnforcer("resources/model_t.conf","resources/p_t.csv")
//
//	ok,err:= e.Enforce(sub,"domain1", obj, act)
//	if err==nil && ok {
//		log.Println("运行通过")
//	}
//
//
//}

























//type Role struct {
//}
//
//func (r *Role) TableName() {
//	fmt.Println("劳资设置了roles表名")
//}
//
//type Model interface {
//	TableName()
//}
//type GORM struct { //假冒Gorm
//}
//
//func (g *GORM) Do(model interface{}) { //假设就是 执行数据库方法
//	//下面是一坨反射代码
//	t := reflect.TypeOf(model)
//	v := reflect.ValueOf(model)
//
//	if t.Implements(reflect.TypeOf((*Model)(nil)).Elem()) {
//		if method, ok := t.MethodByName("TableName"); ok {
//			method.Func.Call([]reflect.Value{v})
//		} else {
//			fmt.Println("fwef")
//		}
//
//	}
//
//	fmt.Println("执行数据库方法")
//}
//func main() {
//	gorm := &GORM{}
//	role := &Role{}
//	gorm.Do(role)
//}
