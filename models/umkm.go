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

type TempatUsaha struct {
	IDSiKUMKM          int    `json:"id_si_kumkm" bson:"id_si_kumkm"`
	NamaLengkap        string `json:"nama_lengkap" bson:"nama_lengkap"`
	NamaKomersil       string `json:"nama_komersil" bson:"nama_komersil"`
	AlamatIDProv       int    `json:"alamat_id_prov" bson:"alamat_id_prov"`
	AlamatIDKabKot     int    `json:"alamat_id_kab_kot" bson:"alamat_id_kab_kot"`
	AlamatIDKec        int    `json:"alamat_id_kec" bson:"alamat_id_kec"`
	AlamatIDDesaKel    int    `json:"alamat_id_desa_kel" bson:"alamat_id_desa_kel"`
	AlamatProvinsi     string `json:"alamat_provinsi" bson:"alamat_provinsi"`
	AlamatKabKot       string `json:"alamat_kab_kot" bson:"alamat_kab_kot"`
	AlamatKec          string `json:"alamat_kec" bson:"alamat_kec"`
	AlamatKel          string `json:"alamat_kel" bson:"alamat_kel"`
	IDJenisTempatUsaha int    `json:"id_jenis_tempat_usaha" bson:"id_jenis_tempat_usaha"`
	JenisTempatUsaha   string `json:"jenis_tempat_usaha" bson:"jenis_tempat_usaha"`
}

type KarakteristikPerusahaan struct {
	IDSiKUMKM          int    `json:"id_si_kumkm" bson:"id_si_kumkm"`
	NamaLengkap        string `json:"nama_lengkap" bson:"nama_lengkap"`
	NamaKomersil       string `json:"nama_komersil" bson:"nama_komersil"`
	AlamatIDProv       int    `json:"alamat_id_prov" bson:"alamat_id_prov"`
	AlamatIDKabKot     int    `json:"alamat_id_kabkot" bson:"alamat_id_kabkot"`
	AlamatIDKec        int    `json:"alamat_id_kec" bson:"alamat_id_kec"`
	AlamatIDDesaKel    int    `json:"alamat_id_desa_kel" bson:"alamat_id_desa_kel"`
	AlamatProvinsi     string `json:"alamat_provinsi" bson:"alamat_provinsi"`
	AlamatKabKot       int    `json:"alamat_kabkot" bson:"alamat_kabkot"`
	AlamatKec          string `json:"alamat_kec" bson:"alamat_kec"`
	AlamatKel          string `json:"alamat_kel" bson:"alamat_kel"`
	IDJenisTempatUsaha int    `json:"id_jenis_tempat_usaha" bson:"id_jenis_tempat_usaha"`
	JenisTempatUsaha   string `json:"jenis_tempat_usaha" bson:"jenis_tempat_usaha"`
	IDStatusBadanUsaha int    `json:"id_status_badan_usaha" bson:"id_status_badan_usaha"`
	StatusBadanUsaha   string `json:"status_badan_usaha" bson:"status_badan_usaha"`
	ProdukUtama        string `json:"produk_utama" bson:"produk_utama"`
	IDJenisKegiatan    int    `json:"id_jenis_kegiatan" bson:"id_jenis_kegiatan"`
	JenisKegiatan      string `json:"jenis_kegiatan" bson:"jenis_kegiatan"`
	KegiatanUtama      string `json:"kegiatan_utama" bson:"kegiatan_utama"`
	BulanMulaiOperasi  int    `json:"bulan_mulai_operasi" bson:"bulan_mulai_operasi"`
	TahunMulaiOperasi  int    `json:"tahun_mulai_operasi" bson:"tahun_mulai_operasi"`
}

