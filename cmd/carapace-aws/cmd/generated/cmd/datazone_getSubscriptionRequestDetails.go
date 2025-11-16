package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getSubscriptionRequestDetailsCmd = &cobra.Command{
	Use:   "get-subscription-request-details",
	Short: "Gets the details of the specified subscription request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getSubscriptionRequestDetailsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_getSubscriptionRequestDetailsCmd).Standalone()

		datazone_getSubscriptionRequestDetailsCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain in which to get the subscription request details.")
		datazone_getSubscriptionRequestDetailsCmd.Flags().String("identifier", "", "The identifier of the subscription request the details of which to get.")
		datazone_getSubscriptionRequestDetailsCmd.MarkFlagRequired("domain-identifier")
		datazone_getSubscriptionRequestDetailsCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_getSubscriptionRequestDetailsCmd)
}
