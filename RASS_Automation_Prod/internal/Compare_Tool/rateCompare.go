package compare_tool

import (
	Api_Client "Rass/internal/API"
	Obe_Cliente "Rass/internal/OBE"
	"context"
	"fmt"
	"log"
	"math"
	"sync"
	"time"
)

type Automatic struct {
	AllResorts []Data `json:"data"`
}

type Data struct {
	ResortCode string    `json:"resortCode"`
	BeginDate  string    `json:"beginDate"`
	Length     int       `json:"length"`
	Fails      int       `json:"fails"`
	Data       []Results `json:"results"`
}

type Results struct {
	RoomCategory    string  `json:"roomCategory"`
	ObePrice        float64 `json:"obePrice"`
	RassPrice       float64 `json:"rassPrice"`
	OBEMoreThanRASS float64 `json:"obeMoreThanRass"`
	RASSMoreThanOBE float64 `json:"rassMoreThanObe"`
}

// Run tests to compare room rates between API and DB
func Rate_Compare_Test(resortN string, beginDate string, lengtofStay int) ([]Results, int) {

	var datastruct []Results

	//var msg string
	var value = 000.0

	var faildatacount = 0
	var passdatacount = 0

	// Rass Api Data input

	var rateParams Api_Client.RateParams

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
	rateParams.ReservationDate = beginDate
	rateParams.ResortCodes = []string{resortN}
	rateParams.Rates = []int{lengtofStay}
	rateParams.DisableOccupancy = true
	rateParams.CategoryCode = "ALL"
	rateParams.RequiredFileds = nil

	var rass_Results = Api_Client.CallRates(rateParams)

	//OBE mysql Data input

	c := Obe_Cliente.New()

	temp := Obe_Cliente.Rate_Params{
		Tmpresort:            resortN,
		Tmpbegin_date:        beginDate,
		Tmp_days_stay:        lengtofStay,
		Pi_RoomCat:           "all",
		Tmpminrooms:          1,
		Tmpadults:            2,
		Tmpchildren:          0,
		P_infant:             0,
		Tmpgateway:           "MIA",
		Pi_state:             "FLORIDA",
		Pi_zone:              3,
		Pi_country:           "USA",
		Pi_ssg_no:            0,
		Pi_res_insert_source: "WEBBOOK",
		Pi_rate_structure:    "USA",
	}

	a := c.Test_rate(context.TODO(), temp)

	if len(rass_Results.RateTransport.Resorts) > 0 {

		for _, resort := range rass_Results.RateTransport.Resorts {

			for i := 0; i < len(resort.RoomCategory); i++ {

				var total = (resort.RoomCategory[i].Rate[0].AdultRate * float64(resort.RoomCategory[i].Rate[0].Length)) * float64(resort.RoomCategory[i].Rate[0].Adults)

				for _, v := range a {

					if v.Room_Cat == resort.RoomCategory[i].CategoryCode {

						if v.AfterPromoRate >= math.Floor(total*100)/100 {
							value = (v.AfterPromoRate - math.Floor(total*100)/100)

						} else if math.Floor(total*100)/100 >= v.AfterPromoRate {

							value = (math.Floor(total*100)/100 - v.AfterPromoRate)

						}

						if math.Floor(total*100)/100 == v.AfterPromoRate || math.Floor(value*100)/100 <= 0.50 {

							// msg += "********Room Category***********\n"
							// msg += fmt.Sprintln("       *" + resort.RoomCategory[i].CategoryCode + "*\n")

							// msg += fmt.Sprintln("OBE DATA Total Price = ", v.AfterPromoRate, "=> <=", math.Floor(total*100)/100, "Rass Data Total Price")

							// msg += fmt.Sprintln("STATUS:", "\033[32m", "PASS", "\033[0m", " ")

							// msg += fmt.Sprintln("*******************************")
							passdatacount++

						} else if v.AfterPromoRate > math.Floor(total*100)/100 {

							// msg += "\n"
							// msg += fmt.Sprintf("Resort: %s Begin Date: %s Length: %d\n", resort.ResortCode, temp.Tmpbegin_date, temp.Tmp_days_stay)

							// msg += "Room Category: " + resort.RoomCategory[i].CategoryCode + "\n"

							// msg += fmt.Sprintln("OBE Price = ", v.AfterPromoRate)
							// msg += fmt.Sprintln("Rass Price = ", math.Floor(total*100)/100)

							// msg += fmt.Sprintln("OBE > RASS by: ", math.Floor(value*100)/100)

							tempResults := Results{
								RoomCategory:    resort.RoomCategory[i].CategoryCode,
								ObePrice:        v.AfterPromoRate,
								RassPrice:       (math.Floor(total*100) / 100),
								OBEMoreThanRASS: (math.Floor(value*100) / 100),
								RASSMoreThanOBE: 0.0,
							}

							datastruct = append(datastruct, tempResults)

							faildatacount++

						} else {

							// msg += "\n"
							// msg += fmt.Sprintf("Resort: %s Begin Date: %s Length: %d\n", resort.ResortCode, temp.Tmpbegin_date, temp.Tmp_days_stay)

							// msg += "Room Category: " + resort.RoomCategory[i].CategoryCode + "\n"

							// msg += fmt.Sprintln("OBE Price = ", v.AfterPromoRate)

							// msg += fmt.Sprintln("Rass Price = ", math.Floor(total*100)/100)

							// msg += fmt.Sprintln("RASS > OBE by: ", math.Floor(value*100)/100)

							tempResults := Results{
								RoomCategory:    resort.RoomCategory[i].CategoryCode,
								ObePrice:        v.AfterPromoRate,
								RassPrice:       (math.Floor(total*100) / 100),
								OBEMoreThanRASS: 0,
								RASSMoreThanOBE: (math.Floor(value*100) / 100),
							}

							datastruct = append(datastruct, tempResults)

							faildatacount++
						}

					}

				}

			}
		}
	}

	return datastruct, faildatacount

}

