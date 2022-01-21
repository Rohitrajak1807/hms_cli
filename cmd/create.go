/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		createGuest()
	},
}

var guest models.GuestIn

func init() {
	guestCmd.AddCommand(createCmd)
	guestCmd.Flags().StringVar(&guest.Name, "name", "", "specify the guest name.")
	guestCmd.Flags().StringVar(&guest.CheckInDate, "check-in", "", "specify the check in date")
	guestCmd.Flags().StringVar(&guest.CheckInDate, "check-out", "", "specify the check out date")
	guestCmd.MarkFlagRequired("name")
	guestCmd.MarkFlagRequired("check-in")
	guestCmd.MarkFlagRequired("check-out")
}

func createGuest()  {
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
