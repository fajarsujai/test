package kelurahan

type CreateKelurahanInput struct {
	KelurahanName string `json:"kelurahan_name" binding:"required"`
	KecamatanID   int    `json:"kecamatan_id" binding:"required"`
}

type GetKelurahanDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
