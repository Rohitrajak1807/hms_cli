/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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

	"github.com/spf13/cobra"
)

// checkoutCmd represents the checkout command
var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

func changeCheckOutDate()  {
	url := fmt.Sprintf("http://localhost:4000/guest/%d", guestId)
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