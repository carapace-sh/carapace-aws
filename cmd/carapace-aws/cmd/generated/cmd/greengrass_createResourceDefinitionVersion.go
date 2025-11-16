package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_createResourceDefinitionVersionCmd = &cobra.Command{
	Use:   "create-resource-definition-version",
	Short: "Creates a version of a resource definition that has already been defined.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_createResourceDefinitionVersionCmd).Standalone()

	greengrass_createResourceDefinitionVersionCmd.Flags().String("amzn-client-token", "", "A client token used to correlate requests and responses.")
	greengrass_createResourceDefinitionVersionCmd.Flags().String("resource-definition-id", "", "The ID of the resource definition.")
	greengrass_createResourceDefinitionVersionCmd.Flags().String("resources", "", "A list of resources.")
	greengrass_createResourceDefinitionVersionCmd.MarkFlagRequired("resource-definition-id")
	greengrassCmd.AddCommand(greengrass_createResourceDefinitionVersionCmd)
}
