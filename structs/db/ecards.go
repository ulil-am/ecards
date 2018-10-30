package db

//Ecards - Ecards struct for DB
type ECards struct {
	CardNumber int    `bson:"card_number"`
	Name       string `bson:"name"`
	ExpiryDate string `bson:"expiry_date"`
	Company    string `bson:"company"`
}
