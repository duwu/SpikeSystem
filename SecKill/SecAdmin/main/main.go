package main

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/duwu/SpikeSystem/SecKill/SecAdmin/router"
)

func main() {
	err := initAll()
	if err != nil {
		panic(fmt.Sprintf("init database failed, err:%v", err))
		return
	}
	beego.Run()
}
