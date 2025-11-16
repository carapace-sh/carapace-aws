package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_listDomainDeliverabilityCampaignsCmd = &cobra.Command{
	Use:   "list-domain-deliverability-campaigns",
	Short: "Retrieve deliverability data for all the campaigns that used a specific domain to send email during a specified time range.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_listDomainDeliverabilityCampaignsCmd).Standalone()

	sesv2_listDomainDeliverabilityCampaignsCmd.Flags().String("end-date", "", "The last day that you want to obtain deliverability data for.")
	sesv2_listDomainDeliverabilityCampaignsCmd.Flags().String("next-token", "", "A token that’s returned from a previous call to the `ListDomainDeliverabilityCampaigns` operation.")
	sesv2_listDomainDeliverabilityCampaignsCmd.Flags().String("page-size", "", "The maximum number of results to include in response to a single call to the `ListDomainDeliverabilityCampaigns` operation.")
	sesv2_listDomainDeliverabilityCampaignsCmd.Flags().String("start-date", "", "The first day that you want to obtain deliverability data for.")
	sesv2_listDomainDeliverabilityCampaignsCmd.Flags().String("subscribed-domain", "", "The domain to obtain deliverability data for.")
	sesv2_listDomainDeliverabilityCampaignsCmd.MarkFlagRequired("end-date")
	sesv2_listDomainDeliverabilityCampaignsCmd.MarkFlagRequired("start-date")
	sesv2_listDomainDeliverabilityCampaignsCmd.MarkFlagRequired("subscribed-domain")
	sesv2Cmd.AddCommand(sesv2_listDomainDeliverabilityCampaignsCmd)
}
