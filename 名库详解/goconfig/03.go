package main

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/goconfig"
	"github.com/gpmgo/gopm/modules/log"
)

func main() {
	// 可以同时加载多个文件，后面的文件会覆盖前面的相同键的值
	cfg,err := goconfig.LoadConfigFile("conf.ini",)
	if err !=nil{
		log.Debug("无法加载配置文件")
	}

	// 递归查询
	mygithub,err:=cfg.GetValue("","mygithub")
	fmt.Println(mygithub)
	//https://github.com/Unknwon

	// 子分区继承父分区
	name,err:=cfg.GetValue("dir.Go名库讲解","name")
	fmt.Println(name)

	// 子分区继承父分区
	age,err:=cfg.GetValue("dir.Go名库讲解","age")
	fmt.Println(age)

	sec,err:=cfg.GetSection("auto increment")
	fmt.Println(sec)
	fmt.Println(sec["#1"])


}
