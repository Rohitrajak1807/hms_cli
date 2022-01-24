/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
Example: hms_cli hotel show
Shows the stats of hotel
Init a hotel with 0 occupants
*/
package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Rohitrajak1807/hms_cli/models"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "hms_cli hotel show",
	Long:  `This command shows the stats of hotel like number of room occupied, name of the hotel etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		showHotelStats()
	},
}

func init() {
	hotelCmd.AddCommand(showCmd)
}

func showHotelStats() {
	url := "http://localhost:4000/hotel"
	responseBytes := getHotelData(url)
	hotel := models.Hotel{}
	if err := json.Unmarshal(responseBytes, &hotel); err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", hotel)
}

func getHotelData(url string) []byte {
	req, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return responseBytes
}
