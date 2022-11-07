package obe

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

type Client struct {
	Config *conf
}

type conf struct {
	DB_User       string `yaml:"OBE_user"`
	DB_Pass       string `yaml:"OBE_pass"`
	DB_Host       string `yaml:"OBE_host"`
	DB_Stage_Host string `yaml:"Stagegold2_obe_host"`
	DB_Stage_User string `yaml:"Stagegold2_obe_user"`
	DB_Stage_Pass string `yaml:"Stagegold2_obe_pass"`
}

func New() *Client {
	return &Client{
		Config: Credentials(),
	}
}

// Function that grabs the credentials from the .yaml file and populates the "conf" struct with its contents. Return the *conf struct
func Credentials() *conf {
	var c *conf

	//If from main path = credentials.yaml
	// IF is from Test folder = ../../credentials.yaml
	yamlFile, err := os.ReadFile("credentials.yaml")
	if err != nil {
		log.Println("Error reading yaml file: #"+err.Error(), true)
	}
	//With Unmarshal we pass the values from our variable into our "conf" struct
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Println("Unmarshal: "+err.Error(), true)
	}

	return c
}

func Connect(ctx context.Context, user string, pass string, dbHost string) *sql.DB {
	deb, err := sql.Open("mysql", user+":"+pass+"@"+dbHost)

	if err := deb.Ping(); err != nil {
		deb.Close()
	}

	if err != nil {
		panic(err.Error())
	} else {
		return deb
	}

}

type Quote_Availability struct {
	Tmpsession_id        string
	Tmpresort            string
	Tmpbegin_date        string
	Tmp_days_stay        int
	Tmpminrooms          int
	Tmpadults            int
	Tmpchildren          int
	P_infant             int
	Tmpgateway           string
	Tmppromo_id          int
	Pi_resort_type       string
	Pi_region            string
	Pi_state             string
	Pi_zone              int
	Pi_country           string
	Pi_ssg_no            int
	Pi_res_insert_source string
	Pi_rate_structure    string
	Pi_priv_structure    string
	P_mesage             string
	Pi_air_exist         string
	Pi_name_type         string
	Pi_name_code         string
	Pi_RoomCat           string
	Pi_priv_sale_code    string
}

type ReturnValues struct {
	SESSION_ID                string
	ROOM_CATEGORY             string
	Room_Class                string
	DESCRIPTION               string
	PAX_BASE_RATE             float64
	CHILD_BASE_RATE           float64
	PROMO_AMOUNT              float64
	AFTER_PROMO_RATE          float64
	STATUS                    string
	SIGNET_APPLY              string
	SIGNET_GAIN               string
	SIGNET_POINTS             float64
	WEB_DESCRIPTION           string
	WEB_IMAGE_CODE            float64
	HANDICAP                  string
	BUTLER                    string
	CONCIERGE                 string
	CATEGORY_VIEW             string
	ROOM_SERVICE              string
	MAP_LOCATION              string
	VRX                       string
	FLOORPLAN_IMAGE_CODE      float64
	ADDITIONAL_IMAGE_CODES    string
	BEDDING_NAME              string
	PRIVATE_CAR_TRANSFER      string
	ROLLS_ROYCE_TRANSFER      string
	CATEGORY_RANK             float64
	BMW_TRANSFER              string
	VIP_ARRIVAL_SERVICE       string
	DISCOUNT_ID               float64
	DISCOUNT_CODE             float64
	DISCOUNT_DESCRIPTION      string
	DISCOUNT_AMOUNT           float64
	PRIVATE_DISCOUNT_ID       float64
	PRIVATE_DISCOUNT_CODE     float64
	PRIVATE_DISCOUNT_DESC     string
	PRIVATE_DISCOUNT_AMT      float64
	DEPOSIT_DUE_DATE          string
	DEPOSIT_AMOUNT            float64
	BALANCE_DUE_DATE          string
	DEPOSIT_REFUNDABLE_YN     string
	DAYS_TO_BALANCE_DEPOSIT   float64
	Balance_deposit_amount    float64
	Available_rooms           float64
	Webspc_discount_id        string
	Webspc_disc_amt           float64
	Webspc_disc_desc          string
	Webspc_disc_code          string
	Full_payment_disc_id      float64
	Full_payment_disc_amt     float64
	FULL_PAYMENT_DISC_FORMULA string
	Full_payment_disc_factor  string
	NON_REFUNDABLE_DEPOSIT    string
	Max_occupancy             float64
}

