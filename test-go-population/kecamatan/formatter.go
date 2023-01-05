package kecamatan

type KecamatanFormatter struct {
	ID            int    `json:"id"`
	KecamatanName string `json:"kecamatan_name"`
	KotaID        int    `json:"kota_id"`
}

func FormatKecamatan(kecamatan Kecamatan) KecamatanFormatter {
	formatter := KecamatanFormatter{
		ID:            kecamatan.ID,
		KecamatanName: kecamatan.KecamatanName,
		KotaID:        kecamatan.KotaID,
	}

	return formatter
}

func FormatKecamatans(kecamatans []Kecamatan) []KecamatanFormatter {
	kecamatansFormatter := []KecamatanFormatter{}

	for _, kecamatan := range kecamatans {
		kecamatanFormatter := FormatKecamatan(kecamatan)
		kecamatansFormatter = append(kecamatansFormatter, kecamatanFormatter)
	}

	return kecamatansFormatter
}
