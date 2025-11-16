package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_createExportJobCmd = &cobra.Command{
	Use:   "create-export-job",
	Short: "Creates an export job for a data source and destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_createExportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_createExportJobCmd).Standalone()

		sesv2_createExportJobCmd.Flags().String("export-data-source", "", "The data source for the export job.")
		sesv2_createExportJobCmd.Flags().String("export-destination", "", "The destination for the export job.")
		sesv2_createExportJobCmd.MarkFlagRequired("export-data-source")
		sesv2_createExportJobCmd.MarkFlagRequired("export-destination")
	})
	sesv2Cmd.AddCommand(sesv2_createExportJobCmd)
}
