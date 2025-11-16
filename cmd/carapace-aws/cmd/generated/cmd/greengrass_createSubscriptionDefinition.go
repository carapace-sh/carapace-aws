package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_createSubscriptionDefinitionCmd = &cobra.Command{
	Use:   "create-subscription-definition",
	Short: "Creates a subscription definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_createSubscriptionDefinitionCmd).Standalone()

	greengrass_createSubscriptionDefinitionCmd.Flags().String("amzn-client-token", "", "A client token used to correlate requests and responses.")
	greengrass_createSubscriptionDefinitionCmd.Flags().String("initial-version", "", "Information about the initial version of the subscription definition.")
	greengrass_createSubscriptionDefinitionCmd.Flags().String("name", "", "The name of the subscription definition.")
	greengrass_createSubscriptionDefinitionCmd.Flags().String("tags", "", "Tag(s) to add to the new resource.")
	greengrassCmd.AddCommand(greengrass_createSubscriptionDefinitionCmd)
}
