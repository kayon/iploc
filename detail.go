package iploc

type Detail struct {
	IP    IP
	Start IP
	End   IP
	Location
}

func (detail Detail) String() string {
	return detail.Location.String()
}

func (detail Detail) Bytes() []byte {
	return detail.Location.Bytes()
}

func (detail Detail) InIP(ip IP) bool {
	return detail.Start.Compare(ip) < 1 && detail.End.Compare(ip) > -1
}

func (detail Detail) In(rawIP string) bool {
	ip, err := ParseIP(rawIP)
	if err != nil {
		return false
	}
	return detail.InIP(ip)
}

func (detail Detail) InUint(uintIP uint32) bool {
	return detail.InIP(ParseUintIP(uintIP))
}

func (detail *Detail) fill() *Detail {
	// TODO
	return detail
}

type Location struct {
	Country string
	Region  string
	raw     string
}

func (location Location) String() string {
	return location.raw
}

func (location Location) Bytes() []byte {
	return []byte(location.raw)
}

func parseLocation(country, region []byte) Location {
	location := Location{
		Country: string(country),
		Region:  string(region),
	}
	location.raw = location.Country
	if region != nil {
		location.raw += " " + location.Region
	}
	return location
}
