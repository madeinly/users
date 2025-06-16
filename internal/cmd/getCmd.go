/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/madeinly/users/internal/parser"
	"github.com/madeinly/users/internal/repo"
	"github.com/spf13/cobra"
)

var userID string

func init() {

	getCmd.Flags().StringVar(
		&userID,               // Pointer to store value
		string(parser.FormID), // Flag name
		"",                    // Default value
		"UserID for user",     // Description
	)

	getCmd.MarkFlagRequired(string(parser.FormID))
	rootCmd.AddCommand(getCmd)

}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get a single user by its id",
	Run: func(cmd *cobra.Command, args []string) {
		v := parser.NewUserParser()

		userID := v.ValidateID(userID)

		if v.HasErrors() {
			userErrors, _ := json.MarshalIndent(v.Errors, "", " ")
			fmt.Println(string(userErrors))
			return
		}

		user, err := repo.GetUserByID(userID)

		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(user)
	},
}