type IdentitasPengusaha struct {
	IDSiKUMKM          interface{} `json:"id_si_kumkm" bson:"id_si_kumkm"`
	NamaLengkap        string      `json:"nama_lengkap" bson:"nama_lengkap"`
	NamaKomersil       string      `json:"nama_komersil" bson:"nama_komersil"`
	AlamatIDProv       int         `json:"alamat_id_prov" bson:"alamat_id_prov"`
	AlamatIDKabKot     int         `json:"alamat_id_kabkot" bson:"alamat_id_kabkot"`
	AlamatIDKec        int         `json:"alamat_id_kec" bson:"alamat_id_kec"`
	AlamatIDDesaKel    int         `json:"alamat_id_desa_kel" bson:"alamat_id_desa_kel"`
	AlamatProvinsi     string      `json:"alamat_provinsi" bson:"alamat_provinsi"`
	AlamatKabKot       interface{} `json:"alamat_kabkot" bson:"alamat_kabkot"`
	AlamatKec          string      `json:"alamat_kec" bson:"alamat_kec"`
	AlamatKel          string      `json:"alamat_kel" bson:"alamat_kel"`
	IDJenisTempatUsaha int         `json:"id_jenis_tempat_usaha" bson:"id_jenis_tempat_usaha"`
	JenisTempatUsaha   string      `json:"jenis_tempat_usaha" bson:"jenis_tempat_usaha"`
	JenisKelamin       string      `json:"jenis_kelamin" bson:"jenis_kelamin"`
	Umur               int         `json:"umur" bson:"umur"`
}

type Izin struct {
	Key   int    `json:"key" bson:"key"`
	Name  string `json:"name" bson:"name"`
	Value int    `json:"value" bson:"value"`
}

type IzinDanStandarisasiUsaha struct {
	Nib                   string     `json:"nib" bson:"nib"`
	Id_si_kumkm           int        `json:"id_si_kumkm" bson:"id_si_kumkm"`
	Nama_lengkap          string     `json:"nama_lengkap" bson:"nama_lengkap"`
	Nama_komersil         string     `json:"nama_komersil" bson:"nama_komersil"`
	Alamat_id_prov        int        `json:"alamat_id_prov" bson:"alamat_id_prov"`
	Alamat_id_kabkot      int        `json:"alamat_id_kabkot" bson:"alamat_id_kabkot"`
	Alamat_id_kec         int        `json:"alamat_id_kec" bson:"alamat_id_kec"`
	Alamat_id_desa_kel    int        `json:"alamat_id_desa_kel" bson:"alamat_id_desa_kel"`
	Alamat_provinsi       string     `json:"alamat_provinsi" bson:"alamat_provinsi"`
	Alamat_kabkot         string     `json:"alamat_kabkot" bson:"alamat_kabkot"`
	Alamat_kec            string     `json:"alamat_kec" bson:"alamat_kec"`
	Alamat_kel            string     `json:"alamat_kel" bson:"alamat_kel"`
	Id_jenis_tempat_usaha int        `json:"id_jenis_tempat_usaha" bson:"id_jenis_tempat_usaha"`
	Jenis_tempat_usaha    string     `json:"jenis_tempat_usaha" bson:"jenis_tempat_usaha"`
	Izin                  [][][]Izin `json:"izin" bson:"izin"`
	Izin_lain             string     `json:"izin_lain" bson:"izin_lain"`
	Standarisasi          string     `json:"standarisasi" bson:"standarisasi"`
	Standarisasi_lain     string     `json:"standarisasi_lain" bson:"standarisasi_lain"`
}

type Penghargaan struct {
	Key   int    `json:"key" bson:"key"`
	Name  string `json:"name" bson:"name"`
	Value int    `json:"value" bson:"value"`
}
type PenghargaanUsaha struct {
	Id_si_kumkm           int               `json:"id_si_kumkm" bson:"id_si_kumkm"`
	Nama_lengkap          string            `json:"nama_lengkap" bson:"nama_lengkap"`
	Nama_komersil         string            `json:"nama_komersil" bson:"nama_komersil"`
	Alamat_id_prov        int               `json:"alamat_id_prov" bson:"alamat_id_prov"`
	Alamat_id_kabkot      int               `json:"alamat_id_kabkot" bson:"alamat_id_kabkot"`
	Alamat_id_kec         int               `json:"alamat_id_kec" bson:"alamat_id_kec"`
	Alamat_id_desa_kel    int               `json:"alamat_id_desa_kel" bson:"alamat_id_desa_kel"`
	Alamat_provinsi       string            `json:"alamat_provinsi" bson:"alamat_provinsi"`
	Alamat_kabkot         string            `json:"alamat_kabkot" bson:"alamat_kabkot"`
	Alamat_kec            string            `json:"alamat_kec" bson:"alamat_kec"`
	Alamat_kel            string            `json:"alamat_kel" bson:"alamat_kel"`
	Id_jenis_tempat_usaha int               `json:"id_jenis_tempat_usaha" bson:"id_jenis_tempat_usaha"`
	Jenis_tempat_usaha    string            `json:"jenis_tempat_usaha" bson:"jenis_tempat_usaha"`
	Penghargaan           [][][]Penghargaan `json:"penghargaan" bson:"penghargaan"`
}

