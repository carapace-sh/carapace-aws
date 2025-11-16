package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemas_createRegistryCmd = &cobra.Command{
	Use:   "create-registry",
	Short: "Creates a registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemas_createRegistryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(schemas_createRegistryCmd).Standalone()

		schemas_createRegistryCmd.Flags().String("description", "", "A description of the registry to be created.")
		schemas_createRegistryCmd.Flags().String("registry-name", "", "The name of the registry.")
		schemas_createRegistryCmd.Flags().String("tags", "", "Tags to associate with the registry.")
		schemas_createRegistryCmd.MarkFlagRequired("registry-name")
	})
	schemasCmd.AddCommand(schemas_createRegistryCmd)
}
