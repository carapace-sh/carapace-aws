package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_getCoreDefinitionVersionCmd = &cobra.Command{
	Use:   "get-core-definition-version",
	Short: "Retrieves information about a core definition version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_getCoreDefinitionVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_getCoreDefinitionVersionCmd).Standalone()

		greengrass_getCoreDefinitionVersionCmd.Flags().String("core-definition-id", "", "The ID of the core definition.")
		greengrass_getCoreDefinitionVersionCmd.Flags().String("core-definition-version-id", "", "The ID of the core definition version.")
		greengrass_getCoreDefinitionVersionCmd.MarkFlagRequired("core-definition-id")
		greengrass_getCoreDefinitionVersionCmd.MarkFlagRequired("core-definition-version-id")
	})
	greengrassCmd.AddCommand(greengrass_getCoreDefinitionVersionCmd)
}
