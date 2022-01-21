/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		showAllGuests()
	},
}

func init() {
	guestCmd.AddCommand(allCmd)
}

func showAllGuests() {
	url := "http://localhost:4000/guests"
	responseBytes := getAllGuests(url)
	var guests []models.Guest
	err := json.Unmarshal(responseBytes, &guests)
	if err != nil {
		log.Fatal(err)
	}
	for _, guest := range guests {
		log.Printf("%+v\n", guest)
	}
}

func getAllGuests(url string) []byte  {
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