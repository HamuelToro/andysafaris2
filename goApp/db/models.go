package db

type QandA struct {
	Question    string
	Description string
}

type PriceStructure struct {
	One   PricingSafari
	Two   PricingSafari
	Three PricingSafari
	Four  PricingSafari
	Extra PricingSafari
}

type PricingSafari struct {
	Economy string
	Comfort string
	Luxury  string
}

type AccomodationSafari struct {
	Economy []string
	Comfort []string
	Luxury  []string
}

type SafariBreakDown struct {
	Day      int16
	Title    string
	Activity string
	Image    string
}

type Safari struct {
	ID           int32
	Title        string
	Guide        string
	Days         int16
	Description  string
	Image        string
	LowestPrice  string
	Prices       map[string]PriceStructure
	Accomodation AccomodationSafari
	TourIncludes []string
	Breakdown    []SafariBreakDown
}

type Vehicle struct {
	ID                     int16
	Name                   string
	Passengers             int16
	Suitcases              int16
	Type                   string
	PricePerKM             string
	PricePerMonth          string // added
	PricePerMothWithDriver string
	PriceMonthsContract    string
	Images                 []string
	Make                   VehicleMake
	ComesWith              []string // added
}

type VehicleMake struct {
	Engine        string
	InteriorColor string
	Power         string
	FuelType      string
	Length        string
	ExteriorColor string
	Transmission  string
	Extras        string
}