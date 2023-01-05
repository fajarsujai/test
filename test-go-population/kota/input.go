package kota

type CreateKotaInput struct {
	KotaName   string `json:"kota_name" binding:"required"`
	ProvinsiID int    `json:"provinsi_id" binding:"required"`
}

type GetKotaDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
