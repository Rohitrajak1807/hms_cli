/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
Example: hms_cli guest checkout --date="2022-01-31" --id=16
Edit checkout date for a guest
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/Rohitrajak1807/hms_cli/apiroutes"
	"github.com/spf13/cobra"
)

// checkoutCmd represents the checkout command
var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "This command enables the guest to update the checkout date.",
	Long: `Note that the total payment will be updated automatically depending on the checkout date.
			If the new checkout date is before the old checkout date, then payment amount will decrease
			and if it is after checkout date, amount will increase.
			
			Update Checkout date by adding flag --date --id
			
				--date		provide checkout date here in this format: YYYY-MM-DD
				--id		enter the id of that particular guest whose checkout date needs to be altered`,
	Run: func(cmd *cobra.Command, args []string) {
		changeCheckOutDate()
	},
}

var newCheckOut string
var guestId int

func init() {
	guestCmd.AddCommand(checkoutCmd)
	checkoutCmd.Flags().IntVar(&guestId, "id", -1, "Specifies the id of a guest.")
	checkoutCmd.Flags().StringVar(&newCheckOut, "date", "", "Change the checkout date of a guest.")
	checkoutCmd.MarkFlagRequired("id")
	checkoutCmd.MarkFlagRequired("date")
}

func changeCheckOutDate() {
	url := fmt.Sprintf("%s/%d", apiroutes.GuestCheckoutRoute, guestId)
	_, err := time.Parse("2006-01-02", newCheckOut)
	if err != nil {
		log.Fatal(err)
	}
	jsonStr, err := json.Marshal(map[string]string{
		"date": newCheckOut,
	})
	if err != nil {
		log.Fatal(err)
	}
	response := requestCheckOutChange(url, jsonStr)
	log.Println(string(response))
}

func requestCheckOutChange(url string, requestBody []byte) []byte {
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal(err)
	}
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	log.Println("StatusCode:", response.StatusCode)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}
