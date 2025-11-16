package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_updateConnectorDefinitionCmd = &cobra.Command{
	Use:   "update-connector-definition",
	Short: "Updates a connector definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_updateConnectorDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_updateConnectorDefinitionCmd).Standalone()

		greengrass_updateConnectorDefinitionCmd.Flags().String("connector-definition-id", "", "The ID of the connector definition.")
		greengrass_updateConnectorDefinitionCmd.Flags().String("name", "", "The name of the definition.")
		greengrass_updateConnectorDefinitionCmd.MarkFlagRequired("connector-definition-id")
	})
	greengrassCmd.AddCommand(greengrass_updateConnectorDefinitionCmd)
}
