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

	"github.com/Rohitrajak1807/hms_cli/models"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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