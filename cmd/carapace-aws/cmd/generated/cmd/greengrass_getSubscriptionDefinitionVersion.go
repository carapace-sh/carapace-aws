package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_getSubscriptionDefinitionVersionCmd = &cobra.Command{
	Use:   "get-subscription-definition-version",
	Short: "Retrieves information about a subscription definition version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_getSubscriptionDefinitionVersionCmd).Standalone()

	greengrass_getSubscriptionDefinitionVersionCmd.Flags().String("next-token", "", "The token for the next set of results, or ''null'' if there are no additional results.")
	greengrass_getSubscriptionDefinitionVersionCmd.Flags().String("subscription-definition-id", "", "The ID of the subscription definition.")
	greengrass_getSubscriptionDefinitionVersionCmd.Flags().String("subscription-definition-version-id", "", "The ID of the subscription definition version.")
	greengrass_getSubscriptionDefinitionVersionCmd.MarkFlagRequired("subscription-definition-id")
	greengrass_getSubscriptionDefinitionVersionCmd.MarkFlagRequired("subscription-definition-version-id")
	greengrassCmd.AddCommand(greengrass_getSubscriptionDefinitionVersionCmd)
}
