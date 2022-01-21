/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		deleteGuest()
	},
}

var toBeDeleted int

func init() {
	guestCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().IntVar(&toBeDeleted, "id", -1, "specifies the id of the guest to be deleted.")
	deleteCmd.MarkFlagRequired("id")
}

func deleteGuest()  {
	url := fmt.Sprintf("http://localhost:4000/guest/%d", toBeDeleted)
	body := requestDeleteGuest(url)
	log.Println(string(body))

}

func requestDeleteGuest(url string) []byte {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
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
