package main

import (
	"net/http"
	"reflect"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
	"time"
)

/*
   对绑定解析到结构体上的参数，自定义验证功能
   比如我们要对 name 字段做校验，要不能为空，并且不等于 admin ，类似这种需求，就无法 binding 现成的方法
   需要我们自己验证方法才能实现 官网示例（https://godoc.org/gopkg.in/go-playground/validator.v8#hdr-Custom_Functions）
   这里需要下载引入下 gopkg.in/go-playground/validator.v8
*/
type Person struct {
	Age int `form:"age" binding:"required,gt=10"`
	// 2、在参数 binding 上使用自定义的校验方法函数注册时候的名称
	Name    string `form:"name" binding:"NotNullAndAdmin"`
	Address string `form:"address" binding:"required"`
}

// Booking contains binded and validated data.
type Booking struct {
	//定义一个预约的时间大于今天的时间
	CheckIn time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	//gtfield=CheckIn退出的时间大于预约的时间
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

// 自定义的校验方法
func nameNotNullAndAdmin(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if value, ok := field.Interface().(string); ok {
		// 字段不能为空，并且不等于  admin
		return value != "" && !("admin" == value)
	}
	return true
}

func bookableDate(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	//field.Interface().(time.Time)获取参数值并且转换为时间格式
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		if today.Unix() > date.Unix() {
			return false
		}
	}
	return true
}

func main() {
	r := gin.Default()

	// 3、将我们自定义的校验方法注册到 validator中
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 这里的 key 和 fn 可以不一样最终在 struct 使用的是 key
		v.RegisterValidation("NotNullAndAdmin", nameNotNullAndAdmin)
		//绑定第一个参数是验证的函数第二个参数是自定义的验证函数
		v.RegisterValidation("bookabledate", bookableDate)
	}

	/*
	   curl -X GET "http://127.0.0.1:8080/testing?name=&age=12&address=beijing"
	   curl -X GET "http://127.0.0.1:8080/testing?name=lmh&age=12&address=beijing"
	   curl -X GET "http://127.0.0.1:8080/testing?name=adz&age=12&address=beijing"
	*/
	r.GET("/person", func(c *gin.Context) {
		var person Person
		if e := c.ShouldBind(&person); e == nil {
			c.String(http.StatusOK, "%v", person)
		} else {
			c.String(http.StatusOK, "person bind err:%v", e.Error())
		}
	})
	// curl -X GET "http://localhost:8080/5lmh?check_in=2019-11-07&check_out=2019-11-20"
	// curl -X GET "http://localhost:8080/5lmh?check_in=2019-09-07&check_out=2019-11-20"
	// curl -X GET "http://localhost:8080/5lmh?check_in=2019-11-07&check_out=2019-11-01"
	r.GET("/book", func(c *gin.Context) {
		var b Booking
		if err := c.ShouldBindWith(&b, binding.Query); err == nil {
			c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
	r.Run()
}
