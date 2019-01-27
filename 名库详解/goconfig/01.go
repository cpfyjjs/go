package main

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/goconfig"
	"github.com/gpmgo/gopm/modules/log"
	"reflect"
)

func main() {
	cfg,err := goconfig.LoadConfigFile("conf.ini")
	if err !=nil{
		log.Debug("无法加载配置文件")
	}

	value,err := cfg.GetValue(goconfig.DEFAULT_SECTION,"author")
	if err!= nil{
		log.Debug("无法获取")
	}
	fmt.Printf("%s > %s %s\n",goconfig.DEFAULT_SECTION,"author",value)


	isInsert := cfg.SetValue(goconfig.DEFAULT_SECTION,"email","756206487@qq.com")
	if isInsert{
		fmt.Println("插入新值成功")
	}

	comment := cfg.GetSectionComments("courses")
	fmt.Println(comment)

	isInsert = cfg.SetSectionComments("courses","# 这是新的分区注释")
	if isInsert{
		fmt.Println("更改成功")
	}

	vInt,err := cfg.Int("must","int")

	if err!= nil{
		fmt.Println("获取int值失败")
	}

	fmt.Println("type :",reflect.TypeOf(vInt)," value :",vInt)

	vBool := cfg.MustBool("must","bool")
	fmt.Println(vBool)

	vBool = cfg.MustBool("must","bool404")
	fmt.Println(vBool)

	vBool = cfg.MustBool("must","bool404",true)
	fmt.Println(vBool)

	ok := cfg.DeleteKey("must","string")
	if ok{
		fmt.Println("删除成功")
	}
	err = goconfig.SaveConfigFile(cfg,"conf_save.ini")
	if err != nil{
		fmt.Println(err)
	}


}
