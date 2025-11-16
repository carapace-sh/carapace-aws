package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_acceptSubscriptionRequestCmd = &cobra.Command{
	Use:   "accept-subscription-request",
	Short: "Accepts a subscription request to a specific asset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_acceptSubscriptionRequestCmd).Standalone()

	datazone_acceptSubscriptionRequestCmd.Flags().String("asset-permissions", "", "The asset permissions of the accept subscription request.")
	datazone_acceptSubscriptionRequestCmd.Flags().String("asset-scopes", "", "The asset scopes of the accept subscription request.")
	datazone_acceptSubscriptionRequestCmd.Flags().String("decision-comment", "", "A description that specifies the reason for accepting the specified subscription request.")
	datazone_acceptSubscriptionRequestCmd.Flags().String("domain-identifier", "", "The Amazon DataZone domain where the specified subscription request is being accepted.")
	datazone_acceptSubscriptionRequestCmd.Flags().String("identifier", "", "The unique identifier of the subscription request that is to be accepted.")
	datazone_acceptSubscriptionRequestCmd.MarkFlagRequired("domain-identifier")
	datazone_acceptSubscriptionRequestCmd.MarkFlagRequired("identifier")
	datazoneCmd.AddCommand(datazone_acceptSubscriptionRequestCmd)
}
