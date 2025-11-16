package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_getResourceDefinitionVersionCmd = &cobra.Command{
	Use:   "get-resource-definition-version",
	Short: "Retrieves information about a resource definition version, including which resources are included in the version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_getResourceDefinitionVersionCmd).Standalone()

	greengrass_getResourceDefinitionVersionCmd.Flags().String("resource-definition-id", "", "The ID of the resource definition.")
	greengrass_getResourceDefinitionVersionCmd.Flags().String("resource-definition-version-id", "", "The ID of the resource definition version.")
	greengrass_getResourceDefinitionVersionCmd.MarkFlagRequired("resource-definition-id")
	greengrass_getResourceDefinitionVersionCmd.MarkFlagRequired("resource-definition-version-id")
	greengrassCmd.AddCommand(greengrass_getResourceDefinitionVersionCmd)
}
