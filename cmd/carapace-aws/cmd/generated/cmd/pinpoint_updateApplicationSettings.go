package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_updateApplicationSettingsCmd = &cobra.Command{
	Use:   "update-application-settings",
	Short: "Updates the settings for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_updateApplicationSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_updateApplicationSettingsCmd).Standalone()

		pinpoint_updateApplicationSettingsCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_updateApplicationSettingsCmd.Flags().String("write-application-settings-request", "", "")
		pinpoint_updateApplicationSettingsCmd.MarkFlagRequired("application-id")
		pinpoint_updateApplicationSettingsCmd.MarkFlagRequired("write-application-settings-request")
	})
	pinpointCmd.AddCommand(pinpoint_updateApplicationSettingsCmd)
}
