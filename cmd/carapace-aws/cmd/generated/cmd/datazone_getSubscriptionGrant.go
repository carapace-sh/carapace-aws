package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getSubscriptionGrantCmd = &cobra.Command{
	Use:   "get-subscription-grant",
	Short: "Gets the subscription grant in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getSubscriptionGrantCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_getSubscriptionGrantCmd).Standalone()

		datazone_getSubscriptionGrantCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which the subscription grant exists.")
		datazone_getSubscriptionGrantCmd.Flags().String("identifier", "", "The ID of the subscription grant.")
		datazone_getSubscriptionGrantCmd.MarkFlagRequired("domain-identifier")
		datazone_getSubscriptionGrantCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_getSubscriptionGrantCmd)
}
