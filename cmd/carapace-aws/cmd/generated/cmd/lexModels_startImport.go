package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_startImportCmd = &cobra.Command{
	Use:   "start-import",
	Short: "Starts a job to import a resource to Amazon Lex.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_startImportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexModels_startImportCmd).Standalone()

		lexModels_startImportCmd.Flags().String("merge-strategy", "", "Specifies the action that the `StartImport` operation should take when there is an existing resource with the same name.")
		lexModels_startImportCmd.Flags().String("payload", "", "A zip archive in binary format.")
		lexModels_startImportCmd.Flags().String("resource-type", "", "Specifies the type of resource to export.")
		lexModels_startImportCmd.Flags().String("tags", "", "A list of tags to add to the imported bot.")
		lexModels_startImportCmd.MarkFlagRequired("merge-strategy")
		lexModels_startImportCmd.MarkFlagRequired("payload")
		lexModels_startImportCmd.MarkFlagRequired("resource-type")
	})
	lexModelsCmd.AddCommand(lexModels_startImportCmd)
}
