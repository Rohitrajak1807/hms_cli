/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
Example: hms_cli hotel init --cost=1000 --name="Hotel 1" --rooms=5
Init a hotel with 0 occupants
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Rohitrajak1807/hms_cli/models"
	"github.com/spf13/cobra"
)

var hotel models.HotelIn

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the hotel with 0 occupants ",
	Long: `This command helps in creating a hotel.
	
			Example: hms_cli hotel create --cost --name --rooms
				
				--cost		This flag represents the cost of a single room which can be changed by the user
							according to this only, the total payment will be calculated
				--name		Specify the name of the hotel by using this flag
				--rooms		Specify how many rooms are available in the desired hotel.`,

	Run: func(cmd *cobra.Command, args []string) {
		initHotel()
	},
}

func init() {
	hotelCmd.AddCommand(initCmd)
	initCmd.Flags().StringVar(&hotel.Name, "name", "", "Set hotel name")
	initCmd.Flags().IntVar(&hotel.TotalRooms, "rooms", 0, "Rooms in the hotel")
	initCmd.Flags().IntVar(&hotel.CostPerDay, "cost", 0, "cost per day of hotel room")
	initCmd.MarkFlagRequired("name")
	initCmd.MarkFlagRequired("rooms")
	initCmd.MarkFlagRequired("cost")
}

func initHotel() {
	url := "http://localhost:4000/hotel"
	jsonStr, err := json.Marshal(hotel)
	if err != nil {
		log.Fatal(err)
	}
	body := makeCreateRequest(url, jsonStr)
	log.Println("response:", string(body))
}

func makeCreateRequest(url string, requestBody []byte) []byte {

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	log.Println("StatusCode:", res.StatusCode)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}