type BahanBaku struct {
	Key   int    `json:"key" bson:"key"`
	Name  string `json:"name" bson:"name"`
	Value int    `json:"value" bson:"value"`
}

type BahanBakuSkalaSumber struct {
	Key   int    `json:"key" bson:"key"`
	Name  string `json:"name" bson:"name"`
	Value int    `json:"value" bson:"value"`
}

type BahanBakuPenolong struct {
	Id_si_kumkm             int                        `json:"id_si_kumkm" bson:"id_si_kumkm"`
	Nama_lengkap            string                     `json:"nama_lengkap" bson:"nama_lengkap"`
	Nama_komersil           string                     `json:"nama_komersil" bson:"nama_komersil"`
	Alamat_id_prov          int                        `json:"alamat_id_prov" bson:"alamat_id_prov"`
	Alamat_id_kabkot        int                        `json:"alamat_id_kabkot" bson:"alamat_id_kabkot"`
	Alamat_id_kec           int                        `json:"alamat_id_kec" bson:"alamat_id_kec"`
	Alamat_id_desa_kel      int                        `json:"alamat_id_desa_kel" bson:"alamat_id_desa_kel"`
	Alamat_provinsi         string                     `json:"alamat_provinsi" bson:"alamat_provinsi"`
	Alamat_kabkot           string                     `json:"alamat_kabkot" bson:"alamat_kabkot"`
	Alamat_kec              string                     `json:"alamat_kec" bson:"alamat_kec"`
	Alamat_kel              string                     `json:"alamat_kel" bson:"alamat_kel"`
	Id_jenis_tempat_usaha   int                        `json:"id_jenis_tempat_usaha" bson:"id_jenis_tempat_usaha"`
	Jenis_tempat_usaha      string                     `json:"jenis_tempat_usaha" bson:"jenis_tempat_usaha"`
	Bahan_baku              [][][]BahanBaku            `json:"bahan_baku" bson:"bahan_baku"`
	Bahan_baku_skala_sumber [][][]BahanBakuSkalaSumber `json:"bahan_baku_skala_sumber" bson:"bahan_baku_skala_sumber"`
}

type Produck struct {
	Id_produk          int    `json:"id_produk" bson:"id_produk"`
	Nama_produk        string `json:"nama_produk" bson:"nama_produk"`
	Kuantitas_produksi int    `json:"kuantitas_produksi" bson:"kuantitas_produksi"`
	Id_satuan          int    `json:"id_satuan" bson:"id_satuan"`
	Satuan             string `json:"satuan" bson:"satuan"`
	Nilai_jual_total   int    `json:"nilai_jual_total" bson:"nilai_jual_total"`
}

