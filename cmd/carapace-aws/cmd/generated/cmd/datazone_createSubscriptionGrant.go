package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_createSubscriptionGrantCmd = &cobra.Command{
	Use:   "create-subscription-grant",
	Short: "Creates a subsscription grant in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_createSubscriptionGrantCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_createSubscriptionGrantCmd).Standalone()

		datazone_createSubscriptionGrantCmd.Flags().String("asset-target-names", "", "The names of the assets for which the subscription grant is created.")
		datazone_createSubscriptionGrantCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that is provided to ensure the idempotency of the request.")
		datazone_createSubscriptionGrantCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which the subscription grant is created.")
		datazone_createSubscriptionGrantCmd.Flags().String("environment-identifier", "", "The ID of the environment in which the subscription grant is created.")
		datazone_createSubscriptionGrantCmd.Flags().String("granted-entity", "", "The entity to which the subscription is to be granted.")
		datazone_createSubscriptionGrantCmd.Flags().String("subscription-target-identifier", "", "The ID of the subscription target for which the subscription grant is created.")
		datazone_createSubscriptionGrantCmd.MarkFlagRequired("domain-identifier")
		datazone_createSubscriptionGrantCmd.MarkFlagRequired("environment-identifier")
		datazone_createSubscriptionGrantCmd.MarkFlagRequired("granted-entity")
	})
	datazoneCmd.AddCommand(datazone_createSubscriptionGrantCmd)
}
