/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
Example: hms_cli guest create --check-in="2022-01-21" --check-out="2022-01-26" --name="abhay"
Creates a new guest
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/Rohitrajak1807/hms_cli/apiroutes"
	"github.com/Rohitrajak1807/hms_cli/models"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a guest in the hotel database by adding the required details.",
	Long: `This command helps user to create a guest by adding details.
			Note that the payment tab will update automatically depending on
			the number of days the guest is staying in the hotel. 
			
			Usage Example: hms_cli guest create --check-in --check-out --name
			
				--check-in		Enter the checkin date by using this flag in the command. The date should in the format YYYY-MM-DD.
				--check-out		Enter the checkout date by using this flag in the command. The date should in the format YYYY-MM-DD
				--name			Enter the name of the guest by using this flag in command. Example: --name="Jacob"`,

	Run: func(cmd *cobra.Command, args []string) {
		createGuest()
	},
}

var guest models.GuestIn

func init() {
	guestCmd.AddCommand(createCmd)
	createCmd.Flags().StringVar(&guest.Name, "name", "", "specify the guest name.")
	createCmd.Flags().StringVar(&guest.CheckInDate, "check-in", "", "specify the check in date.")
	createCmd.Flags().StringVar(&guest.CheckOutDate, "check-out", "", "specify the check out date.")
	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("check-in")
	createCmd.MarkFlagRequired("check-out")
}

func createGuest() {
	url := apiroutes.GuestCreateRoute
	_, checkInErr := time.Parse("2006-01-02", guest.CheckInDate)
	if checkInErr != nil {
		log.Fatal(checkInErr)
	}
	_, checkOutErr := time.Parse("2006-01-02", guest.CheckOutDate)
	if checkOutErr != nil {
		log.Fatal(checkOutErr)
	}

	jsonString, err := json.Marshal(guest)
	if err != nil {
		log.Fatal(err)
	}
	responseBytes := makeCreateGuestRequest(url, jsonString)
	log.Println("response:", string(responseBytes))
}

func makeCreateGuestRequest(url string, requestBody []byte) []byte {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal(err)
	}
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	log.Println("StatusCode:", response.StatusCode)
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return responseBytes
}
