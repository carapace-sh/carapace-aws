package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmDataExports_updateExportCmd = &cobra.Command{
	Use:   "update-export",
	Short: "Updates an existing data export by overwriting all export parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmDataExports_updateExportCmd).Standalone()

	bcmDataExports_updateExportCmd.Flags().String("export", "", "The name and query details for the export.")
	bcmDataExports_updateExportCmd.Flags().String("export-arn", "", "The Amazon Resource Name (ARN) for this export.")
	bcmDataExports_updateExportCmd.MarkFlagRequired("export")
	bcmDataExports_updateExportCmd.MarkFlagRequired("export-arn")
	bcmDataExportsCmd.AddCommand(bcmDataExports_updateExportCmd)
}
