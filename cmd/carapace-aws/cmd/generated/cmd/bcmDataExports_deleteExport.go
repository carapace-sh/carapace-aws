package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmDataExports_deleteExportCmd = &cobra.Command{
	Use:   "delete-export",
	Short: "Deletes an existing data export.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmDataExports_deleteExportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bcmDataExports_deleteExportCmd).Standalone()

		bcmDataExports_deleteExportCmd.Flags().String("export-arn", "", "The Amazon Resource Name (ARN) for this export.")
		bcmDataExports_deleteExportCmd.MarkFlagRequired("export-arn")
	})
	bcmDataExportsCmd.AddCommand(bcmDataExports_deleteExportCmd)
}
