/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
Example: hms_cli guest delete --id=19
delete a guest by id
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
	Short: "hms_cli guest delete",
	Long:  `Delete a particular guest by adding flag --id=`,
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

func deleteGuest() {
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