func rangeDate(start, end time.Time) func() time.Time {

	y, m, d := start.Date()
	start = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	y, m, d = end.Date()
	end = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)

	return func() time.Time {
		if start.After(end) {
			return time.Time{}
		}
		date := start
		start = start.AddDate(0, 0, 1)
		return date
	}
}

func testRangeDate(start, end time.Time, resortCode string, wg *sync.WaitGroup, c1 chan []Data) {
	defer wg.Done()

	var temp = 7

	var returnData []Data

	for rd := rangeDate(start, end); ; {
		date := rd()

		if date.IsZero() {
			break
		}
		// fmt.Println(date.Format("2006-01-02"))

		if temp == 7 {

			var temp_date = date.Format("2006-01-02")
			msg, fails := Rate_Compare_Test(resortCode, temp_date, 7)

			log.Println("Finished ", resortCode, " for dates ", temp_date, ", fails: ", fails)

			if fails > 0 {
				var tmpData Data
				tmpData.Data = msg
				tmpData.ResortCode = resortCode
				tmpData.BeginDate = temp_date
				tmpData.Length = temp
				tmpData.Fails = fails

				returnData = append(returnData, tmpData)
			}

			temp = 0
		}

		temp++

	}
	c1 <- returnData
}

func Starter(year int) Automatic {

	var returnStruct Automatic

	c1 := make(chan []Data)
	c2 := make(chan []Data)
	c3 := make(chan []Data)
	c4 := make(chan []Data)
	c5 := make(chan []Data)
	c6 := make(chan []Data)
	c7 := make(chan []Data)
	c8 := make(chan []Data)
	c9 := make(chan []Data)
	c10 := make(chan []Data)
	c11 := make(chan []Data)
	c12 := make(chan []Data)
	c13 := make(chan []Data)
	c14 := make(chan []Data)
	c15 := make(chan []Data)
	c16 := make(chan []Data)
	// resorts list

	arr := [16]string{"SMB", "SRC", "SNG", "SWH", "SGO", "BRP", "SRB", "SEB", "SLU", "SGL", "SHC", "SAT", "SLS", "SBD", "SBR", "SCR"}

	start := time.Date(year+1, time.January, 1, 0, 0, 0, 0, time.FixedZone("ET", -4*3600))
	end := start.AddDate(0, 0, 365)
	fmt.Println(start.Format("2006-01-02"), "-", end.Format("2006-01-02"))

	var wg sync.WaitGroup

	wg.Add(16)

	go testRangeDate(start, end, arr[0], &wg, c1)
	go testRangeDate(start, end, arr[1], &wg, c2)
	go testRangeDate(start, end, arr[2], &wg, c3)
	go testRangeDate(start, end, arr[3], &wg, c4)
	go testRangeDate(start, end, arr[4], &wg, c5)
	go testRangeDate(start, end, arr[5], &wg, c6)
	go testRangeDate(start, end, arr[6], &wg, c7)
	go testRangeDate(start, end, arr[7], &wg, c8)
	go testRangeDate(start, end, arr[8], &wg, c9)
	go testRangeDate(start, end, arr[9], &wg, c10)
	go testRangeDate(start, end, arr[10], &wg, c11)
	go testRangeDate(start, end, arr[11], &wg, c12)
	go testRangeDate(start, end, arr[12], &wg, c13)
	go testRangeDate(start, end, arr[13], &wg, c14)
	go testRangeDate(start, end, arr[14], &wg, c15)
	go testRangeDate(start, end, arr[15], &wg, c16)

	//wg.Wait()

	for i := 0; i < 16; i++ {
		// Await both of these values
		// simultaneously, printing each one as it arrives.
		select {
		case msg1 := <-c1:
			returnStruct.AllResorts = append(returnStruct.AllResorts, msg1...)
		case msg2 := <-c2:
			returnStruct.AllResorts = append(returnStruct.AllResorts, msg2...)
		case msg3 := <-c3:
			returnStruct.AllResorts = append(returnStruct.AllResorts, msg3...)
		case msg4 := <-c4:
			returnStruct.AllResorts = append(returnStruct.AllResorts, msg4...)
		case msg5 := <-c5:
			returnStruct.AllResorts = append(returnStruct.AllResorts, msg5...)
		case msg6 := <-c6:
			returnStruct.AllResorts = append(returnStruct.AllResorts, msg6...)
		case msg7 := <-c7:
			returnStruct.AllResorts = append(returnStruct.AllResorts, msg7...)
		case msg8 := <-c8:
			returnStruct.AllResorts = append(returnStruct.AllResorts, msg8...)
		case msg9 := <-c9:
			returnStruct.AllResorts = append(returnStruct.AllResorts, msg9...)
		case msg10 := <-c10:
			returnStruct.AllResorts = append(returnStruct.AllResorts, msg10...)
		case msg11 := <-c11:
			returnStruct.AllResorts = append(returnStruct.AllResorts, msg11...)
		case msg12 := <-c12:
			returnStruct.AllResorts = append(returnStruct.AllResorts, msg12...)
		case msg13 := <-c13:
			returnStruct.AllResorts = append(returnStruct.AllResorts, msg13...)
		case msg14 := <-c14:
			returnStruct.AllResorts = append(returnStruct.AllResorts, msg14...)
		case msg15 := <-c15:
			returnStruct.AllResorts = append(returnStruct.AllResorts, msg15...)
		case msg16 := <-c16:
			returnStruct.AllResorts = append(returnStruct.AllResorts, msg16...)
		}
	}

	wg.Wait()

	log.Println("Done!")

	return returnStruct
}
