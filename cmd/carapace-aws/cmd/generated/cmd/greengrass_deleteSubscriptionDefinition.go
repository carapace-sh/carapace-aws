package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_deleteSubscriptionDefinitionCmd = &cobra.Command{
	Use:   "delete-subscription-definition",
	Short: "Deletes a subscription definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_deleteSubscriptionDefinitionCmd).Standalone()

	greengrass_deleteSubscriptionDefinitionCmd.Flags().String("subscription-definition-id", "", "The ID of the subscription definition.")
	greengrass_deleteSubscriptionDefinitionCmd.MarkFlagRequired("subscription-definition-id")
	greengrassCmd.AddCommand(greengrass_deleteSubscriptionDefinitionCmd)
}
