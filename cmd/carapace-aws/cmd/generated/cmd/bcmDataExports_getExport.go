package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmDataExports_getExportCmd = &cobra.Command{
	Use:   "get-export",
	Short: "Views the definition of an existing data export.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmDataExports_getExportCmd).Standalone()

	bcmDataExports_getExportCmd.Flags().String("export-arn", "", "The Amazon Resource Name (ARN) for this export.")
	bcmDataExports_getExportCmd.MarkFlagRequired("export-arn")
	bcmDataExportsCmd.AddCommand(bcmDataExports_getExportCmd)
}
