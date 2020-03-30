package models

type Detail struct {
	ProvinceState string      `json:"provinceState"`
	CountryRegion string      `json:"countryRegion"`
	LastUpdate    int64       `json:"lastUpdate"`
	Lat           float64     `json:"lat"`
	Long          float64     `json:"long"`
	Confirmed     int         `json:"confirmed"`
	Recovered     int         `json:"recovered"`
	Deaths        int         `json:"deaths"`
	Active        int         `json:"active"`
	Admin2        interface{} `json:"admin2"`
	Fips          interface{} `json:"fips"`
	CombinedKey   string      `json:"combinedKey"`
	IncidentRate  interface{} `json:"incidentRate"`
	PeopleTested  interface{} `json:"peopleTested"`
	Iso2          string      `json:"iso2"`
	Iso3          string      `json:"iso3"`
}
