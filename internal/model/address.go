package model

type RequestAddressSearch struct {
	Query string `json:"query"`
}

type RequestGeocode struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Address struct {
	Source       string      `json:"source"`
	Result       string      `json:"result"`
	PostalCode   string      `json:"postal_code"`
	Country      string      `json:"country"`
	Region       string      `json:"region"`
	CityArea     string      `json:"city_area"`
	CityDistrict string      `json:"city_district"`
	Street       string      `json:"street"`
	House        string      `json:"house"`
	GeoLat       string      `json:"geo_lat"`
	GeoLon       string      `json:"geo_lon"`
	QCGeo        interface{} `json:"qc_geo"`
}

type ResponseAddress struct {
	Addresses []*Address `json:"addresses"`
}
