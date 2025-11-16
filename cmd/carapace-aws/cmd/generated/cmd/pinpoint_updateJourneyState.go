package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_updateJourneyStateCmd = &cobra.Command{
	Use:   "update-journey-state",
	Short: "Cancels (stops) an active journey.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_updateJourneyStateCmd).Standalone()

	pinpoint_updateJourneyStateCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_updateJourneyStateCmd.Flags().String("journey-id", "", "The unique identifier for the journey.")
	pinpoint_updateJourneyStateCmd.Flags().String("journey-state-request", "", "")
	pinpoint_updateJourneyStateCmd.MarkFlagRequired("application-id")
	pinpoint_updateJourneyStateCmd.MarkFlagRequired("journey-id")
	pinpoint_updateJourneyStateCmd.MarkFlagRequired("journey-state-request")
	pinpointCmd.AddCommand(pinpoint_updateJourneyStateCmd)
}
