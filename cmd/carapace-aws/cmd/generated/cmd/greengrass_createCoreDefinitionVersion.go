package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_createCoreDefinitionVersionCmd = &cobra.Command{
	Use:   "create-core-definition-version",
	Short: "Creates a version of a core definition that has already been defined.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_createCoreDefinitionVersionCmd).Standalone()

	greengrass_createCoreDefinitionVersionCmd.Flags().String("amzn-client-token", "", "A client token used to correlate requests and responses.")
	greengrass_createCoreDefinitionVersionCmd.Flags().String("core-definition-id", "", "The ID of the core definition.")
	greengrass_createCoreDefinitionVersionCmd.Flags().String("cores", "", "A list of cores in the core definition version.")
	greengrass_createCoreDefinitionVersionCmd.MarkFlagRequired("core-definition-id")
	greengrassCmd.AddCommand(greengrass_createCoreDefinitionVersionCmd)
}
