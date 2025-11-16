package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_deleteSubscriptionRequestCmd = &cobra.Command{
	Use:   "delete-subscription-request",
	Short: "Deletes a subscription request in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_deleteSubscriptionRequestCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_deleteSubscriptionRequestCmd).Standalone()

		datazone_deleteSubscriptionRequestCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which the subscription request is deleted.")
		datazone_deleteSubscriptionRequestCmd.Flags().String("identifier", "", "The ID of the subscription request that is deleted.")
		datazone_deleteSubscriptionRequestCmd.MarkFlagRequired("domain-identifier")
		datazone_deleteSubscriptionRequestCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_deleteSubscriptionRequestCmd)
}
