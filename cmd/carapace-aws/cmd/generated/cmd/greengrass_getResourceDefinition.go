package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_getResourceDefinitionCmd = &cobra.Command{
	Use:   "get-resource-definition",
	Short: "Retrieves information about a resource definition, including its creation time and latest version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_getResourceDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_getResourceDefinitionCmd).Standalone()

		greengrass_getResourceDefinitionCmd.Flags().String("resource-definition-id", "", "The ID of the resource definition.")
		greengrass_getResourceDefinitionCmd.MarkFlagRequired("resource-definition-id")
	})
	greengrassCmd.AddCommand(greengrass_getResourceDefinitionCmd)
}
