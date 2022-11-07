package api

import (
	"log"
	"testing"
)

func TestAvailability(t *testing.T) {
	var availParams AvailabilityParams

	availParams.ReservationDate = "2023-05-10"
	availParams.Length = 1
	availParams.Category_Code = "all"
	availParams.ResortCodes = []string{"SMB", "SRC"}

	log.Println(CallAvailability(availParams))
}

func TestRates(t *testing.T) {
	var rateParams RateParams

	rateParams.Gateway = "ALL"
	rateParams.Zone = "ALL"
	rateParams.State = "ALL"
	rateParams.Country = "ALL"
	rateParams.Signet = false
	rateParams.Wholesaler = false
	rateParams.Adults = 2
	rateParams.Children = 0
	rateParams.Horizon = 1
	rateParams.RateStructure = "USA"
	rateParams.BookingSource = "ALL"
	rateParams.ReservationDate = "2023-10-05"
	rateParams.ResortCodes = []string{"SGO"}
	rateParams.Rates = []int{7}
	rateParams.DisableOccupancy = true
	rateParams.CategoryCode = "ALL"
	rateParams.RequiredFileds = nil

	log.Println(CallRates(rateParams))
}

func TestPromotion(t *testing.T) {
	var promotionParams PromotionParams

	promotionParams.PromotionIds = []int{12187}

	log.Println(CallPromotion(promotionParams))
}

func TestDiscount(t *testing.T) {
	var discountParams DiscountParams

	discountParams.DiscountIds = []int{100, 101, 102}

	log.Println(CallDiscount(discountParams))
}
