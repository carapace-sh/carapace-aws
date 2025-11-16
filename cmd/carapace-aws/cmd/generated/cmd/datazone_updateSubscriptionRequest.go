package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_updateSubscriptionRequestCmd = &cobra.Command{
	Use:   "update-subscription-request",
	Short: "Updates a specified subscription request in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_updateSubscriptionRequestCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_updateSubscriptionRequestCmd).Standalone()

		datazone_updateSubscriptionRequestCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain in which a subscription request is to be updated.")
		datazone_updateSubscriptionRequestCmd.Flags().String("identifier", "", "The identifier of the subscription request that is to be updated.")
		datazone_updateSubscriptionRequestCmd.Flags().String("request-reason", "", "The reason for the `UpdateSubscriptionRequest` action.")
		datazone_updateSubscriptionRequestCmd.MarkFlagRequired("domain-identifier")
		datazone_updateSubscriptionRequestCmd.MarkFlagRequired("identifier")
		datazone_updateSubscriptionRequestCmd.MarkFlagRequired("request-reason")
	})
	datazoneCmd.AddCommand(datazone_updateSubscriptionRequestCmd)
}
