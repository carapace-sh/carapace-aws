package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_updateResourceDefinitionCmd = &cobra.Command{
	Use:   "update-resource-definition",
	Short: "Updates a resource definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_updateResourceDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_updateResourceDefinitionCmd).Standalone()

		greengrass_updateResourceDefinitionCmd.Flags().String("name", "", "The name of the definition.")
		greengrass_updateResourceDefinitionCmd.Flags().String("resource-definition-id", "", "The ID of the resource definition.")
		greengrass_updateResourceDefinitionCmd.MarkFlagRequired("resource-definition-id")
	})
	greengrassCmd.AddCommand(greengrass_updateResourceDefinitionCmd)
}
