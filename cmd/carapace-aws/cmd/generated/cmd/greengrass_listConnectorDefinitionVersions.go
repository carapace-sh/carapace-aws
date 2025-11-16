package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_listConnectorDefinitionVersionsCmd = &cobra.Command{
	Use:   "list-connector-definition-versions",
	Short: "Lists the versions of a connector definition, which are containers for connectors.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_listConnectorDefinitionVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_listConnectorDefinitionVersionsCmd).Standalone()

		greengrass_listConnectorDefinitionVersionsCmd.Flags().String("connector-definition-id", "", "The ID of the connector definition.")
		greengrass_listConnectorDefinitionVersionsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
		greengrass_listConnectorDefinitionVersionsCmd.Flags().String("next-token", "", "The token for the next set of results, or ''null'' if there are no additional results.")
		greengrass_listConnectorDefinitionVersionsCmd.MarkFlagRequired("connector-definition-id")
	})
	greengrassCmd.AddCommand(greengrass_listConnectorDefinitionVersionsCmd)
}