type ProduksiSelamaSebulan struct {
	Id_si_kumkm           int       `json:"id_si_kumkm" bson:"id_si_kumkm"`
	Nama_lengkap          string    `json:"nama_lengkap" bson:"nama_lengkap"`
	Nama_komersil         string    `json:"nama_komersil" bson:"nama_komersil"`
	Alamat_id_prov        int       `json:"alamat_id_prov" bson:"alamat_id_prov"`
	Alamat_id_kabkot      int       `json:"alamat_id_kabkot" bson:"alamat_id_kabkot"`
	Alamat_id_kec         int       `json:"alamat_id_kec" bson:"alamat_id_kec"`
	Alamat_id_desa_kel    int       `json:"alamat_id_desa_kel" bson:"alamat_id_desa_kel"`
	Alamat_provinsi       string    `json:"alamat_provinsi" bson:"alamat_provinsi"`
	Alamat_kabkot         string    `json:"alamat_kabkot" bson:"alamat_kabkot"`
	Alamat_kec            string    `json:"alamat_kec" bson:"alamat_kec"`
	Alamat_kel            string    `json:"alamat_kel" bson:"alamat_kel"`
	Id_jenis_tempat_usaha int       `json:"id_jenis_tempat_usaha" bson:"id_jenis_tempat_usaha"`
	Jenis_tempat_usaha    string    `json:"jenis_tempat_usaha" bson:"jenis_tempat_usaha"`
	Produk                []Produck `json:"produk" bson:"produk"`
}

type Mitra_produksi struct {
	Id_pola_kemitraan    int    `json:"id_pola_kemitraan" bson:"id_pola_kemitraan"`
	Pola_kemitraan       string `json:"pola_kemitraan" bson:"pola_kemitraan"`
	Id_lingkup_kemitraan int    `json:"id_lingkup_kemitraan" bson:"id_lingkup_kemitraan"`
	Lingkup_kemitraan    string `json:"lingkup_kemitraan" bson:"lingkup_kemitraan"`
	Nama_mitra           string `json:"nama_mitra" bson:"nama_mitra"`
	Alamat               string `json:"alamat" bson:"alamat"`
	Telp_hp              string `json:"telp_hp" bson:"telp_hp"`
}

type MitraProduksi struct {
	Id_si_kumkm           int                  `json:"id_si_kumkm" bson:"id_si_kumkm"`
	Nama_lengkap          string               `json:"nama_lengkap" bson:"nama_lengkap"`
	Nama_komersil         string               `json:"nama_komersil" bson:"nama_komersil"`
	Alamat_id_prov        int                  `json:"alamat_id_prov" bson:"alamat_id_prov"`
	Alamat_id_kabkot      int                  `json:"alamat_id_kabkot" bson:"alamat_id_kabkot"`
	Alamat_id_kec         int                  `json:"alamat_id_kec" bson:"alamat_id_kec"`
	Alamat_id_desa_kel    int                  `json:"alamat_id_desa_kel" bson:"alamat_id_desa_kel"`
	Alamat_provinsi       string               `json:"alamat_provinsi" bson:"alamat_provinsi"`
	Alamat_kabkot         string               `json:"alamat_kabkot" bson:"alamat_kabkot"`
	Alamat_kec            string               `json:"alamat_kec" bson:"alamat_kec"`
	Alamat_kel            string               `json:"alamat_kel" bson:"alamat_kel"`
	Id_jenis_tempat_usaha int                  `json:"id_jenis_tempat_usaha" bson:"id_jenis_tempat_usaha"`
	Jenis_tempat_usaha    string               `json:"jenis_tempat_usaha" bson:"jenis_tempat_usaha"`
	Mitra_produksi        [][][]Mitra_produksi `json:"mitra_produksi" bson:"mitra_produksi"`
}
type PengeluaranKeuangan struct {
	Id_si_kumkm           int                    `json:"id_si_kumkm" bson:"id_si_kumkm"`
	Nama_lengkap          string                 `json:"nama_lengkap" bson:"nama_lengkap"`
	Nama_komersil         string                 `json:"nama_komersil" bson:"nama_komersil"`
	Alamat_id_prov        int                    `json:"alamat_id_prov" bson:"alamat_id_prov"`
	Alamat_id_kabkot      int                    `json:"alamat_id_kabkot" bson:"alamat_id_kabkot"`
	Alamat_id_kec         int                    `json:"alamat_id_kec" bson:"alamat_id_kec"`
	Alamat_id_desa_kel    int                    `json:"alamat_id_desa_kel" bson:"alamat_id_desa_kel"`
	Alamat_provinsi       string                 `json:"alamat_provinsi" bson:"alamat_provinsi"`
	Alamat_kabkot         string                 `json:"alamat_kabkot" bson:"alamat_kabkot"`
	Alamat_kec            string                 `json:"alamat_kec" bson:"alamat_kec"`
	Alamat_kel            string                 `json:"alamat_kel" bson:"alamat_kel"`
	Id_jenis_tempat_usaha int                    `json:"id_jenis_tempat_usaha" bson:"id_jenis_tempat_usaha"`
	Jenis_tempat_usaha    string                 `json:"jenis_tempat_usaha" bson:"jenis_tempat_usaha"`
	Pengeluaran_keuangan  []Pengeluaran_keuangan `json:"pengeluaran_keuangan" bson:"pengeluaran_keuangan"`
}

