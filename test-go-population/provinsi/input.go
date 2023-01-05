package provinsi

type CreateProvinsiInput struct {
	ProvinsiName string `json:"provinsi_name" binding:"required"`
	ProvinsiCode string `json:"provinsi_code" binding:"required"`
}

type GetProvinsiDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
