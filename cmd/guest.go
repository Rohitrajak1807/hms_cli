/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
The guest subcommand 
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// guestCmd represents the guest command
var guestCmd = &cobra.Command{
	Use:   "guest",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}

func init() {
	rootCmd.AddCommand(guestCmd)
}
