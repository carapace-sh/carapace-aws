package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_createUserCmd = &cobra.Command{
	Use:   "create-user",
	Short: "Creates a new user in the user pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_createUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_createUserCmd).Standalone()

		appstream_createUserCmd.Flags().String("authentication-type", "", "The authentication type for the user.")
		appstream_createUserCmd.Flags().String("first-name", "", "The first name, or given name, of the user.")
		appstream_createUserCmd.Flags().String("last-name", "", "The last name, or surname, of the user.")
		appstream_createUserCmd.Flags().String("message-action", "", "The action to take for the welcome email that is sent to a user after the user is created in the user pool.")
		appstream_createUserCmd.Flags().String("user-name", "", "The email address of the user.")
		appstream_createUserCmd.MarkFlagRequired("authentication-type")
		appstream_createUserCmd.MarkFlagRequired("user-name")
	})
	appstreamCmd.AddCommand(appstream_createUserCmd)
}
