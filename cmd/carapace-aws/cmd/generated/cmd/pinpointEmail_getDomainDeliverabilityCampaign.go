package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_getDomainDeliverabilityCampaignCmd = &cobra.Command{
	Use:   "get-domain-deliverability-campaign",
	Short: "Retrieve all the deliverability data for a specific campaign.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_getDomainDeliverabilityCampaignCmd).Standalone()

	pinpointEmail_getDomainDeliverabilityCampaignCmd.Flags().String("campaign-id", "", "The unique identifier for the campaign.")
	pinpointEmail_getDomainDeliverabilityCampaignCmd.MarkFlagRequired("campaign-id")
	pinpointEmailCmd.AddCommand(pinpointEmail_getDomainDeliverabilityCampaignCmd)
}
