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

	"github.com/Rohitrajak1807/hms_cli/models"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "hms_cli guest create",
	Long:  `Add guest data by adding flags --check-in="" --check-out="" --name=""`,
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
	url := "http://localhost:4000/guest"
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
