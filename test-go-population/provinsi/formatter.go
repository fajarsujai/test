package provinsi

type ProvinsiFormatter struct {
	ID           int    `json:"id"`
	ProvinsiName string `json:"provinsi_name"`
	ProvinsiCode string `json:"provinsi_code"`
}

func FormatProvinsi(provinsi Provinsi) ProvinsiFormatter {
	formatter := ProvinsiFormatter{
		ID:           provinsi.ID,
		ProvinsiName: provinsi.ProvinsiName,
		ProvinsiCode: provinsi.ProvinsiCode,
	}

	return formatter
}

func FormatProvinsis(provinsis []Provinsi) []ProvinsiFormatter {
	provinsisFormatter := []ProvinsiFormatter{}

	for _, provinsi := range provinsis {
		provinsiFormatter := FormatProvinsi(provinsi)
		provinsisFormatter = append(provinsisFormatter, provinsiFormatter)
	}

	return provinsisFormatter
}

// type ProvinsiDetailFormatter struct {
// 	ID           int    `json:"id"`
// 	ProvinsiName string `json:"provinsi_name"`
// 	ProvinsiCode string `json:"provinsi_code"`
// }

// func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
// 	campaignDetailFormatter := CampaignDetailFormatter{}
// 	campaignDetailFormatter.ID = campaign.ID
// 	campaignDetailFormatter.Name = campaign.Name

// 	return campaignDetailFormatter
// }
