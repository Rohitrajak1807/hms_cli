/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
Example: hms_cli guest get --id=1
get a guest by id
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Rohitrajak1807/hms_cli/models"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "hms_cli guest get",
	Long:  `Get details of a particluar guest by adding flag --id=`,
	Run: func(cmd *cobra.Command, args []string) {
		getAGuest()
	},
}

var id int

func init() {
	guestCmd.AddCommand(getCmd)
	guestCmd.LocalFlags().IntVar(&id, "id", -1, "accepts an integer id for the guest")
	guestCmd.MarkFlagRequired("id")
}

func getAGuest() {
	url := fmt.Sprintf("http://localhost:4000/guest/%d", id)
	body := getGuestData(url)
	guest := models.Guest{}
	if err := json.Unmarshal(body, &guest); err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", guest)
}

func getGuestData(url string) []byte {
	req, err := http.NewRequest(http.MethodGet, url, nil)
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
