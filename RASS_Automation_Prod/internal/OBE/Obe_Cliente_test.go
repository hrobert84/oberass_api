package obe

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func TestRunpr(t *testing.T) {

	c := New()

	var temp Quote_Availability

	temp.Tmpsession_id = "QQAAZZWWSSXX"
	temp.Tmpresort = "SMB"
	temp.Tmpbegin_date = "2023-05-11"
	temp.Tmp_days_stay = 7
	temp.Tmpminrooms = 1
	temp.Tmpadults = 2
	temp.Tmpchildren = 0
	temp.P_infant = 0
	temp.Tmpgateway = "MIA"
	temp.Tmppromo_id = 0
	temp.Pi_resort_type = "S"
	temp.Pi_region = "JAMAICA"
	temp.Pi_state = "FLORIDA"
	temp.Pi_zone = 3
	temp.Pi_country = "USA"
	temp.Pi_ssg_no = 0
	temp.Pi_res_insert_source = "ALL"
	temp.Pi_rate_structure = "USA"
	temp.Pi_priv_structure = "x"
	temp.P_mesage = "@pmsg"
	temp.Pi_air_exist = "N"
	temp.Pi_name_type = "D"
	temp.Pi_name_code = "X"

	var ctx context.Context

	a := c.Quote_availability(ctx, temp)

	fmt.Println("Resort")

	for _, v := range a {

		fmt.Println(v)

	}
}

func TestStage(t *testing.T) {
	c := New()

	temp := Stage_Params{Num: 7, Date: "2023-01-01", Room_Cat: "all", Resort_Code: "SMB"}

	a := c.Test_availability(context.TODO(), temp)

	for _, v := range a {
		log.Println(v)
	}
}

func TestRates(t *testing.T) {
	c := New()

	temp := Rate_Params{
		Tmpresort:            "SGO",
		Tmpbegin_date:        "2023-12-25",
		Tmp_days_stay:        7,
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

	for _, v := range a {

		fmt.Println(v)
	}
}
