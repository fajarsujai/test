package kelurahan

type KelurahanFormatter struct {
	ID            int    `json:"id"`
	KelurahanName string `json:"kelurahan_name"`
	KecamatanID   int    `json:"kecamatan_id"`
}

func FormatKelurahan(kelurahan Kelurahan) KelurahanFormatter {
	formatter := KelurahanFormatter{
		ID:            kelurahan.ID,
		KelurahanName: kelurahan.KelurahanName,
		KecamatanID:   kelurahan.KecamatanID,
	}

	return formatter
}

func FormatKelurahans(kelurahans []Kelurahan) []KelurahanFormatter {
	kelurahansFormatter := []KelurahanFormatter{}

	for _, kelurahan := range kelurahans {
		kelurahanFormatter := FormatKelurahan(kelurahan)
		kelurahansFormatter = append(kelurahansFormatter, kelurahanFormatter)
	}

	return kelurahansFormatter
}
