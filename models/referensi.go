package models

import "time"

type Provinsi struct {
	ID           int    `json:"id" bson :"id"`
	NoProvinsi   int    `json:"no_provinsi" bson :"no_provinsi"`
	NamaProvinsi string `json:"nama_provinsi" bson :"nama_provinsi"`
	CreatedAt    time.Time `json:"created_at bson:"created_at"`
	UpdatedAt    time.Time `json: updated_at bson:"updated_at"`
}
type Kabupaten struct {
	ID            int    `json:"id" bson:"id"`
	NoKabupaten   int    `json:"no_kabupaten" bson:"no_kabupaten"`
	NamaKabupaten string `json:"nama_kabupaten" bson:"nama_kabupaten"`
	ProvinsiID    int    `json:"provinsi_id" bson:"provinsi_id"`
	CreatedAt     time.Time `json:"created_at bson:"created_at"`
	UpdatedAt     time.Time `json: updated_at bson:"updated_at"`
}

type Kecamatan struct {
	ID            int    `json:"id" bson:"id"`
	NoKecamatan   int    `json:"no_kecamatan" bson:"no_kecamatan"`
	NamaKecamatan string `json:"nama_kecamatan" bson:"nama_kecamatan"`
	ProvinsiID    int    `json:"provinsi_id" bson:"provinsi_id"`
	KabupatenID   int    `json:"kabupaten_id" bson:"kabupaten_id"`
	CreatedAt     time.Time `json:"created_at bson:"created_at"`
	UpdatedAt     time.Time `json: updated_at bson:"updated_at"`
}

type Kelurahan struct {
	ID            int    `json:"id" bson:"id"`
	NoKelurahan   int    `json:"no_kelurahan" bson:"no_kelurahan"`
	NamaKelurahan string `json:"nama_kelurahan" bson:"nama_kelurahan"`
	ProvinsiID    int    `json:"provinsi_id" bson:"provinsi_id"`
	KabupatenID   int    `json:"kabupaten_id" bson:"kabupaten_id"`
	KecamatanID   int    `json:"kecamatan_id" bson:"kecamatan_id"`
	CreatedAt     time.Time `json:"created_at bson:"created_at"`
	UpdatedAt     time.Time `json: updated_at bson:"updated_at"`
}
