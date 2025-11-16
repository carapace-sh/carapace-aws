package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_updateJourneyCmd = &cobra.Command{
	Use:   "update-journey",
	Short: "Updates the configuration and other settings for a journey.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_updateJourneyCmd).Standalone()

	pinpoint_updateJourneyCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_updateJourneyCmd.Flags().String("journey-id", "", "The unique identifier for the journey.")
	pinpoint_updateJourneyCmd.Flags().String("write-journey-request", "", "")
	pinpoint_updateJourneyCmd.MarkFlagRequired("application-id")
	pinpoint_updateJourneyCmd.MarkFlagRequired("journey-id")
	pinpoint_updateJourneyCmd.MarkFlagRequired("write-journey-request")
	pinpointCmd.AddCommand(pinpoint_updateJourneyCmd)
}
