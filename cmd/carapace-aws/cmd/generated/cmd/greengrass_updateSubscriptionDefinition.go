package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_updateSubscriptionDefinitionCmd = &cobra.Command{
	Use:   "update-subscription-definition",
	Short: "Updates a subscription definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_updateSubscriptionDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_updateSubscriptionDefinitionCmd).Standalone()

		greengrass_updateSubscriptionDefinitionCmd.Flags().String("name", "", "The name of the definition.")
		greengrass_updateSubscriptionDefinitionCmd.Flags().String("subscription-definition-id", "", "The ID of the subscription definition.")
		greengrass_updateSubscriptionDefinitionCmd.MarkFlagRequired("subscription-definition-id")
	})
	greengrassCmd.AddCommand(greengrass_updateSubscriptionDefinitionCmd)
}
