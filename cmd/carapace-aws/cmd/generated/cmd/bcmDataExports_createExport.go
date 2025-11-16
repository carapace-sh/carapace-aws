package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmDataExports_createExportCmd = &cobra.Command{
	Use:   "create-export",
	Short: "Creates a data export and specifies the data query, the delivery preference, and any optional resource tags.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmDataExports_createExportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bcmDataExports_createExportCmd).Standalone()

		bcmDataExports_createExportCmd.Flags().String("export", "", "The details of the export, including data query, name, description, and destination configuration.")
		bcmDataExports_createExportCmd.Flags().String("resource-tags", "", "An optional list of tags to associate with the specified export.")
		bcmDataExports_createExportCmd.MarkFlagRequired("export")
	})
	bcmDataExportsCmd.AddCommand(bcmDataExports_createExportCmd)
}
