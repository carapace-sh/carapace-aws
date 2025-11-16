package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_deletePublishingDestinationCmd = &cobra.Command{
	Use:   "delete-publishing-destination",
	Short: "Deletes the publishing definition with the specified `destinationId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_deletePublishingDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(guardduty_deletePublishingDestinationCmd).Standalone()

		guardduty_deletePublishingDestinationCmd.Flags().String("destination-id", "", "The ID of the publishing destination to delete.")
		guardduty_deletePublishingDestinationCmd.Flags().String("detector-id", "", "The unique ID of the detector associated with the publishing destination to delete.")
		guardduty_deletePublishingDestinationCmd.MarkFlagRequired("destination-id")
		guardduty_deletePublishingDestinationCmd.MarkFlagRequired("detector-id")
	})
	guarddutyCmd.AddCommand(guardduty_deletePublishingDestinationCmd)
}
