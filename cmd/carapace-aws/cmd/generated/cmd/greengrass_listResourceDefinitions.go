package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_listResourceDefinitionsCmd = &cobra.Command{
	Use:   "list-resource-definitions",
	Short: "Retrieves a list of resource definitions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_listResourceDefinitionsCmd).Standalone()

	greengrass_listResourceDefinitionsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	greengrass_listResourceDefinitionsCmd.Flags().String("next-token", "", "The token for the next set of results, or ''null'' if there are no additional results.")
	greengrassCmd.AddCommand(greengrass_listResourceDefinitionsCmd)
}
