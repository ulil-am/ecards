package helper

import (
	"github.com/astaxie/beego"
)

//CheckErr ...
func CheckErr(val string, err error) {
	if err != nil {
		beego.Warning(val, err)
	}
}
