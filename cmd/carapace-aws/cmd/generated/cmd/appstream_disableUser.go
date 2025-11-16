package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_disableUserCmd = &cobra.Command{
	Use:   "disable-user",
	Short: "Disables the specified user in the user pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_disableUserCmd).Standalone()

	appstream_disableUserCmd.Flags().String("authentication-type", "", "The authentication type for the user.")
	appstream_disableUserCmd.Flags().String("user-name", "", "The email address of the user.")
	appstream_disableUserCmd.MarkFlagRequired("authentication-type")
	appstream_disableUserCmd.MarkFlagRequired("user-name")
	appstreamCmd.AddCommand(appstream_disableUserCmd)
}
