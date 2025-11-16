package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_getCoreDefinitionCmd = &cobra.Command{
	Use:   "get-core-definition",
	Short: "Retrieves information about a core definition version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_getCoreDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_getCoreDefinitionCmd).Standalone()

		greengrass_getCoreDefinitionCmd.Flags().String("core-definition-id", "", "The ID of the core definition.")
		greengrass_getCoreDefinitionCmd.MarkFlagRequired("core-definition-id")
	})
	greengrassCmd.AddCommand(greengrass_getCoreDefinitionCmd)
}