type Pengeluaran_keuangan struct {
	Id_item_pengeluaran     int     `json:"id_item_pengeluaran" bson:"id_item_pengeluaran"`
	Item_pengeluaran        string  `json:"item_pengeluaran" bson:"item_pengeluaran"`
	Banyak_volume           float64 `json:"banyak_volume" bson:"banyak_volume"`
	Nilai_rp_tahun_lalu     float64 `json:"nilai_rp_tahun_lalu" bson:"nilai_rp_tahun_lalu"`
	Nilai_rp_bulan_lalu     float64 `json:"nilai_rp_bulan_lalu" bson:"nilai_rp_bulan_lalu"`
	Nilai_rp_tahun_berjalan float64 `json:"nilai_rp_tahun_berjalan" bson:"nilai_rp_tahun_berjalan"`
}

type PembinaanPelatihan struct {
	Key   int    `json:"key" bson:"key"`
	Name  string `json:"name" bson:"name"`
	Value int    `json:"value" bson:"value"`
}

type PembinaanModal struct {
	Key   int    `json:"key" bson:"key"`
	Name  string `json:"name" bson:"name"`
	Value int    `json:"value" bson:"value"`
}

type PenyelenggaraanModal struct {
	Key   int    `json:"key" bson:"key"`
	Name  string `json:"name" bson:"name"`
	Value int    `json:"value" bson:"value"`
}

type DataBadanUsaha struct {
	IDSiKUMKM            interface{}                `json:"id_si_kumkm" bson:"id_si_kumkm"`
	IDDataBadanUsaha     int                        `json:"id_data_badan_usaha" bson:"id_data_badan_usaha"`
	NamaLengkap          string                     `json:"nama_lengkap" bson:"nama_lengkap"`
	NamaKomersil         string                     `json:"nama_komersil" bson:"nama_komersil"`
	AlamatIDProv         int                        `json:"alamat_id_prov" bson:"alamat_id_prov"`
	AlamatIDKabKot       int                        `json:"alamat_id_kabkot" bson:"alamat_id_kabkot"`
	AlamatIDKec          int                        `json:"alamat_id_kec" bson:"alamat_id_kec"`
	AlamatIDDesaKel      int                        `json:"alamat_id_desa_kel" bson:"alamat_id_desa_kel"`
	AlamatProvinsi       string                     `json:"alamat_provinsi" bson:"alamat_provinsi"`
	AlamatKabKot         interface{}                `json:"alamat_kabkot" bson:"alamat_kabkot"`
	AlamatKec            string                     `json:"alamat_kec" bson:"alamat_kec"`
	AlamatKel            string                     `json:"alamat_kel" bson:"alamat_kel"`
	IDJenisTempatUsaha   int                        `json:"id_jenis_tempat_usaha" bson:"id_jenis_tempat_usaha"`
	JenisTempatUsaha     string                     `json:"jenis_tempat_usaha" bson:"jenis_tempat_usaha"`
	PembinaanPelatiahan  [][][]PembinaanPelatihan   `json:"pembinaan_pelatiahan" bson:"pembinaan_pelatiahan"`
	PembinaanModal       [][][]PembinaanModal       `json:"pembinaan_modal" bson:"pembinaan_modal"`
	PenyelenggaraanModal [][][]PenyelenggaraanModal `json:"penyelenggaraan_modal" bson:"penyelenggaraan_modal"`
}
