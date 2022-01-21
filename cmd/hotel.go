/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// hotelCmd represents the hotel command
var hotelCmd = &cobra.Command{
	Use:   "hotel",
	Short: "Check hotel stats or initialize a hotel with 0 occupants",
	Long: `Displays hotel statistics when the show flag is active.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(hotelCmd)
}