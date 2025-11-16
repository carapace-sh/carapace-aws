package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_listDomainDeliverabilityCampaignsCmd = &cobra.Command{
	Use:   "list-domain-deliverability-campaigns",
	Short: "Retrieve deliverability data for all the campaigns that used a specific domain to send email during a specified time range.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_listDomainDeliverabilityCampaignsCmd).Standalone()

	pinpointEmail_listDomainDeliverabilityCampaignsCmd.Flags().String("end-date", "", "The last day, in Unix time format, that you want to obtain deliverability data for.")
	pinpointEmail_listDomainDeliverabilityCampaignsCmd.Flags().String("next-token", "", "A token that’s returned from a previous call to the `ListDomainDeliverabilityCampaigns` operation.")
	pinpointEmail_listDomainDeliverabilityCampaignsCmd.Flags().String("page-size", "", "The maximum number of results to include in response to a single call to the `ListDomainDeliverabilityCampaigns` operation.")
	pinpointEmail_listDomainDeliverabilityCampaignsCmd.Flags().String("start-date", "", "The first day, in Unix time format, that you want to obtain deliverability data for.")
	pinpointEmail_listDomainDeliverabilityCampaignsCmd.Flags().String("subscribed-domain", "", "The domain to obtain deliverability data for.")
	pinpointEmail_listDomainDeliverabilityCampaignsCmd.MarkFlagRequired("end-date")
	pinpointEmail_listDomainDeliverabilityCampaignsCmd.MarkFlagRequired("start-date")
	pinpointEmail_listDomainDeliverabilityCampaignsCmd.MarkFlagRequired("subscribed-domain")
	pinpointEmailCmd.AddCommand(pinpointEmail_listDomainDeliverabilityCampaignsCmd)
}
