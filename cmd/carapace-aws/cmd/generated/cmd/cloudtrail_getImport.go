package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_getImportCmd = &cobra.Command{
	Use:   "get-import",
	Short: "Returns information about a specific import.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_getImportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudtrail_getImportCmd).Standalone()

		cloudtrail_getImportCmd.Flags().String("import-id", "", "The ID for the import.")
		cloudtrail_getImportCmd.MarkFlagRequired("import-id")
	})
	cloudtrailCmd.AddCommand(cloudtrail_getImportCmd)
}
