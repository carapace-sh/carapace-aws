package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_deleteSubscriptionTargetCmd = &cobra.Command{
	Use:   "delete-subscription-target",
	Short: "Deletes a subscription target in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_deleteSubscriptionTargetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_deleteSubscriptionTargetCmd).Standalone()

		datazone_deleteSubscriptionTargetCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which the subscription target is deleted.")
		datazone_deleteSubscriptionTargetCmd.Flags().String("environment-identifier", "", "The ID of the Amazon DataZone environment in which the subscription target is deleted.")
		datazone_deleteSubscriptionTargetCmd.Flags().String("identifier", "", "The ID of the subscription target that is deleted.")
		datazone_deleteSubscriptionTargetCmd.MarkFlagRequired("domain-identifier")
		datazone_deleteSubscriptionTargetCmd.MarkFlagRequired("environment-identifier")
		datazone_deleteSubscriptionTargetCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_deleteSubscriptionTargetCmd)
}
