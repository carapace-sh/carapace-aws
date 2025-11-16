package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_getDomainDeliverabilityCampaignCmd = &cobra.Command{
	Use:   "get-domain-deliverability-campaign",
	Short: "Retrieve all the deliverability data for a specific campaign.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_getDomainDeliverabilityCampaignCmd).Standalone()

	sesv2_getDomainDeliverabilityCampaignCmd.Flags().String("campaign-id", "", "The unique identifier for the campaign.")
	sesv2_getDomainDeliverabilityCampaignCmd.MarkFlagRequired("campaign-id")
	sesv2Cmd.AddCommand(sesv2_getDomainDeliverabilityCampaignCmd)
}
