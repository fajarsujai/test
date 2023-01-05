package kecamatan

type CreateKecamatanInput struct {
	KecamatanName string `json:"kecamatan_name" binding:"required"`
	KotaID        int    `json:"kota_id" binding:"required"`
}

type GetKecamatanDetailInput struct {
	ID int `uri:"id", binding:"required"`
}
