package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_createJourneyCmd = &cobra.Command{
	Use:   "create-journey",
	Short: "Creates a journey for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_createJourneyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_createJourneyCmd).Standalone()

		pinpoint_createJourneyCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_createJourneyCmd.Flags().String("write-journey-request", "", "")
		pinpoint_createJourneyCmd.MarkFlagRequired("application-id")
		pinpoint_createJourneyCmd.MarkFlagRequired("write-journey-request")
	})
	pinpointCmd.AddCommand(pinpoint_createJourneyCmd)
}
