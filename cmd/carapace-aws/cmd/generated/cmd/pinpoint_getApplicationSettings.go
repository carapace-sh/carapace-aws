package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getApplicationSettingsCmd = &cobra.Command{
	Use:   "get-application-settings",
	Short: "Retrieves information about the settings for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getApplicationSettingsCmd).Standalone()

	pinpoint_getApplicationSettingsCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_getApplicationSettingsCmd.MarkFlagRequired("application-id")
	pinpointCmd.AddCommand(pinpoint_getApplicationSettingsCmd)
}
