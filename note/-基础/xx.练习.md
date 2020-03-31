# 面向对象编程应用

- 编写一个Student结构体，包含name、gender、age、id、score字段，分别为string、string、int、int、float64类型
- 结构体中声明一个say方法，返回string类型，方法返回信息中包含所有字段值
- 在main方法中，创建Student结构体实例(变量)，并访问say方法，并将调用结果打印输出

```go
package main
import "fmt"
func main(){
	stu := Student{
		name:"tom",
		gender:"male",
		age:18,
		id:1000,
		score:33.9,
	}
	fmt.Println(stu.say())
}

type Student struct{
	name string
	gender string
	age int 
	id int
	score float64
}

func (stu *Student) say() string{
	infoStr := fmt.Sprintf("stu info name=[%v] gender=[%v] age=[%v] id=[%v] score=[%v]",
		stu.name,
		stu.gender,
		stu.age,
		stu.id,
		stu.score,
	)
	return infoStr
}
```



# 家庭收支记账软件

- 模拟实现基于文本界面的《家庭记账软件》
- 该软件能够记录家庭的收入、支出，并能够打印收支明细表

```go
package utils
import (
	"fmt"
)

type FamilyAccount struct {
	//声明必须的字段.

	//声明一个字段，保存接收用户输入的选项
	key  string
	//声明一个字段，控制是否退出for
	loop bool
	//定义账户的余额 []
	balance float64
	//每次收支的金额
	money float64
	//每次收支的说明
	note string
	//定义个字段，记录是否有收支的行为
	flag bool
	//收支的详情使用字符串来记录
	//当有收支时，只需要对details 进行拼接处理即可
	details string
}

//编写要给工厂模式的构造方法，返回一个*FamilyAccount实例
func NewFamilyAccount() *FamilyAccount { 

	return &FamilyAccount{
		key : "",
		loop : true,
		balance : 10000.0,
		money : 0.0,
		note : "",
		flag : false,
		details : "收支\t账户金额\t收支金额\t说    明",
	}

} 

//将显示明细写成一个方法
func (this *FamilyAccount) showDetails() {
	fmt.Println("-----------------当前收支明细记录-----------------")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有收支明细... 来一笔吧!")
	}
} 

//将登记收入写成一个方法，和*FamilyAccount绑定
func (this *FamilyAccount) income() {
	
	fmt.Println("本次收入金额:")
	fmt.Scanln(&this.money)
	this.balance += this.money // 修改账户余额
	fmt.Println("本次收入说明:")
	fmt.Scanln(&this.note)
	//将这个收入情况，拼接到details变量
	//收入    11000           1000            有人发红包
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

//将登记支出写成一个方法，和*FamilyAccount绑定
func (this *FamilyAccount) pay() {
	fmt.Println("本次支出金额:")
	fmt.Scanln(&this.money)
	//这里需要做一个必要的判断
	if this.money > this.balance {
		fmt.Println("余额的金额不足")
		//break
	}
	this.balance -= this.money
	fmt.Println("本次支出说明:")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

//将退出系统写成一个方法,和*FamilyAccount绑定
func (this *FamilyAccount) exit() {

	fmt.Println("你确定要退出吗? y/n")
	choice := ""
	for {

		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入 y/n")
	}

	if choice == "y" {
		this.loop = false
	}
}


//给该结构体绑定相应的方法
//显示主菜单
func (this *FamilyAccount) MainMenu() {

	for {
		fmt.Println("\n-----------------家庭收支记账软件-----------------")
		fmt.Println("                  1 收支明细")
		fmt.Println("                  2 登记收入")
		fmt.Println("                  3 登记支出")
		fmt.Println("                  4 退出软件")
		fmt.Print("请选择(1-4)：")
		fmt.Scanln(&this.key)
		switch this.key {
			case "1":
				this.showDetails()
			case "2":
				this.income()
			case "3":
				this.pay()
			case "4":
				this.exit()
			default :
				fmt.Println("请输入正确的选项..")		
		}

		if !this.loop {
			break 
		}

	}
}
```

- main.go

```go
package main
import (
	"go_code/familyaccount/utils"
	"fmt"
)
func main() {

	fmt.Println("这个是面向对象的方式完成~~")
	utils.NewFamilyAccount().MainMenu() //思路非常清晰
}
```

