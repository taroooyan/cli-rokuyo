package main

import (
  "fmt"
  "time"
  "net/http"
  "io/ioutil"
  "encoding/json"
)

// define json
type DateInfo struct {
  Now Now `json:"_now"`
  Century int
  Date string
  Date_ja string
  Day int
  Error string `json:"error"`
  Eto string
  Eto_kana string
  Gengo string
  Gengo_full string
  Holiday string
  Julian int
  Month int
  Month_en string
  Month_end int
  Month_ja string
  Moon int
  Moon_en string
  Moon_ja string
  Old_date string
  Old_day int
  Old_leap bool
  Old_month int
  Old_year int
  Rokuyo string
  Sunrise string
  Sunset string
  Timezone string
  Week int
  Week_en string
  Week_ja string
  Week_number int
  Week_number_of_year int
  Year int
  Year_ja int
}

type Now struct {
  Date string
  Datetime string
  Day int
  Hour int
  Minute int
  Month int
  Second int
  Week int
  Week_en string
  Week_ja string
  Year int
}

func getDateInfo(date string) string {
  const url = "https://dateinfoapi.appspot.com/v1"
  req, _:= http.NewRequest("GET", url+"?date="+date, nil)
  client := new(http.Client)

  res, _ := client.Do(req)
  defer res.Body.Close()

  body, _ := ioutil.ReadAll(res.Body)

  var dateInfo = new(DateInfo)
  json.Unmarshal([]byte(body), &dateInfo)

  fmt.Println(dateInfo)
  return "end"
}

func main() {
  t := time.Now()
  const layout = "2006-01-02"
  fmt.Println(t.Format(layout))
  getDateInfo(t.Format(layout))
}
