package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_deleteJourneyCmd = &cobra.Command{
	Use:   "delete-journey",
	Short: "Deletes a journey from an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_deleteJourneyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_deleteJourneyCmd).Standalone()

		pinpoint_deleteJourneyCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_deleteJourneyCmd.Flags().String("journey-id", "", "The unique identifier for the journey.")
		pinpoint_deleteJourneyCmd.MarkFlagRequired("application-id")
		pinpoint_deleteJourneyCmd.MarkFlagRequired("journey-id")
	})
	pinpointCmd.AddCommand(pinpoint_deleteJourneyCmd)
}
