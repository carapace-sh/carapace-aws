package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_listSubscriptionDefinitionVersionsCmd = &cobra.Command{
	Use:   "list-subscription-definition-versions",
	Short: "Lists the versions of a subscription definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_listSubscriptionDefinitionVersionsCmd).Standalone()

	greengrass_listSubscriptionDefinitionVersionsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	greengrass_listSubscriptionDefinitionVersionsCmd.Flags().String("next-token", "", "The token for the next set of results, or ''null'' if there are no additional results.")
	greengrass_listSubscriptionDefinitionVersionsCmd.Flags().String("subscription-definition-id", "", "The ID of the subscription definition.")
	greengrass_listSubscriptionDefinitionVersionsCmd.MarkFlagRequired("subscription-definition-id")
	greengrassCmd.AddCommand(greengrass_listSubscriptionDefinitionVersionsCmd)
}
