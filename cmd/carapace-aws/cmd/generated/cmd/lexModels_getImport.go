package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_getImportCmd = &cobra.Command{
	Use:   "get-import",
	Short: "Gets information about an import job started with the `StartImport` operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_getImportCmd).Standalone()

	lexModels_getImportCmd.Flags().String("import-id", "", "The identifier of the import job information to return.")
	lexModels_getImportCmd.MarkFlagRequired("import-id")
	lexModelsCmd.AddCommand(lexModels_getImportCmd)
}
