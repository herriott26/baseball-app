package main

import (
    "fmt"
    "log"
	"net/http"
	"encoding/json"
	"time"
)

type WeeklyCallUps struct {
	TransactionAll struct {
		CopyRight    string `json:"copyRight"`
		QueryResults struct {
			Created   string `json:"created"`
			TotalSize string `json:"totalSize"`
			Row       []struct {
				TransDateCd          string `json:"trans_date_cd"`
				FromTeamID           string `json:"from_team_id"`
				OrigAsset            string `json:"orig_asset"`
				FinalAssetType       string `json:"final_asset_type"`
				Player               string `json:"player"`
				ResolutionCd         string `json:"resolution_cd"`
				FinalAsset           string `json:"final_asset"`
				NameDisplayFirstLast string `json:"name_display_first_last"`
				TypeCd               string `json:"type_cd"`
				NameSort             string `json:"name_sort"`
				ResolutionDate       string `json:"resolution_date"`
				ConditionalSw        string `json:"conditional_sw"`
				Team                 string `json:"team"`
				Type                 string `json:"type"`
				NameDisplayLastFirst string `json:"name_display_last_first"`
				TransactionID        string `json:"transaction_id"`
				TransDate            string `json:"trans_date"`
				EffectiveDate        string `json:"effective_date"`
				PlayerID             string `json:"player_id"`
				OrigAssetType        string `json:"orig_asset_type"`
				FromTeam             string `json:"from_team"`
				TeamID               string `json:"team_id"`
				Note                 string `json:"note"`
			} `json:"row"`
		} `json:"queryResults"`
	} `json:"transaction_all"`
}

func main() {

	//API Call http://lookup-service-prod.mlb.com/json/named.transaction_all.bam?sport_code='mlb'&start_date='20190414'&end_date='20190419'

	//get starting and end time for call-up report
	now := time.Now()
	enddate := now.Format("20060102")
	start := now.AddDate(0,0,-7)
	startdate := start.Format("20060102")

	fmt.Println("Starting Weelky Callup Aggregator...")
	fmt.Println("Calling http://lookup-service-prod.mlb.com API")
	url := fmt.Sprintf("http://lookup-service-prod.mlb.com/json/named.transaction_all.bam?sport_code='mlb'&start_date='%s'&end_date='%s'", startdate, enddate)
	fmt.Println(url)

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	fmt.Println("Calling API complete")

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record WeeklyCallUps

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	// This will iterate through all the weekly callups
	fmt.Println("Current date: ", enddate)
	fmt.Println("Start date: ", startdate)
	fmt.Println("Date created:", record.TransactionAll.QueryResults.Created)
	fmt.Println("Date created:", record.TransactionAll.QueryResults.TotalSize)
	fmt.Println("Player 0:", record.TransactionAll.QueryResults.Row[0].Player)
	fmt.Println("Player 0 ID:", record.TransactionAll.QueryResults.Row[0].PlayerID)
	fmt.Println("Player 0 Transaction Date:", record.TransactionAll.QueryResults.Row[0].TransDate)
	fmt.Println("Player 0 Team:", record.TransactionAll.QueryResults.Row[0].Team)
	fmt.Println("Player 0 Notes:", record.TransactionAll.QueryResults.Row[0].Note)

	player25 := record.TransactionAll.QueryResults.Row[25]
	fmt.Println(player25)
}