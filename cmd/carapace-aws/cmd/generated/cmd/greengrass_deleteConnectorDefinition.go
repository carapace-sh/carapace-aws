package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_deleteConnectorDefinitionCmd = &cobra.Command{
	Use:   "delete-connector-definition",
	Short: "Deletes a connector definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_deleteConnectorDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_deleteConnectorDefinitionCmd).Standalone()

		greengrass_deleteConnectorDefinitionCmd.Flags().String("connector-definition-id", "", "The ID of the connector definition.")
		greengrass_deleteConnectorDefinitionCmd.MarkFlagRequired("connector-definition-id")
	})
	greengrassCmd.AddCommand(greengrass_deleteConnectorDefinitionCmd)
}
