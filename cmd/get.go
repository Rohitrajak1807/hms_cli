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

	"github.com/Rohitrajak1807/hms_cli/apiroutes"
	"github.com/Rohitrajak1807/hms_cli/models"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get details of a particluar guest by providing the Guest ID.",
	Long: `This command helps the user to get the details of a particular
			guest by giving the ID as the argument.
			
			Example: hms_cli guest get --id
			
				--id: ID of that particular guest whose details needs to be fetched.`,
	
	Run: func(cmd *cobra.Command, args []string) {
		getAGuest()
	},
}

var id int

func init() {
	guestCmd.AddCommand(getCmd)
	getCmd.Flags().IntVar(&id, "id", -1, "accepts an integer id for the guest")
	getCmd.MarkFlagRequired("id")
}

func getAGuest() {
	url := fmt.Sprintf("%s/%d", apiroutes.GuestGetRoute,id)
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
	if response.StatusCode != http.StatusOK {
		log.Fatal(response.Status)
	}

	return responseBytes
}
