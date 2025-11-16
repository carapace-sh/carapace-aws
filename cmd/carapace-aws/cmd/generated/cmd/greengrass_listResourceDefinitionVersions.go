package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_listResourceDefinitionVersionsCmd = &cobra.Command{
	Use:   "list-resource-definition-versions",
	Short: "Lists the versions of a resource definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_listResourceDefinitionVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_listResourceDefinitionVersionsCmd).Standalone()

		greengrass_listResourceDefinitionVersionsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
		greengrass_listResourceDefinitionVersionsCmd.Flags().String("next-token", "", "The token for the next set of results, or ''null'' if there are no additional results.")
		greengrass_listResourceDefinitionVersionsCmd.Flags().String("resource-definition-id", "", "The ID of the resource definition.")
		greengrass_listResourceDefinitionVersionsCmd.MarkFlagRequired("resource-definition-id")
	})
	greengrassCmd.AddCommand(greengrass_listResourceDefinitionVersionsCmd)
}
