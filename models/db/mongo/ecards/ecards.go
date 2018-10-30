package ecards

import (
	db "ecards/models/db/mongo"

	"github.com/astaxie/beego"
	"github.com/globalsign/mgo"
)

type Ecards struct{}

func init() {
	var d Ecards
	d.Index()
}

func (d *Ecards) GetColl() (sess *mgo.Session, coll *mgo.Collection, err error) {
	sess, err = db.Connect()
	if err != nil {
		beego.Warning("Error get collection Ecards", err)
		return
	}

	coll = sess.DB("ecards").C("ecards")

	return
}

func (d *Ecards) Index() (err error) {
	sess, coll, err := d.GetColl()
	return
}

func (d *Ecards) InsertECards(v interface{}) (err error) {
	sess, coll, err := d.GetColl()
	defer sess.Close()
	if err != nil {
		return
	}

	err = coll.Insert(v)

	return
}
