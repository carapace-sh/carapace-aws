package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_createConnectorDefinitionVersionCmd = &cobra.Command{
	Use:   "create-connector-definition-version",
	Short: "Creates a version of a connector definition which has already been defined.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_createConnectorDefinitionVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_createConnectorDefinitionVersionCmd).Standalone()

		greengrass_createConnectorDefinitionVersionCmd.Flags().String("amzn-client-token", "", "A client token used to correlate requests and responses.")
		greengrass_createConnectorDefinitionVersionCmd.Flags().String("connector-definition-id", "", "The ID of the connector definition.")
		greengrass_createConnectorDefinitionVersionCmd.Flags().String("connectors", "", "A list of references to connectors in this version, with their corresponding configuration settings.")
		greengrass_createConnectorDefinitionVersionCmd.MarkFlagRequired("connector-definition-id")
	})
	greengrassCmd.AddCommand(greengrass_createConnectorDefinitionVersionCmd)
}
