package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_deleteCoreDefinitionCmd = &cobra.Command{
	Use:   "delete-core-definition",
	Short: "Deletes a core definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_deleteCoreDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_deleteCoreDefinitionCmd).Standalone()

		greengrass_deleteCoreDefinitionCmd.Flags().String("core-definition-id", "", "The ID of the core definition.")
		greengrass_deleteCoreDefinitionCmd.MarkFlagRequired("core-definition-id")
	})
	greengrassCmd.AddCommand(greengrass_deleteCoreDefinitionCmd)
}
