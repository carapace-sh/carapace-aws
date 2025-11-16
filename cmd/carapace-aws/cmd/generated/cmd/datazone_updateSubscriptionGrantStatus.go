package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_updateSubscriptionGrantStatusCmd = &cobra.Command{
	Use:   "update-subscription-grant-status",
	Short: "Updates the status of the specified subscription grant status in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_updateSubscriptionGrantStatusCmd).Standalone()

	datazone_updateSubscriptionGrantStatusCmd.Flags().String("asset-identifier", "", "The identifier of the asset the subscription grant status of which is to be updated.")
	datazone_updateSubscriptionGrantStatusCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain in which a subscription grant status is to be updated.")
	datazone_updateSubscriptionGrantStatusCmd.Flags().String("failure-cause", "", "Specifies the error message that is returned if the operation cannot be successfully completed.")
	datazone_updateSubscriptionGrantStatusCmd.Flags().String("identifier", "", "The identifier of the subscription grant the status of which is to be updated.")
	datazone_updateSubscriptionGrantStatusCmd.Flags().String("status", "", "The status to be updated as part of the `UpdateSubscriptionGrantStatus` action.")
	datazone_updateSubscriptionGrantStatusCmd.Flags().String("target-name", "", "The target name to be updated as part of the `UpdateSubscriptionGrantStatus` action.")
	datazone_updateSubscriptionGrantStatusCmd.MarkFlagRequired("asset-identifier")
	datazone_updateSubscriptionGrantStatusCmd.MarkFlagRequired("domain-identifier")
	datazone_updateSubscriptionGrantStatusCmd.MarkFlagRequired("identifier")
	datazone_updateSubscriptionGrantStatusCmd.MarkFlagRequired("status")
	datazoneCmd.AddCommand(datazone_updateSubscriptionGrantStatusCmd)
}
