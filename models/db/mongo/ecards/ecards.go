package ecards

import (
	db "ecards/models/db/mongo"

	"github.com/astaxie/beego"
	"github.com/globalsign/mgo"
)

// Ecards ...
type Ecards struct{}

func init() {
	var d Ecards
	d.Index()
}

// GetColl ...
func (d *Ecards) GetColl() (sess *mgo.Session, coll *mgo.Collection, err error) {
	sess, err = db.Connect()
	if err != nil {
		beego.Warning("Error get collection Ecards", err)
		return
	}

	coll = sess.DB("ecards").C("ecards")

	return
}

// Index ...
func (d *Ecards) Index() (err error) {
	sess, coll, err := d.GetColl()
	defer sess.Close()
	if err != nil {
		return
	}

	index := mgo.Index{
		Key:        []string{"card_number"},
		Unique:     true,
		DropDups:   false,
		Background: false,
		Sparse:     false,
	}

	err = coll.EnsureIndex(index)
	return
}

// InsertECards ...
func (d *Ecards) InsertECards(v interface{}) (err error) {
	beego.Debug("Models db ===> ", v)
	sess, coll, err := d.GetColl()
	defer sess.Close()
	beego.Debug(err)
	if err != nil {
		return
	}

	err = coll.Insert(v)
	beego.Debug(err)
	return
}
