package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_createSubscriptionRequestCmd = &cobra.Command{
	Use:   "create-subscription-request",
	Short: "Creates a subscription request in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_createSubscriptionRequestCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_createSubscriptionRequestCmd).Standalone()

		datazone_createSubscriptionRequestCmd.Flags().String("asset-permissions", "", "The asset permissions of the subscription request.")
		datazone_createSubscriptionRequestCmd.Flags().String("asset-scopes", "", "The asset scopes of the subscription request.")
		datazone_createSubscriptionRequestCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that is provided to ensure the idempotency of the request.")
		datazone_createSubscriptionRequestCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which the subscription request is created.")
		datazone_createSubscriptionRequestCmd.Flags().String("metadata-forms", "", "The metadata form included in the subscription request.")
		datazone_createSubscriptionRequestCmd.Flags().String("request-reason", "", "The reason for the subscription request.")
		datazone_createSubscriptionRequestCmd.Flags().String("subscribed-listings", "", "The published asset for which the subscription grant is to be created.")
		datazone_createSubscriptionRequestCmd.Flags().String("subscribed-principals", "", "The Amazon DataZone principals for whom the subscription request is created.")
		datazone_createSubscriptionRequestCmd.MarkFlagRequired("domain-identifier")
		datazone_createSubscriptionRequestCmd.MarkFlagRequired("request-reason")
		datazone_createSubscriptionRequestCmd.MarkFlagRequired("subscribed-listings")
		datazone_createSubscriptionRequestCmd.MarkFlagRequired("subscribed-principals")
	})
	datazoneCmd.AddCommand(datazone_createSubscriptionRequestCmd)
}
