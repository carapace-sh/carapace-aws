package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_getConnectorDefinitionCmd = &cobra.Command{
	Use:   "get-connector-definition",
	Short: "Retrieves information about a connector definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_getConnectorDefinitionCmd).Standalone()

	greengrass_getConnectorDefinitionCmd.Flags().String("connector-definition-id", "", "The ID of the connector definition.")
	greengrass_getConnectorDefinitionCmd.MarkFlagRequired("connector-definition-id")
	greengrassCmd.AddCommand(greengrass_getConnectorDefinitionCmd)
}
