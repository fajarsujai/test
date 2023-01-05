package memberdetail

type CreateMemberDetailInput struct {
	UserID      int    `json:"user_id" binding:"required"`
	Nik         string `json:"nik" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Address     string `json:"address" binding:"required"`
	ProvinsiID  int    `json:"provinsi_id" binding:"required"`
	KotaID      int    `json:"kota_id" binding:"required"`
	KecamatanID int    `json:"kecamatan_id" binding:"required"`
	KelurahanID int    `json:"kelurahan_id" binding:"required"`
}

type GetMemberDetailDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
