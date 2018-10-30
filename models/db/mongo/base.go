package mongo

import (
	"github.com/astaxie/beego"
	"github.com/globalsign/mgo"
)

//Connect ...
func Connect() (sess *mgo.Session, err error) {
	sess, err = mgo.Dial("mongodb://172.17.0.1:27017")
	if err != nil {
		beego.Warning("Error connect to MongoDB", err)
		return
	}

	mgo.SetDebug(true)

	sess.SetMode(mgo.Monotonic, true)

	return
}
