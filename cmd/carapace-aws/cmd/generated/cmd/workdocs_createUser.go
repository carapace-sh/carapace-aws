package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_createUserCmd = &cobra.Command{
	Use:   "create-user",
	Short: "Creates a user in a Simple AD or Microsoft AD directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_createUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workdocs_createUserCmd).Standalone()

		workdocs_createUserCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
		workdocs_createUserCmd.Flags().String("email-address", "", "The email address of the user.")
		workdocs_createUserCmd.Flags().String("given-name", "", "The given name of the user.")
		workdocs_createUserCmd.Flags().String("organization-id", "", "The ID of the organization.")
		workdocs_createUserCmd.Flags().String("password", "", "The password of the user.")
		workdocs_createUserCmd.Flags().String("storage-rule", "", "The amount of storage for the user.")
		workdocs_createUserCmd.Flags().String("surname", "", "The surname of the user.")
		workdocs_createUserCmd.Flags().String("time-zone-id", "", "The time zone ID of the user.")
		workdocs_createUserCmd.Flags().String("username", "", "The login name of the user.")
		workdocs_createUserCmd.MarkFlagRequired("given-name")
		workdocs_createUserCmd.MarkFlagRequired("password")
		workdocs_createUserCmd.MarkFlagRequired("surname")
		workdocs_createUserCmd.MarkFlagRequired("username")
	})
	workdocsCmd.AddCommand(workdocs_createUserCmd)
}
