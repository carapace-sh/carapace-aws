package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemas_updateRegistryCmd = &cobra.Command{
	Use:   "update-registry",
	Short: "Updates a registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemas_updateRegistryCmd).Standalone()

	schemas_updateRegistryCmd.Flags().String("description", "", "The description of the registry to update.")
	schemas_updateRegistryCmd.Flags().String("registry-name", "", "The name of the registry.")
	schemas_updateRegistryCmd.MarkFlagRequired("registry-name")
	schemasCmd.AddCommand(schemas_updateRegistryCmd)
}
