package config

import (
	"os"

	"github.com/aiteung/atdb"
)

var MongoString string = os.Getenv("MONGOSTRING")

var DBUmkmmongoinfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "umkm",
}

var Umkmmongoconn = atdb.MongoConnect(DBUmkmmongoinfo)
