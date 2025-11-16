package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_createImportJobCmd = &cobra.Command{
	Use:   "create-import-job",
	Short: "Creates an import job for a data destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_createImportJobCmd).Standalone()

	sesv2_createImportJobCmd.Flags().String("import-data-source", "", "The data source for the import job.")
	sesv2_createImportJobCmd.Flags().String("import-destination", "", "The destination for the import job.")
	sesv2_createImportJobCmd.MarkFlagRequired("import-data-source")
	sesv2_createImportJobCmd.MarkFlagRequired("import-destination")
	sesv2Cmd.AddCommand(sesv2_createImportJobCmd)
}
