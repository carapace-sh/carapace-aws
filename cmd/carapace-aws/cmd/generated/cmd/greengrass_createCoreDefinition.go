package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_createCoreDefinitionCmd = &cobra.Command{
	Use:   "create-core-definition",
	Short: "Creates a core definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_createCoreDefinitionCmd).Standalone()

	greengrass_createCoreDefinitionCmd.Flags().String("amzn-client-token", "", "A client token used to correlate requests and responses.")
	greengrass_createCoreDefinitionCmd.Flags().String("initial-version", "", "Information about the initial version of the core definition.")
	greengrass_createCoreDefinitionCmd.Flags().String("name", "", "The name of the core definition.")
	greengrass_createCoreDefinitionCmd.Flags().String("tags", "", "Tag(s) to add to the new resource.")
	greengrassCmd.AddCommand(greengrass_createCoreDefinitionCmd)
}
