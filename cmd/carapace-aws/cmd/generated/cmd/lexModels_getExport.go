package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_getExportCmd = &cobra.Command{
	Use:   "get-export",
	Short: "Exports the contents of a Amazon Lex resource in a specified format.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_getExportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexModels_getExportCmd).Standalone()

		lexModels_getExportCmd.Flags().String("export-type", "", "The format of the exported data.")
		lexModels_getExportCmd.Flags().String("name", "", "The name of the bot to export.")
		lexModels_getExportCmd.Flags().String("resource-type", "", "The type of resource to export.")
		lexModels_getExportCmd.Flags().String("version", "", "The version of the bot to export.")
		lexModels_getExportCmd.MarkFlagRequired("export-type")
		lexModels_getExportCmd.MarkFlagRequired("name")
		lexModels_getExportCmd.MarkFlagRequired("resource-type")
		lexModels_getExportCmd.MarkFlagRequired("version")
	})
	lexModelsCmd.AddCommand(lexModels_getExportCmd)
}
