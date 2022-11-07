package api

//Return Availability Structs
type AvailabiltyTransport struct {
	Transport ResortsStruct `json:"availability-transport"`
}

type ResortsStruct struct {
	Resorts []ResortInfo `json:"resorts"`
}

type ResortInfo struct {
	ResortCode   string         `json:"resort-code"`
	RoomCategory []RoomCategory `json:"room-category"`
}

type RoomCategory struct {
	CategoryCode     string                   `json:"category-code"`
	Rate             []RateInfo               `json:"rate"`
	RoomAvailability []RoomAvailabilityStruct `json:"availability"`
}

type RoomAvailabilityStruct struct {
	ReservationDate string `json:"reservation-date"`
	AvailableRooms  int    `json:"available-rooms"`
}

//Retrun Rates Struct
type RatesTransport struct {
	RateTransport ResortsStruct `json:"rates-transport"`
}

type RateInfo struct {
	Zone       string  `json:"zone"`
	State      string  `json:"state"`
	Country    string  `json:"country"`
	Signet     bool    `json:"signet"`
	Wholesaler bool    `json:"wholesaler"`
	BeginDate  string  `json:"begin-date"`
	Length     int     `json:"length"`
	AdultRate  float64 `json:"adult-rate"`
	ChildRate  float64 `json:"child-rate"`
	Adults     int     `json:"adults"`
	Children   int     `json:"childrens"`
	Promotions string  `json:"promotions"`
	Active     bool    `json:"active"`
}

//Return Promotions Struct
type PromotionStruct struct {
	Id          int    `json:"promotion-id"`
	Name        string `json:"promotion-name"`
	Description string `json:"promotion-description"`
}

//Return Discounts Struct
type DiscountsStruct struct {
	Id          int    `json:"discount-id"`
	Name        string `json:"discount-name"`
	Description string `json:"discount-description"`
}

//Params Structs
type AvailabilityParams struct {
	ReservationDate string
	Length          int
	Category_Code   string
	ResortCodes     []string
}

type RateParams struct {
	Gateway          string
	Zone             string
	State            string
	Country          string
	Signet           bool
	Wholesaler       bool
	Adults           int
	Children         int
	Horizon          int
	RateStructure    string
	BookingSource    string
	ReservationDate  string
	CategoryCode     string
	DisableOccupancy bool
	RequiredFileds   []string
	ResortCodes      []string
	Rates            []int
}

type PromotionParams struct {
	PromotionIds []int
}

type DiscountParams struct {
	DiscountIds []int
}
