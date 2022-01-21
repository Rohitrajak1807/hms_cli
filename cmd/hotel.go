/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// hotelCmd represents the hotel command
var hotelCmd = &cobra.Command{
	Use:   "hotel",
	Short: "Check hotel stats or initialize a hotel with 0 occupants",
	Long: `Displays hotel statistics when the show flag is active.`,
	Run: func(cmd *cobra.Command, args []string) {
		showFlag, _ := cmd.Flags().GetBool("show")
		if showFlag {
			showHotelStats()
		}
	},
}

func init() {
	rootCmd.AddCommand(hotelCmd)
	hotelCmd.Flags().BoolP("show", "s", false, "Show hotel stats")

}

type Hotel struct {
	Name string `json:"name"`
	TotalRooms int `json:"total_rooms"`
	OccupiedRooms int `json:"occupied_rooms"`
	CostPerDay int `json:"cost_per_day"`
}

func showHotelStats()  {
	url := "http://localhost:4000/hotel"
	responseBytes := getHotelData(url)
	hotel := Hotel{}
	if err := json.Unmarshal(responseBytes, &hotel); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Name: %v\n", hotel.Name)
	fmt.Printf("TotalRooms: %v\n", hotel.TotalRooms)
	fmt.Printf("OccupiedRooms: %v\n", hotel.OccupiedRooms)
	fmt.Printf("CostPerDay: %v\n", hotel.CostPerDay)
}

func getHotelData(url string) []byte  {
	req, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		log.Fatal("Cannot reach server")
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
