package vo

import "encoding/json"

type Address struct {
	addressLine1 string
	addressLine2 string
	province     string
	city         string
	district     string
	subDistrict  string
	postCode     string
}

func NewAddress(addressLine1 string, addressLine2 string, province string, city string, district string, subDistrict string, postCode string) *Address {
	return &Address{addressLine1: addressLine1, addressLine2: addressLine2, province: province, city: city, district: district, subDistrict: subDistrict, postCode: postCode}
}

func (a *Address) AddressLine1() string {
	return a.addressLine1
}

func (a *Address) AddressLine2() string {
	return a.addressLine2
}

func (a *Address) Province() string {
	return a.province
}

func (a *Address) City() string {
	return a.city
}

func (a *Address) District() string {
	return a.district
}

func (a *Address) SubDistrict() string {
	return a.subDistrict
}

func (a *Address) PostCode() string {
	return a.postCode
}

func (a *Address) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		AddressLine1 string `json:"address_line_1"`
		AddressLine2 string `json:"address_line_2"`
		Province     string `json:"province"`
		City         string `json:"city"`
		District     string `json:"district"`
		SubDistrict  string `json:"sub_district"`
		PostCode     string `json:"post_code"`
	}{
		AddressLine1: a.addressLine1,
		AddressLine2: a.addressLine2,
		Province:     a.province,
		City:         a.city,
		District:     a.district,
		SubDistrict:  a.subDistrict,
		PostCode:     a.postCode,
	})
}

func (a *Address) UnmarshalJSON(bytes []byte) error {
	var tmp struct {
		AddressLine1 string `json:"address_line_1"`
		AddressLine2 string `json:"address_line_2"`
		Province     string `json:"province"`
		City         string `json:"city"`
		District     string `json:"district"`
		SubDistrict  string `json:"sub_district"`
		PostCode     string `json:"post_code"`
	}
	err := json.Unmarshal(bytes, &tmp)
	if err != nil {
		return err
	}
	a.addressLine1 = tmp.AddressLine1
	a.addressLine2 = tmp.AddressLine2
	a.province = tmp.Province
	a.city = tmp.City
	a.district = tmp.District
	a.subDistrict = tmp.SubDistrict
	a.postCode = tmp.PostCode
	return nil
}
