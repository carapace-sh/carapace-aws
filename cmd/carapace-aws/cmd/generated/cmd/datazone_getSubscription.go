package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getSubscriptionCmd = &cobra.Command{
	Use:   "get-subscription",
	Short: "Gets a subscription in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getSubscriptionCmd).Standalone()

	datazone_getSubscriptionCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which the subscription exists.")
	datazone_getSubscriptionCmd.Flags().String("identifier", "", "The ID of the subscription.")
	datazone_getSubscriptionCmd.MarkFlagRequired("domain-identifier")
	datazone_getSubscriptionCmd.MarkFlagRequired("identifier")
	datazoneCmd.AddCommand(datazone_getSubscriptionCmd)
}