type Stage_Params struct {
	Num         int
	Date        string
	Room_Cat    string
	Resort_Code string
}

type StageReturnValues struct {
	Resort          string
	Room_Cat        string
	Date            string
	Available_Rooms int
}

type Rate_Params struct {
	Tmpresort            string
	Tmpbegin_date        string
	Tmp_days_stay        int
	Pi_RoomCat           string
	Tmpminrooms          int
	Tmpadults            int
	Tmpchildren          int
	P_infant             int
	Tmpgateway           string
	Pi_state             string
	Pi_zone              int
	Pi_country           string
	Pi_ssg_no            int
	Pi_res_insert_source string
	Pi_rate_structure    string
	Pi_priv_sale_code    string
}

type RateReturnValues struct {
	Resort         string
	Room_Cat       string
	Region         string
	Zone           int64
	State          string
	Country        string
	Signet_Apply   string
	BeginDate      string
	Length         int64
	AfterPromoRate float64
	Status         string
	AvgAdultRate   float64
	PromotionId    string
	SessionId      string
	MaxOccupancy   int64
}

func (c *Client) Quote_availability(ctx context.Context, quote Quote_Availability) []ReturnValues {

	var returnArray []ReturnValues

	var returnStruct ReturnValues

	opdb := Connect(ctx, c.Config.DB_Stage_User, c.Config.DB_Stage_Pass, c.Config.DB_Stage_Host)

	defer opdb.Close()

	fmt.Println(quote.Tmpsession_id, quote.Tmpresort, quote.Tmpbegin_date, quote.Tmp_days_stay, quote.Tmpminrooms, quote.Tmpadults, quote.Tmpchildren, quote.P_infant, quote.Tmpgateway, quote.Tmppromo_id, quote.Pi_resort_type, quote.Pi_region, quote.Pi_state, quote.Pi_zone, quote.Pi_country, quote.Pi_ssg_no, quote.Pi_res_insert_source, quote.Pi_rate_structure, quote.Pi_priv_structure, quote.P_mesage, quote.Pi_air_exist, quote.Pi_name_type, quote.Pi_name_code)

	stmt, err := opdb.Query("CALL PRC.quote_availability(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);", quote.Tmpsession_id, quote.Tmpresort, quote.Tmpbegin_date, quote.Tmp_days_stay, quote.Tmpminrooms, quote.Tmpadults, quote.Tmpchildren, quote.P_infant, quote.Tmpgateway, quote.Tmppromo_id, quote.Pi_resort_type, quote.Pi_region, quote.Pi_state, quote.Pi_zone, quote.Pi_country, quote.Pi_ssg_no, quote.Pi_res_insert_source, quote.Pi_rate_structure, quote.Pi_priv_structure, quote.P_mesage, quote.Pi_air_exist, quote.Pi_name_type, quote.Pi_name_code)

	if err != nil {
		log.Println("STMT: " + err.Error())
	}
	defer stmt.Close()

	rows, err := opdb.Query("CALL `PRC`.`get_availability`(?)", quote.Tmpsession_id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var sessionId, roomCat, Additional_Image_codes, roomClass, description, status, signalApply, mapLocation, signetGain, webDescription, handicap, butler, concierge, categoryView, roomService, vrx, beddingName, privateCarTransfer, rollsRoyceTransfer, bmwTransfer, vipArrivalService, discountDescrption, privateDiscountDesc, balanceDueDate, depositRefundableYn, WDD, despositeDueDate, fullPaymentDiscountFormula, fullPaymentDiscFactor, NonRefDep, WDI, WDC sql.NullString
		var pax_base_rate, promoAmount, child_base_rate, afterPromoRate, signetPoint, webImageCode, floorPlanImageCodes, categoryRank, discountAmount, privateDiscountAmt, depositAmount, BalanceToDepositAmount, discountId, discountCode, privateDiscountID, privateDiscountCode, AvailableRooms, daysToBalanceDeposit, WDA, MaxOccp, full_payment_dic_id, full_payment_disc_amt sql.NullFloat64

		err := rows.Scan(&sessionId, &roomCat, &roomClass, &description, &pax_base_rate, &child_base_rate, &promoAmount, &afterPromoRate, &status, &signalApply, &signetGain, &signetPoint, &webDescription, &webImageCode, &handicap, &butler, &concierge, &categoryView, &roomService, &mapLocation, &vrx, &floorPlanImageCodes, &Additional_Image_codes, &beddingName, &privateCarTransfer, &rollsRoyceTransfer, &categoryRank, &bmwTransfer, &vipArrivalService, &discountId, &discountCode, &discountDescrption, &discountAmount, &privateDiscountID, &privateDiscountCode, &privateDiscountDesc, &privateDiscountAmt, &despositeDueDate, &depositAmount, &balanceDueDate, &depositRefundableYn, &daysToBalanceDeposit, &BalanceToDepositAmount, &AvailableRooms, &WDI, &WDC, &WDA, &WDD, &full_payment_dic_id, &full_payment_disc_amt, &fullPaymentDiscountFormula, &fullPaymentDiscFactor, &NonRefDep, &MaxOccp)

		if err != nil {
			log.Fatal(err)
		} else {
			if Additional_Image_codes.Valid {
				returnStruct.ADDITIONAL_IMAGE_CODES = Additional_Image_codes.String

			}
			if sessionId.Valid {
				returnStruct.SESSION_ID = sessionId.String
			}
			if roomCat.Valid {
				returnStruct.ROOM_CATEGORY = roomCat.String
			}
			if roomClass.Valid {
				returnStruct.Room_Class = roomClass.String
			}
			if description.Valid {
				returnStruct.DESCRIPTION = description.String
			}

			if pax_base_rate.Valid {
				returnStruct.PAX_BASE_RATE = pax_base_rate.Float64
			}

			if child_base_rate.Valid {
				returnStruct.CHILD_BASE_RATE = child_base_rate.Float64
			}

			if promoAmount.Valid {
				returnStruct.PROMO_AMOUNT = promoAmount.Float64
			}

			if afterPromoRate.Valid {
				returnStruct.AFTER_PROMO_RATE = afterPromoRate.Float64
			}

			if status.Valid {
				returnStruct.STATUS = status.String
			}
			if signalApply.Valid {
				returnStruct.SIGNET_APPLY = signalApply.String
			}

			if signetGain.Valid {
				returnStruct.SIGNET_GAIN = signetGain.String
			}

			if signetPoint.Valid {
				returnStruct.SIGNET_POINTS = signetPoint.Float64
			}

			if webDescription.Valid {
				returnStruct.WEB_DESCRIPTION = webDescription.String
			}
			if webImageCode.Valid {
				returnStruct.WEB_IMAGE_CODE = webImageCode.Float64
			}

			if handicap.Valid {
				returnStruct.HANDICAP = handicap.String
			}

			if butler.Valid {
				returnStruct.BUTLER = butler.String
			}

			if concierge.Valid {
				returnStruct.CONCIERGE = concierge.String
			}

			if categoryView.Valid {
				returnStruct.CATEGORY_VIEW = categoryView.String
			}

			if roomService.Valid {
				returnStruct.ROOM_SERVICE = roomService.String
			}

			if mapLocation.Valid {
				returnStruct.MAP_LOCATION = mapLocation.String
			}
			if vrx.Valid {
				returnStruct.VRX = vrx.String
			}
			if floorPlanImageCodes.Valid {
				returnStruct.FLOORPLAN_IMAGE_CODE = floorPlanImageCodes.Float64
			}
			if beddingName.Valid {
				returnStruct.BEDDING_NAME = beddingName.String
			}
			if privateCarTransfer.Valid {
				returnStruct.PRIVATE_CAR_TRANSFER = privateCarTransfer.String
			}
			if rollsRoyceTransfer.Valid {
				returnStruct.ROLLS_ROYCE_TRANSFER = rollsRoyceTransfer.String
			}
			if bmwTransfer.Valid {
				returnStruct.BMW_TRANSFER = bmwTransfer.String
			}
			if vipArrivalService.Valid {
				returnStruct.VIP_ARRIVAL_SERVICE = vipArrivalService.String
			}
			if discountId.Valid {
				returnStruct.DISCOUNT_ID = discountId.Float64
			}
			if discountCode.Valid {
				returnStruct.DISCOUNT_CODE = discountCode.Float64
			}
			if discountDescrption.Valid {
				returnStruct.DISCOUNT_DESCRIPTION = discountDescrption.String
			}
			if privateDiscountID.Valid {
				returnStruct.PRIVATE_DISCOUNT_ID = privateDiscountID.Float64
			}
			if privateDiscountCode.Valid {
				returnStruct.PRIVATE_DISCOUNT_CODE = privateDiscountCode.Float64
			}
			if privateDiscountDesc.Valid {
				returnStruct.PRIVATE_DISCOUNT_DESC = privateDiscountDesc.String
			}
			if balanceDueDate.Valid {
				returnStruct.BALANCE_DUE_DATE = balanceDueDate.String
			}
			if depositRefundableYn.Valid {
				returnStruct.DEPOSIT_REFUNDABLE_YN = depositRefundableYn.String
			}
			if WDD.Valid {
				returnStruct.Webspc_disc_desc = WDD.String
			}
			if despositeDueDate.Valid {
				returnStruct.DEPOSIT_DUE_DATE = despositeDueDate.String
			}
			if fullPaymentDiscountFormula.Valid {
				returnStruct.FULL_PAYMENT_DISC_FORMULA = fullPaymentDiscountFormula.String
			}
			if fullPaymentDiscFactor.Valid {
				returnStruct.Full_payment_disc_factor = fullPaymentDiscFactor.String
			}
			if NonRefDep.Valid {
				returnStruct.NON_REFUNDABLE_DEPOSIT = NonRefDep.String
			}
			if AvailableRooms.Valid {
				returnStruct.Available_rooms = AvailableRooms.Float64
			}
			if daysToBalanceDeposit.Valid {
				returnStruct.DAYS_TO_BALANCE_DEPOSIT = daysToBalanceDeposit.Float64
			}
			if WDI.Valid {
				returnStruct.Webspc_discount_id = WDI.String
			}
			if WDC.Valid {
				returnStruct.Webspc_disc_code = WDC.String
			}
			if WDA.Valid {
				returnStruct.Webspc_disc_amt = WDA.Float64
			}
			if MaxOccp.Valid {
				returnStruct.Max_occupancy = MaxOccp.Float64
			}
			if full_payment_dic_id.Valid {
				returnStruct.Full_payment_disc_id = full_payment_dic_id.Float64
			}
			if full_payment_disc_amt.Valid {
				returnStruct.Full_payment_disc_amt = full_payment_disc_amt.Float64
			}
		}
		returnArray = append(returnArray, returnStruct)
	}

	return returnArray

}

