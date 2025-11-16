package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_createSubscriptionTargetCmd = &cobra.Command{
	Use:   "create-subscription-target",
	Short: "Creates a subscription target in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_createSubscriptionTargetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_createSubscriptionTargetCmd).Standalone()

		datazone_createSubscriptionTargetCmd.Flags().String("applicable-asset-types", "", "The asset types that can be included in the subscription target.")
		datazone_createSubscriptionTargetCmd.Flags().String("authorized-principals", "", "The authorized principals of the subscription target.")
		datazone_createSubscriptionTargetCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that is provided to ensure the idempotency of the request.")
		datazone_createSubscriptionTargetCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which subscription target is created.")
		datazone_createSubscriptionTargetCmd.Flags().String("environment-identifier", "", "The ID of the environment in which subscription target is created.")
		datazone_createSubscriptionTargetCmd.Flags().String("manage-access-role", "", "The manage access role that is used to create the subscription target.")
		datazone_createSubscriptionTargetCmd.Flags().String("name", "", "The name of the subscription target.")
		datazone_createSubscriptionTargetCmd.Flags().String("provider", "", "The provider of the subscription target.")
		datazone_createSubscriptionTargetCmd.Flags().String("subscription-target-config", "", "The configuration of the subscription target.")
		datazone_createSubscriptionTargetCmd.Flags().String("type", "", "The type of the subscription target.")
		datazone_createSubscriptionTargetCmd.MarkFlagRequired("applicable-asset-types")
		datazone_createSubscriptionTargetCmd.MarkFlagRequired("authorized-principals")
		datazone_createSubscriptionTargetCmd.MarkFlagRequired("domain-identifier")
		datazone_createSubscriptionTargetCmd.MarkFlagRequired("environment-identifier")
		datazone_createSubscriptionTargetCmd.MarkFlagRequired("manage-access-role")
		datazone_createSubscriptionTargetCmd.MarkFlagRequired("name")
		datazone_createSubscriptionTargetCmd.MarkFlagRequired("subscription-target-config")
		datazone_createSubscriptionTargetCmd.MarkFlagRequired("type")
	})
	datazoneCmd.AddCommand(datazone_createSubscriptionTargetCmd)
}
