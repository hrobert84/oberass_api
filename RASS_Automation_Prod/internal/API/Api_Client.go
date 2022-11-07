package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	Api_host = "https://raswebservice.sandals.com"
)

// Checks Room Availability for the resorts sent as argument on specified date, for specifiend length if stay with specified amount of adults and children with struct AvailabilityParams
func CallAvailability(reqBody AvailabilityParams) AvailabiltyTransport {

	var availability AvailabiltyTransport

	resortCodes := `"resort-codes":[`

	for _, resort := range reqBody.ResortCodes {
		resortCodes += `"` + resort + `",`
	}

	resortCodes = strings.TrimSuffix(resortCodes, ",")
	resortCodes += "]"

	var jsonStr = []byte(`
		{
			"reservation-date": "` + reqBody.ReservationDate + `",
    		"length": ` + fmt.Sprint(reqBody.Length) + `,
			"category-code": "` + fmt.Sprint(reqBody.Category_Code) + `",
			` + resortCodes + `
    		
		}`)

	req, _ := http.NewRequest("POST", Api_host+"/RASSearch/rest/services/v2/search/availability/", bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err == nil && (resp.StatusCode >= 200 && resp.StatusCode < 205) {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &availability)
	} else {
		if err != nil {
			log.Println(err.Error(), false)
		} else {
			log.Println("Response Status: " + strconv.Itoa(resp.StatusCode))
		}
	}

	return availability
}

// Checks Rates for the specified parameters with struct RateParams
func CallRates(reqBody RateParams) RatesTransport {

	var rt RatesTransport

	requiredFields := `"required-fields":[`

	for _, field := range reqBody.RequiredFileds {
		requiredFields += `"` + field + `",`
	}

	requiredFields = strings.TrimSuffix(requiredFields, ",")
	requiredFields += "]"

	resortCodes := `"resort-codes":[`

	for _, resort := range reqBody.ResortCodes {
		resortCodes += `"` + resort + `",`
	}

	resortCodes = strings.TrimSuffix(resortCodes, ",")
	resortCodes += "]"

	rates := `"rates": [`

	for _, rate := range reqBody.Rates {
		rates += fmt.Sprint(rate)
	}

	rates = strings.TrimSuffix(rates, ",")
	rates += "]"

	var jsonStr = []byte(`
	{
		"gateway":"` + reqBody.Gateway + `",
		"zone":"` + reqBody.Zone + `",
		"state":"` + reqBody.State + `",
		"country":"` + reqBody.Country + `",
		"signet":` + fmt.Sprint(reqBody.Signet) + `,
		"wholesaler":` + fmt.Sprint(reqBody.Wholesaler) + `,
		"adults":` + fmt.Sprint(reqBody.Adults) + `,
		"childrens":` + fmt.Sprint(reqBody.Children) + `,
		"horizon":` + fmt.Sprint(reqBody.Horizon) + `,
		"rate-structure":"` + reqBody.RateStructure + `",
		"booking-source":"` + reqBody.BookingSource + `",
		"reservation-date":"` + reqBody.ReservationDate + `",
		"category-code":"` + reqBody.CategoryCode + `",
		"disable-occupancy":"` + fmt.Sprint(reqBody.DisableOccupancy) + `",
		` + requiredFields + `,
		` + resortCodes + `,
		` + rates + `
	}`)

	req, _ := http.NewRequest("POST", Api_host+"/RASSearch/rest/services/v2/search/rates/", bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err == nil && (resp.StatusCode >= 200 && resp.StatusCode < 205) {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &rt)
	} else {
		if err != nil {
			log.Println(err.Error(), false)
		} else {
			log.Println("Response Status Calling Rates: " + strconv.Itoa(resp.StatusCode))
		}
	}

	return rt
}

// Check info of the promotions with specified Id's as parameters with struct PromotionParams
func CallPromotion(reqBody PromotionParams) []PromotionStruct {

	var promotions []PromotionStruct

	promotionIds := `"promotion-ids": [`

	for _, id := range reqBody.PromotionIds {
		promotionIds += fmt.Sprint(id) + ","
	}

	promotionIds = strings.TrimSuffix(promotionIds, ",")
	promotionIds += "]"

	var jsonStr = []byte(`
	{
		` + promotionIds + `
	}`)

	req, _ := http.NewRequest("POST", Api_host+"/RASSearch/rest/services/v2/search/promotion/", bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err == nil && (resp.StatusCode >= 200 && resp.StatusCode < 205) {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &promotions)
	} else {
		if err != nil {
			log.Println(err.Error(), false)
		} else {
			log.Println("Response Status Calling Promotions: " + strconv.Itoa(resp.StatusCode))
		}
	}
	return promotions
}

// Check info of the discounts with specified Id's as parameters with struct DiscountParams
func CallDiscount(reqBody DiscountParams) []DiscountsStruct {

	var ds []DiscountsStruct

	discountIds := `"discount-ids": [`

	for _, id := range reqBody.DiscountIds {
		discountIds += fmt.Sprint(id) + ","
	}

	discountIds = strings.TrimSuffix(discountIds, ",")
	discountIds += "]"

	var jsonStr = []byte(`
	{
		` + discountIds + `
	}`)

	req, _ := http.NewRequest("POST", Api_host+"/RASSearch/rest/services/v2/search/discount/", bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err == nil && (resp.StatusCode >= 200 && resp.StatusCode < 205) {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &ds)
	} else {
		if err != nil {
			log.Println(err.Error(), false)
		} else {
			log.Println("Response Status Calling Discounts: " + strconv.Itoa(resp.StatusCode))
		}
	}

	return ds
}
