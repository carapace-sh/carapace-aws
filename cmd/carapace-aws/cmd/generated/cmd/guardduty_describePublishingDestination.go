package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_describePublishingDestinationCmd = &cobra.Command{
	Use:   "describe-publishing-destination",
	Short: "Returns information about the publishing destination specified by the provided `destinationId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_describePublishingDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(guardduty_describePublishingDestinationCmd).Standalone()

		guardduty_describePublishingDestinationCmd.Flags().String("destination-id", "", "The ID of the publishing destination to retrieve.")
		guardduty_describePublishingDestinationCmd.Flags().String("detector-id", "", "The unique ID of the detector associated with the publishing destination to retrieve.")
		guardduty_describePublishingDestinationCmd.MarkFlagRequired("destination-id")
		guardduty_describePublishingDestinationCmd.MarkFlagRequired("detector-id")
	})
	guarddutyCmd.AddCommand(guardduty_describePublishingDestinationCmd)
}
