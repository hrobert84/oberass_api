package Http_Client

import (
	"Rass/internal/compare_tool"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	mux "github.com/julienschmidt/httprouter"
)

type Route struct {
	Method string
	Path   string
	Handle mux.Handle // httprouter package as mux
}

type Routes []Route

var routes = Routes{
	Route{
		"POST",
		"/run",
		run_test,
	},
	Route{
		"POST",
		"/automatic",
		automatic,
	},
}

type Params struct {
	Resort_Code string `json:"resort_codes"`
	Start_date  string `json:"start_date"`
	Length_stay int    `json:"length_stay"`
}

type FullYearTest struct {
	Year int `json:"year"`
}

type ResortCode struct {
	Id string `json:"code"`
}

// Function that call the API to run
func RunAPI() {
	router := newRouter()
	log.Println("Running in port 5152. Go to /run to execute")
	log.Fatal(http.ListenAndServe(":5152", router))
}

// Constructor for the router
func newRouter() *mux.Router {

	router := mux.New()

	for _, route := range routes {

		router.Handle(route.Method, route.Path, route.Handle)

	}

	return router
}

func automatic(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	body, err := io.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		log.Println(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	request_arguments := FullYearTest{}

	if err := json.Unmarshal(body, &request_arguments); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)

		w.Write([]byte("couldn't parse request body"))

		if err := json.NewEncoder(w).Encode(err); err != nil {

			log.Println(err)

			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}

		}
	} else {
		msg := compare_tool.Starter(request_arguments.Year)
		json.NewEncoder(w).Encode(msg)
	}

}

// Function that handles the request
func run_test(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	body, err := io.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		log.Println(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	request_arguments := Params{}

	if err := json.Unmarshal(body, &request_arguments); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)

		w.Write([]byte("couldn't parse request body"))

		if err := json.NewEncoder(w).Encode(err); err != nil {

			log.Println(err)

			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}

		}
	} else {
		_, err := time.Parse("2006-01-02", request_arguments.Start_date)

		if err != nil {
			resp := make(map[string]string)
			resp["message"] = "start_date format has to be yyyy-mm-dd"
			jsonResp, _ := json.Marshal(resp)
			w.Write(jsonResp)
			return
		}

		values := compare_tool.Data{
			ResortCode: request_arguments.Resort_Code,
			BeginDate:  request_arguments.Start_date,
			Length:     request_arguments.Length_stay,
		}

		msg, fails := compare_tool.Rate_Compare_Test(request_arguments.Resort_Code, request_arguments.Start_date, request_arguments.Length_stay)

		values.Data = msg
		values.Fails = fails

		log.Println(values)

		json.NewEncoder(w).Encode(values)
		return
	}

}
