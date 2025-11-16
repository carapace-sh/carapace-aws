package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_rejectSubscriptionRequestCmd = &cobra.Command{
	Use:   "reject-subscription-request",
	Short: "Rejects the specified subscription request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_rejectSubscriptionRequestCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_rejectSubscriptionRequestCmd).Standalone()

		datazone_rejectSubscriptionRequestCmd.Flags().String("decision-comment", "", "The decision comment of the rejected subscription request.")
		datazone_rejectSubscriptionRequestCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain in which the subscription request was rejected.")
		datazone_rejectSubscriptionRequestCmd.Flags().String("identifier", "", "The identifier of the subscription request that was rejected.")
		datazone_rejectSubscriptionRequestCmd.MarkFlagRequired("domain-identifier")
		datazone_rejectSubscriptionRequestCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_rejectSubscriptionRequestCmd)
}
