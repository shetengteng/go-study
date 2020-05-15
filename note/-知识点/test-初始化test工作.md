- 使用TestMain作为初始化test，并使用m.Run()调用其他tests可以完成一些需要初始化操作的testing，如数据库连接，文件打开，REST服务登录等

```go
func TestMain(m *testing.M){
    fmt.Println("test main")
    ...数据库连接，文件打开，rest服务登录等
    m.Run() // 执行所以其他Test方法，如果没有该语句，则其他Test不执行
}
```





