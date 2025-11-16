package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_getConnectorDefinitionVersionCmd = &cobra.Command{
	Use:   "get-connector-definition-version",
	Short: "Retrieves information about a connector definition version, including the connectors that the version contains.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_getConnectorDefinitionVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_getConnectorDefinitionVersionCmd).Standalone()

		greengrass_getConnectorDefinitionVersionCmd.Flags().String("connector-definition-id", "", "The ID of the connector definition.")
		greengrass_getConnectorDefinitionVersionCmd.Flags().String("connector-definition-version-id", "", "The ID of the connector definition version.")
		greengrass_getConnectorDefinitionVersionCmd.Flags().String("next-token", "", "The token for the next set of results, or ''null'' if there are no additional results.")
		greengrass_getConnectorDefinitionVersionCmd.MarkFlagRequired("connector-definition-id")
		greengrass_getConnectorDefinitionVersionCmd.MarkFlagRequired("connector-definition-version-id")
	})
	greengrassCmd.AddCommand(greengrass_getConnectorDefinitionVersionCmd)
}
