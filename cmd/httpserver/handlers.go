package httpserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/scardozos/add-weeks-tool/cmd/dbclient"
)

type daterange struct {
	StartDate *dbclient.Date
	EndDate   *dbclient.Date
}

type DatesRequest struct {
	DateRange  *daterange
	SingleDate *dbclient.Date
}

type GetDatesResponse struct {
	Res []dbclient.Date `json:"result"`
}

type InsertStaticWeeksResponse struct {
	Res []dbclient.Date `json:"result"`
}

type ErrorInsertStaticWeeksResponse struct {
	Err error `json:"error"`
}

func (s *WeekHandlerRouterContext) GetDates(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res, err := s.DbClient.GetStaticWeeks()
	if err != nil {
		log.Printf("Error found in GetDates: %v", err)
	}
	response := &GetDatesResponse{
		Res: res,
	}
	jsonRes, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error found in GetDates: %v", err)
	}
	w.Header().Set("content-type", "application/json")
	fmt.Fprint(w, string(jsonRes))
}

func (s *WeekHandlerRouterContext) InsertDates(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req DatesRequest
	var res InsertStaticWeeksResponse
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.SingleDate != nil && req.DateRange == nil {
		date, err := s.DbClient.SetStaticWeek(*req.SingleDate)
		if err != nil {
			log.Printf("Error encountered: %v", err)
			fmt.Fprintf(w, "{\"error\":\"%v\"}", err.Error())
			return
		}
		res.Res = append(res.Res, date)
	}
	if req.DateRange != nil && req.SingleDate == nil {
		log.Print("Multiple Dates")
	}
	jsonRes, err := json.Marshal(res)
	if err != nil {
		log.Printf("Error found in InsertDates: %v", err)
		fmt.Fprintf(w, "{\"error\":\"%v\"}", err.Error())
		return

	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%v", string(jsonRes))
	//fmt.Fprint(w, "{\"error\":\"test error\"}")
}
