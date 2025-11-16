package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getSubscriptionTargetCmd = &cobra.Command{
	Use:   "get-subscription-target",
	Short: "Gets the subscription target in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getSubscriptionTargetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_getSubscriptionTargetCmd).Standalone()

		datazone_getSubscriptionTargetCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which the subscription target exists.")
		datazone_getSubscriptionTargetCmd.Flags().String("environment-identifier", "", "The ID of the environment associated with the subscription target.")
		datazone_getSubscriptionTargetCmd.Flags().String("identifier", "", "The ID of the subscription target.")
		datazone_getSubscriptionTargetCmd.MarkFlagRequired("domain-identifier")
		datazone_getSubscriptionTargetCmd.MarkFlagRequired("environment-identifier")
		datazone_getSubscriptionTargetCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_getSubscriptionTargetCmd)
}
