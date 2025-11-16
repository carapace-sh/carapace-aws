package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_createResourceDefinitionCmd = &cobra.Command{
	Use:   "create-resource-definition",
	Short: "Creates a resource definition which contains a list of resources to be used in a group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_createResourceDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_createResourceDefinitionCmd).Standalone()

		greengrass_createResourceDefinitionCmd.Flags().String("amzn-client-token", "", "A client token used to correlate requests and responses.")
		greengrass_createResourceDefinitionCmd.Flags().String("initial-version", "", "Information about the initial version of the resource definition.")
		greengrass_createResourceDefinitionCmd.Flags().String("name", "", "The name of the resource definition.")
		greengrass_createResourceDefinitionCmd.Flags().String("tags", "", "Tag(s) to add to the new resource.")
	})
	greengrassCmd.AddCommand(greengrass_createResourceDefinitionCmd)
}
