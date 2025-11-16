package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_updateSubscriptionTargetCmd = &cobra.Command{
	Use:   "update-subscription-target",
	Short: "Updates the specified subscription target in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_updateSubscriptionTargetCmd).Standalone()

	datazone_updateSubscriptionTargetCmd.Flags().String("applicable-asset-types", "", "The applicable asset types to be updated as part of the `UpdateSubscriptionTarget` action.")
	datazone_updateSubscriptionTargetCmd.Flags().String("authorized-principals", "", "The authorized principals to be updated as part of the `UpdateSubscriptionTarget` action.")
	datazone_updateSubscriptionTargetCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain in which a subscription target is to be updated.")
	datazone_updateSubscriptionTargetCmd.Flags().String("environment-identifier", "", "The identifier of the environment in which a subscription target is to be updated.")
	datazone_updateSubscriptionTargetCmd.Flags().String("identifier", "", "Identifier of the subscription target that is to be updated.")
	datazone_updateSubscriptionTargetCmd.Flags().String("manage-access-role", "", "The manage access role to be updated as part of the `UpdateSubscriptionTarget` action.")
	datazone_updateSubscriptionTargetCmd.Flags().String("name", "", "The name to be updated as part of the `UpdateSubscriptionTarget` action.")
	datazone_updateSubscriptionTargetCmd.Flags().String("provider", "", "The provider to be updated as part of the `UpdateSubscriptionTarget` action.")
	datazone_updateSubscriptionTargetCmd.Flags().String("subscription-target-config", "", "The configuration to be updated as part of the `UpdateSubscriptionTarget` action.")
	datazone_updateSubscriptionTargetCmd.MarkFlagRequired("domain-identifier")
	datazone_updateSubscriptionTargetCmd.MarkFlagRequired("environment-identifier")
	datazone_updateSubscriptionTargetCmd.MarkFlagRequired("identifier")
	datazoneCmd.AddCommand(datazone_updateSubscriptionTargetCmd)
}
