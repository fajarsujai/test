package kota

type KotaFormatter struct {
	ID         int    `json:"id"`
	KotaName   string `json:"kota_name"`
	ProvinsiID int    `json:"provinsi_id"`
}

func FormatKota(kota Kota) KotaFormatter {
	formatter := KotaFormatter{
		ID:         kota.ID,
		KotaName:   kota.KotaName,
		ProvinsiID: kota.ProvinsiID,
	}

	return formatter
}

func FormatKotas(kotas []Kota) []KotaFormatter {
	kotasFormatter := []KotaFormatter{}

	for _, kota := range kotas {
		kotaFormatter := FormatKota(kota)
		kotasFormatter = append(kotasFormatter, kotaFormatter)
	}

	return kotasFormatter
}
