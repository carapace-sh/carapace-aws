package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_deleteExportCmd = &cobra.Command{
	Use:   "delete-export",
	Short: "Removes a previous export and the associated files stored in an S3 bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_deleteExportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_deleteExportCmd).Standalone()

		lexv2Models_deleteExportCmd.Flags().String("export-id", "", "The unique identifier of the export to delete.")
		lexv2Models_deleteExportCmd.MarkFlagRequired("export-id")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_deleteExportCmd)
}
