package main

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/goconfig"
	"github.com/gpmgo/gopm/modules/log"
)

func main() {
	// 可以同时加载多个文件，后面的文件会覆盖前面的相同键的值
	cfg,err := goconfig.LoadConfigFile("conf.ini","conf_save.ini")
	if err !=nil{
		log.Debug("无法加载配置文件")
	}

	email,err:=cfg.GetValue("","author")
	fmt.Println(email)

	cfg.AppendFiles("conf.ini")
	email,err=cfg.GetValue("","author")
	fmt.Println(email)

}
