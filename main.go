package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// define json
type DateInfo struct {
	Now                 Now `json:"_now"`
	Century             int
	Date                string
	Date_ja             string
	Day                 int
	Error               string `json:"error"`
	Eto                 string
	Eto_kana            string
	Gengo               string
	Gengo_full          string
	Holiday             string
	Julian              int
	Month               int
	Month_en            string
	Month_end           int
	Month_ja            string
	Moon                int
	Moon_en             string
	Moon_ja             string
	Old_date            string
	Old_day             int
	Old_leap            bool
	Old_month           int
	Old_year            int
	Rokuyo              string
	Sunrise             string
	Sunset              string
	Timezone            string
	Week                int
	Week_en             string
	Week_ja             string
	Week_number         int
	Week_number_of_year int
	Year                int
	Year_ja             int
}

type Now struct {
	Date     string
	Datetime string
	Day      int
	Hour     int
	Minute   int
	Month    int
	Second   int
	Week     int
	Week_en  string
	Week_ja  string
	Year     int
}

func getDateInfo(date string) *DateInfo {
	const url = "https://dateinfoapi.appspot.com/v1"
	req, _ := http.NewRequest("GET", url+"?date="+date, nil)
	client := new(http.Client)

	res, _ := client.Do(req)
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	dateInfo := new(DateInfo)
	json.Unmarshal([]byte(body), &dateInfo)

	return dateInfo
}

func main() {
	const layout = "2006-01-02"

	// Options
	opt := flag.String("d", time.Now().Format(layout), "Set date to know rokuyo")
	flag.Parse()

	// check date format
	t, err := time.Parse(
		"2006-01-02",
		*opt)
	if err != nil {
		fmt.Println("Error option")
		fmt.Println("Date layout is " + `%yyyy-%mm-%dd`)
		return
	}

	dateInfo := getDateInfo(t.Format(layout))
	fmt.Println(dateInfo.Date + " " + dateInfo.Rokuyo)
}
