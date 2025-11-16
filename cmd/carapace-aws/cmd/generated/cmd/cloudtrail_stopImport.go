package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_stopImportCmd = &cobra.Command{
	Use:   "stop-import",
	Short: "Stops a specified import.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_stopImportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudtrail_stopImportCmd).Standalone()

		cloudtrail_stopImportCmd.Flags().String("import-id", "", "The ID of the import.")
		cloudtrail_stopImportCmd.MarkFlagRequired("import-id")
	})
	cloudtrailCmd.AddCommand(cloudtrail_stopImportCmd)
}
