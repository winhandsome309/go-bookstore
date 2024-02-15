package model

type Provinces struct {
	Province     string `json:"name"`
	ProvinceCode string `json:"code"`
}

type Districts struct {
	District     string `json:"name"`
	DistrictCode string `json:"code"`
}

type Wards struct {
	Ward     string `json:"name"`
	WardCode string `json:"code"`
}
