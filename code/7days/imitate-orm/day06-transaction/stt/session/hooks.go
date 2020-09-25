package session

import (
	"reflect"
	"stt_orm/log"
)

// 实现钩子函数

const (
	BeforeQuery  = "BeforeQuery"
	AfterQuery   = "AfterQuery"
	BeforeUpdate = "BeforeUpdate"
	AfterUpdate  = "AfterUpdate"
	BeforeDelete = "BeforeDelete"
	AfterDelete  = "AfterDelete"
	BeforeInsert = "BeforeInsert"
	AfterInsert  = "AfterInsert"
)

// 调用已经注册的钩子函数
func (s *Session) CallMethod(methodName string, value interface{}) {
	// 通过方法名称获得到方法
	method := reflect.ValueOf(s.RefTable().Model).MethodByName(methodName)
	// value的优先级比 s的RefTable的Model高
	if value != nil {
		method = reflect.ValueOf(value).MethodByName(methodName)
	}
	// 将session作为参数提供给钩子函数
	param := []reflect.Value{reflect.ValueOf(s)}
	if method.IsValid() {
		// 如果调用方法存在有效，则执行
		result := method.Call(param)
		// 判断执行返回的结果中是否有error
		if len(result) > 0 {
			err, ok := result[0].Interface().(error)
			if ok {
				log.Error(err)
			}
		}
	}
	return
}