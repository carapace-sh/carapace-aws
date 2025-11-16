package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_enableUserCmd = &cobra.Command{
	Use:   "enable-user",
	Short: "Enables a user in the user pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_enableUserCmd).Standalone()

	appstream_enableUserCmd.Flags().String("authentication-type", "", "The authentication type for the user.")
	appstream_enableUserCmd.Flags().String("user-name", "", "The email address of the user.")
	appstream_enableUserCmd.MarkFlagRequired("authentication-type")
	appstream_enableUserCmd.MarkFlagRequired("user-name")
	appstreamCmd.AddCommand(appstream_enableUserCmd)
}
