package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_createPublishingDestinationCmd = &cobra.Command{
	Use:   "create-publishing-destination",
	Short: "Creates a publishing destination where you can export your GuardDuty findings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_createPublishingDestinationCmd).Standalone()

	guardduty_createPublishingDestinationCmd.Flags().String("client-token", "", "The idempotency token for the request.")
	guardduty_createPublishingDestinationCmd.Flags().String("destination-properties", "", "The properties of the publishing destination, including the ARNs for the destination and the KMS key used for encryption.")
	guardduty_createPublishingDestinationCmd.Flags().String("destination-type", "", "The type of resource for the publishing destination.")
	guardduty_createPublishingDestinationCmd.Flags().String("detector-id", "", "The ID of the GuardDuty detector associated with the publishing destination.")
	guardduty_createPublishingDestinationCmd.Flags().String("tags", "", "The tags to be added to a new publishing destination resource.")
	guardduty_createPublishingDestinationCmd.MarkFlagRequired("destination-properties")
	guardduty_createPublishingDestinationCmd.MarkFlagRequired("destination-type")
	guardduty_createPublishingDestinationCmd.MarkFlagRequired("detector-id")
	guarddutyCmd.AddCommand(guardduty_createPublishingDestinationCmd)
}
