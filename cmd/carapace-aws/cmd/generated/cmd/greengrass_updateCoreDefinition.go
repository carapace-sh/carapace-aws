package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_updateCoreDefinitionCmd = &cobra.Command{
	Use:   "update-core-definition",
	Short: "Updates a core definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_updateCoreDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_updateCoreDefinitionCmd).Standalone()

		greengrass_updateCoreDefinitionCmd.Flags().String("core-definition-id", "", "The ID of the core definition.")
		greengrass_updateCoreDefinitionCmd.Flags().String("name", "", "The name of the definition.")
		greengrass_updateCoreDefinitionCmd.MarkFlagRequired("core-definition-id")
	})
	greengrassCmd.AddCommand(greengrass_updateCoreDefinitionCmd)
}