func (c *Client) Test_availability(ctx context.Context, params Stage_Params) []StageReturnValues {

	var values []StageReturnValues

	db := Connect(ctx, c.Config.DB_Stage_User, c.Config.DB_Stage_Pass, c.Config.DB_Stage_Host)

	defer db.Close()

	var roomCat string

	if params.Room_Cat != "" {
		roomCat = params.Room_Cat
	} else {
		roomCat = "all"
	}

	rows, err := db.Query("CALL PRC.test_availability(?,?,?,?)", params.Num, params.Date, roomCat, params.Resort_Code)

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			var tempReturnValue StageReturnValues
			var resort, room_cat, reservation_date sql.NullString
			var available_rooms sql.NullInt64
			err2 := rows.Scan(&resort, &room_cat, &reservation_date, &available_rooms)
			if err2 != nil {
				log.Println(err2)
			} else {
				if resort.Valid {
					tempReturnValue.Resort = resort.String
				}
				if room_cat.Valid {
					tempReturnValue.Room_Cat = room_cat.String
				}
				if reservation_date.Valid {
					tempReturnValue.Date = reservation_date.String
				}
				if available_rooms.Valid {
					tempReturnValue.Available_Rooms = int(available_rooms.Int64)
				}
			}
			values = append(values, tempReturnValue)
		}
	}
	return values
}

