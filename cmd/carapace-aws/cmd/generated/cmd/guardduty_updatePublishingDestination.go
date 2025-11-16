package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_updatePublishingDestinationCmd = &cobra.Command{
	Use:   "update-publishing-destination",
	Short: "Updates information about the publishing destination specified by the `destinationId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_updatePublishingDestinationCmd).Standalone()

	guardduty_updatePublishingDestinationCmd.Flags().String("destination-id", "", "The ID of the publishing destination to update.")
	guardduty_updatePublishingDestinationCmd.Flags().String("destination-properties", "", "A `DestinationProperties` object that includes the `DestinationArn` and `KmsKeyArn` of the publishing destination.")
	guardduty_updatePublishingDestinationCmd.Flags().String("detector-id", "", "The ID of the detector associated with the publishing destinations to update.")
	guardduty_updatePublishingDestinationCmd.MarkFlagRequired("destination-id")
	guardduty_updatePublishingDestinationCmd.MarkFlagRequired("detector-id")
	guarddutyCmd.AddCommand(guardduty_updatePublishingDestinationCmd)
}
