package cmd

import (
	"github.com/spf13/cobra"
)

var (
	username string
	email    string
	password string
	roleID   string
	status   string
)

func init() {

	// createCmd.Flags().StringVar(
	// 	&username,                   // Pointer to store value
	// 	string(parser.FormUsername), // Flag name
	// 	"",                          // Default value
	// 	"Username for user",         // Description
	// )

	// createCmd.Flags().StringVar(
	// 	&email,                   // Pointer to store value
	// 	string(parser.FormEmail), // Flag name
	// 	"",                       // Shorthand (optional)
	// 	"email for user",         // Description
	// )

	// createCmd.Flags().StringVar(
	// 	&password,                   // Pointer to store value
	// 	string(parser.FormPassword), // Flag name
	// 	"",                          // Default value
	// 	"password for user",         // Description
	// )

	// createCmd.Flags().StringVar(
	// 	&status,                   // Pointer to store value
	// 	string(parser.FormStatus), // Flag name
	// 	"",                        // Default value
	// 	"status for user",         // Description
	// )

	// createCmd.Flags().StringVar(
	// 	&roleID,                   // Pointer to store value
	// 	string(parser.FormRoleID), // Flag name
	// 	"",                        // Default value
	// 	"role_id for user",        // Description
	// )
	// createCmd.MarkFlagRequired(string(parser.FormUsername))
	// createCmd.MarkFlagRequired(string(parser.FormEmail))
	// createCmd.MarkFlagRequired(string(parser.FormPassword))
	// createCmd.MarkFlagRequired(string(parser.FormRoleID))
	// createCmd.MarkFlagRequired(string(parser.FormStatus))

	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new user",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
