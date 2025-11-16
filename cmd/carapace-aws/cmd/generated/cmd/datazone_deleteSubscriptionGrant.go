package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_deleteSubscriptionGrantCmd = &cobra.Command{
	Use:   "delete-subscription-grant",
	Short: "Deletes and subscription grant in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_deleteSubscriptionGrantCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_deleteSubscriptionGrantCmd).Standalone()

		datazone_deleteSubscriptionGrantCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain where the subscription grant is deleted.")
		datazone_deleteSubscriptionGrantCmd.Flags().String("identifier", "", "The ID of the subscription grant that is deleted.")
		datazone_deleteSubscriptionGrantCmd.MarkFlagRequired("domain-identifier")
		datazone_deleteSubscriptionGrantCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_deleteSubscriptionGrantCmd)
}
