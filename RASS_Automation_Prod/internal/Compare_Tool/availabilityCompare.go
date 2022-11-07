package compare_tool

import (
	Api_Client "Rass/internal/API"
	Obe_Cliente "Rass/internal/OBE"
	"context"
	"fmt"
)

type obedat struct {
	Date         string
	Availability int
}

// rass_availability should match with obe availability 3 dias 2 adultos  1 week
func AvailabilityTest1() {

	var faildatacount = 0
	var passdatacount = 0

	//   Rass data input
	var availParams Api_Client.AvailabilityParams

	availParams.ReservationDate = "2023-01-07"
	availParams.Length = 5
	availParams.Category_Code = "all"
	availParams.ResortCodes = []string{"SMB"}

	var result = Api_Client.CallAvailability(availParams)

	//OBB data input
	c := Obe_Cliente.New()

	temp := Obe_Cliente.Stage_Params{Num: 5, Date: "2023-01-07", Room_Cat: "all", Resort_Code: "SMB"}

	a := c.Test_availability(context.TODO(), temp)

	// Geting  Rass data for

	for _, resort := range result.Transport.Resorts {

		fmt.Println(resort.ResortCode)

		for i := 0; i < len(resort.RoomCategory); i++ {

			fmt.Println(resort.RoomCategory[i].CategoryCode)

			var temparr []obedat

			for _, v := range a {

				if v.Room_Cat == resort.RoomCategory[i].CategoryCode {

					tempdat := obedat{Date: v.Date, Availability: v.Available_Rooms}
					temparr = append(temparr, tempdat)

				}

			}

			for e := 0; e < len(resort.RoomCategory[i].RoomAvailability); e++ {

				if resort.RoomCategory[i].RoomAvailability[e].ReservationDate == temparr[e].Date && resort.RoomCategory[i].RoomAvailability[e].AvailableRooms == temparr[e].Availability {

					fmt.Println(resort.RoomCategory[i].RoomAvailability[e].ReservationDate, " Rass Availability => ", resort.RoomCategory[i].RoomAvailability[e].AvailableRooms, " ", temparr[e].Availability, " <= Obe Availability ", temparr[e].Date)
					passdatacount++

				} else {

					fmt.Println(resort.RoomCategory[i].RoomAvailability[e].ReservationDate, " Rass Availability => ", resort.RoomCategory[i].RoomAvailability[e].AvailableRooms, " ", temparr[e].Availability, " <= Obe Availability ", temparr[e].Date)
					faildatacount++

				}

			}

		}

	}

	fmt.Println("\033[32m", "PASS Data = ", passdatacount)
	fmt.Println("\033[31m", "MisMatch Data = ", faildatacount, "\033[0m")

}
