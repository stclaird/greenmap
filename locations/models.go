package locations

type LocationModel struct {
	Id                    string
	LocationType          string
	LocationName          string
	AddressFirstLine      string
	AddressTown           string
	AddressPostCode       string
	AddressState          string
	AddressCountry        string
	AddressCountryCodeISO string
	LocationPoint         LocationPoint
}

type LocationPoint struct {
	coordinates []uint64
}
