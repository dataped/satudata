package models

type TenagaKerja struct {
	NamaUsaha          string `json:"nama_usaha" bson:"nama_usaha"`
	Provinsi           string `json:"provinsi" bson:"provinsi"`
	Kabupaten          string `json:"kabupaten" bson:"kabupaten"`
	Kecamatan          string `json:"kecamatan" bson:"kecamatan"`
	Desa               string `json:"desa" bson:"desa"`
	JenisTempatUsaha   string `json:"jenis_tempat_usaha" bson:"jenis_tempat_usaha"`
	TkLaki             int    `json:"tk_laki" bson:"tk_laki"`
	TkPerempuan        int    `json:"tk_perempuan" bson:"tk_perempuan"`
	TkDisabilLaki      int    `json:"tk_disabil_laki" bson:"tk_disabil_laki"`
	TkDisabilPerempuan int    `json:"tk_disabil_perempuan" bson:"tk_disabil_perempuan"`
	PendidikanTk       string `json:"pendidikan_tk" bson:"pendidikan_tk"`
}
