package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_createSubscriptionDefinitionVersionCmd = &cobra.Command{
	Use:   "create-subscription-definition-version",
	Short: "Creates a version of a subscription definition which has already been defined.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_createSubscriptionDefinitionVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_createSubscriptionDefinitionVersionCmd).Standalone()

		greengrass_createSubscriptionDefinitionVersionCmd.Flags().String("amzn-client-token", "", "A client token used to correlate requests and responses.")
		greengrass_createSubscriptionDefinitionVersionCmd.Flags().String("subscription-definition-id", "", "The ID of the subscription definition.")
		greengrass_createSubscriptionDefinitionVersionCmd.Flags().String("subscriptions", "", "A list of subscriptions.")
		greengrass_createSubscriptionDefinitionVersionCmd.MarkFlagRequired("subscription-definition-id")
	})
	greengrassCmd.AddCommand(greengrass_createSubscriptionDefinitionVersionCmd)
}