func (c *Client) Test_rate(ctx context.Context, params Rate_Params) []RateReturnValues {

	var values []RateReturnValues

	db := Connect(ctx, c.Config.DB_Stage_User, c.Config.DB_Stage_Pass, c.Config.DB_Stage_Host)

	defer db.Close()

	var roomCat string

	if params.Pi_RoomCat != "" {
		roomCat = params.Pi_RoomCat
	} else {
		roomCat = "all"
	}

	rows, err := db.Query("CALL PRC.TestRates(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);",
		params.Tmpresort, params.Tmpbegin_date, params.Tmp_days_stay, roomCat, params.Tmpminrooms, params.Tmpadults, params.Tmpchildren, params.P_infant, params.Tmpgateway, params.Pi_state, params.Pi_zone, params.Pi_country,
		params.Pi_ssg_no, params.Pi_res_insert_source, params.Pi_rate_structure, params.Pi_priv_sale_code)

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			var resort, room_cat, region, resortType, state, country, signetApply, status, sessionId sql.NullString
			var zone, length, maxOccupancy sql.NullInt64
			var afterPromoRate, avgAdultRate sql.NullFloat64

			var tempRate RateReturnValues

			err := rows.Scan(&resort, &room_cat, &region, &resortType, &zone, &state, &country, &signetApply, &length, &afterPromoRate, &status, &avgAdultRate, &sessionId, &maxOccupancy)

			if err != nil {
				log.Println(err)
			} else {
				if resort.Valid {
					tempRate.Resort = resort.String
				}
				if room_cat.Valid {
					tempRate.Room_Cat = room_cat.String
				}
				if region.Valid {
					tempRate.Region = region.String
				}
				if zone.Valid {
					tempRate.Zone = zone.Int64
				}
				if state.Valid {
					tempRate.State = state.String
				}
				if country.Valid {
					tempRate.Country = country.String
				}
				if signetApply.Valid {
					tempRate.Signet_Apply = signetApply.String
				}

				if length.Valid {
					tempRate.Length = length.Int64
				}
				if afterPromoRate.Valid {
					tempRate.AfterPromoRate = afterPromoRate.Float64
				}
				if status.Valid {
					tempRate.Status = status.String
				}
				if avgAdultRate.Valid {
					tempRate.AvgAdultRate = avgAdultRate.Float64
				}
				if sessionId.Valid {
					tempRate.SessionId = sessionId.String
				}
				if maxOccupancy.Valid {
					tempRate.MaxOccupancy = maxOccupancy.Int64
				}
			}
			values = append(values, tempRate)
		}
	}
	return values
}
