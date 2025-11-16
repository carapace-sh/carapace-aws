package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_createConnectorDefinitionCmd = &cobra.Command{
	Use:   "create-connector-definition",
	Short: "Creates a connector definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_createConnectorDefinitionCmd).Standalone()

	greengrass_createConnectorDefinitionCmd.Flags().String("amzn-client-token", "", "A client token used to correlate requests and responses.")
	greengrass_createConnectorDefinitionCmd.Flags().String("initial-version", "", "Information about the initial version of the connector definition.")
	greengrass_createConnectorDefinitionCmd.Flags().String("name", "", "The name of the connector definition.")
	greengrass_createConnectorDefinitionCmd.Flags().String("tags", "", "Tag(s) to add to the new resource.")
	greengrassCmd.AddCommand(greengrass_createConnectorDefinitionCmd)
}
