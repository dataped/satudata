package helper

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/dataped/satudata/models"
)

func generateRandomData() []models.TenagaKerja {
	rand.Seed(time.Now().UnixNano())
	var data []models.TenagaKerja

	for i := 1; i <= 100000; i++ {
		// Buat data dummy dengan nilai-nilai acak
		// Sesuaikan bagian ini dengan struktur TenagaKerja Anda
		dummy := models.TenagaKerja{
			IDSiKumkm: i,
			// Isi nilai-nilai lain dengan data acak atau tetap
			NamaLengkap: fmt.Sprintf("Usaha Indonesia %d", i),
			// ...
			TenagaKerjaPendididikan: []models.TenagaKerjaPendidikan{
				{
					IDPendidikanFormal: 1,
					PendidikanFormal:   "Tidak Tamat SD",
					TkDibayarLaki:      rand.Intn(10),
					TkPerempuan:        rand.Intn(10),
					// ...
				},
				// Tambahkan lebih banyak data pendidikan jika diperlukan
			},
			// Isi data sesuai dengan struktur TenagaKerja
			// Contoh: IDSiKumkm: i,
			// ...
		}
		data = append(data, dummy)
	}

	return data
}

func main() {
	data := generateRandomData()

	// Menyimpan data ke file JSON
	file, err := os.Create("tenaga_kerja_dummy.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println("File tenaga_kerja_dummy.json berhasil dibuat")
}
