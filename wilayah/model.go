package wilayah

type Province struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type City struct {
	ID         int    `json:"id"`
	Type       string `json:"type"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	FullCode   string `json:"full_code"`
	ProvinceID int    `json:"provinsi_id"`
}

type District struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Code     string `json:"code"`
	FullCode string `json:"full_code"`
	CityID   int    `json:"kabupaten_id"`
}

type Village struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	FullCode   string `json:"full_code"`
	PosCode    string `json:"pos_code"`
	DistrictID int    `json:"kecamatan_id"`
}
