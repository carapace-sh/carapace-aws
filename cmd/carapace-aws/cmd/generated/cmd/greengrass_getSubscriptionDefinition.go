package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_getSubscriptionDefinitionCmd = &cobra.Command{
	Use:   "get-subscription-definition",
	Short: "Retrieves information about a subscription definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_getSubscriptionDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_getSubscriptionDefinitionCmd).Standalone()

		greengrass_getSubscriptionDefinitionCmd.Flags().String("subscription-definition-id", "", "The ID of the subscription definition.")
		greengrass_getSubscriptionDefinitionCmd.MarkFlagRequired("subscription-definition-id")
	})
	greengrassCmd.AddCommand(greengrass_getSubscriptionDefinitionCmd)
}
