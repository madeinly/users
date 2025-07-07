/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// var userID string

func init() {

	// getCmd.Flags().StringVar(
	// 	&userID,               // Pointer to store value
	// 	string(parser.FormID), // Flag name
	// 	"",                    // Default value
	// 	"UserID for user",     // Description
	// )

	// getCmd.MarkFlagRequired(string(parser.FormID))
	rootCmd.AddCommand(getCmd)

}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get a single user by its id",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
