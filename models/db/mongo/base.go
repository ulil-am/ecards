package mongo

import (
	"ecards/helper"
	"ecards/helper/constant"

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

// GetColl - Get Collection
func GetColl(collName string) (sess *mgo.Session, coll *mgo.Collection, err error) {
	sess, err = Connect()
	if err != nil {
		helper.CheckErr("Failed get collection "+collName, err)
		return
	}

	coll = sess.DB(constant.GOAPP).C(collName)

	return
}
