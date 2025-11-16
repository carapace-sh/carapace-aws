package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_listSubscriptionDefinitionsCmd = &cobra.Command{
	Use:   "list-subscription-definitions",
	Short: "Retrieves a list of subscription definitions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_listSubscriptionDefinitionsCmd).Standalone()

	greengrass_listSubscriptionDefinitionsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	greengrass_listSubscriptionDefinitionsCmd.Flags().String("next-token", "", "The token for the next set of results, or ''null'' if there are no additional results.")
	greengrassCmd.AddCommand(greengrass_listSubscriptionDefinitionsCmd)
}
