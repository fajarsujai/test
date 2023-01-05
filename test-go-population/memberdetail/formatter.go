package memberdetail

type MemberDetailFormatter struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Nik         string `json:"nik"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	ProvinsiID  int    `json:"provinsi_id"`
	KotaID      int    `json:"kota_id"`
	KecamatanID int    `json:"kecamatan_id"`
	KelurahanID int    `json:"kelurahan_id"`
}

func FormatMemberDetail(memberdetail MemberDetail) MemberDetailFormatter {
	formatter := MemberDetailFormatter{
		ID:          memberdetail.ID,
		UserID:      memberdetail.UserID,
		Nik:         memberdetail.Nik,
		PhoneNumber: memberdetail.PhoneNumber,
		Address:     memberdetail.Address,
		ProvinsiID:  memberdetail.ProvinsiID,
		KotaID:      memberdetail.KotaID,
		KecamatanID: memberdetail.KecamatanID,
		KelurahanID: memberdetail.KelurahanID,
	}

	return formatter
}

func FormatMemberDetails(memberdetails []MemberDetail) []MemberDetailFormatter {
	memberdetailsFormatter := []MemberDetailFormatter{}

	for _, memberdetail := range memberdetails {
		memberdetailFormatter := FormatMemberDetail(memberdetail)
		memberdetailsFormatter = append(memberdetailsFormatter, memberdetailFormatter)
	}

	return memberdetailsFormatter
}
