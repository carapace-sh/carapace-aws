package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemas_deleteRegistryCmd = &cobra.Command{
	Use:   "delete-registry",
	Short: "Deletes a Registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemas_deleteRegistryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(schemas_deleteRegistryCmd).Standalone()

		schemas_deleteRegistryCmd.Flags().String("registry-name", "", "The name of the registry.")
		schemas_deleteRegistryCmd.MarkFlagRequired("registry-name")
	})
	schemasCmd.AddCommand(schemas_deleteRegistryCmd)
}
