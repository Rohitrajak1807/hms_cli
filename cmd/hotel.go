/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// hotelCmd represents the hotel command
var hotelCmd = &cobra.Command{
	Use:   "hotel",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		showFlag, _ := cmd.Flags().GetBool("show")
		if showFlag {
			showHotelStats()
		}
	},
}

func init() {
	rootCmd.AddCommand(hotelCmd)
	hotelCmd.Flags().BoolP("show", "s", false, "Show hotel stats")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hotelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hotelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func showHotelStats()  {
	
}
