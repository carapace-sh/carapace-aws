package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_cancelSubscriptionCmd = &cobra.Command{
	Use:   "cancel-subscription",
	Short: "Cancels the subscription to the specified asset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_cancelSubscriptionCmd).Standalone()

	datazone_cancelSubscriptionCmd.Flags().String("domain-identifier", "", "The unique identifier of the Amazon DataZone domain where the subscription request is being cancelled.")
	datazone_cancelSubscriptionCmd.Flags().String("identifier", "", "The unique identifier of the subscription that is being cancelled.")
	datazone_cancelSubscriptionCmd.MarkFlagRequired("domain-identifier")
	datazone_cancelSubscriptionCmd.MarkFlagRequired("identifier")
	datazoneCmd.AddCommand(datazone_cancelSubscriptionCmd)
}
