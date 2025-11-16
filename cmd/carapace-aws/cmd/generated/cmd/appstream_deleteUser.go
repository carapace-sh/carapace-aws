package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_deleteUserCmd = &cobra.Command{
	Use:   "delete-user",
	Short: "Deletes a user from the user pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_deleteUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_deleteUserCmd).Standalone()

		appstream_deleteUserCmd.Flags().String("authentication-type", "", "The authentication type for the user.")
		appstream_deleteUserCmd.Flags().String("user-name", "", "The email address of the user.")
		appstream_deleteUserCmd.MarkFlagRequired("authentication-type")
		appstream_deleteUserCmd.MarkFlagRequired("user-name")
	})
	appstreamCmd.AddCommand(appstream_deleteUserCmd)
}
