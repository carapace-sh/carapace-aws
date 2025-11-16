package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getJourneyCmd = &cobra.Command{
	Use:   "get-journey",
	Short: "Retrieves information about the status, configuration, and other settings for a journey.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getJourneyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_getJourneyCmd).Standalone()

		pinpoint_getJourneyCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_getJourneyCmd.Flags().String("journey-id", "", "The unique identifier for the journey.")
		pinpoint_getJourneyCmd.MarkFlagRequired("application-id")
		pinpoint_getJourneyCmd.MarkFlagRequired("journey-id")
	})
	pinpointCmd.AddCommand(pinpoint_getJourneyCmd)
}
