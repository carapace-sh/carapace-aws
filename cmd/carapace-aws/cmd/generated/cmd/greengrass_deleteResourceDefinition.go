package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_deleteResourceDefinitionCmd = &cobra.Command{
	Use:   "delete-resource-definition",
	Short: "Deletes a resource definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_deleteResourceDefinitionCmd).Standalone()

	greengrass_deleteResourceDefinitionCmd.Flags().String("resource-definition-id", "", "The ID of the resource definition.")
	greengrass_deleteResourceDefinitionCmd.MarkFlagRequired("resource-definition-id")
	greengrassCmd.AddCommand(greengrass_deleteResourceDefinitionCmd)
}
