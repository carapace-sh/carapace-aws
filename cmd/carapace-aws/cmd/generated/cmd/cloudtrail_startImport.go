package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_startImportCmd = &cobra.Command{
	Use:   "start-import",
	Short: "Starts an import of logged trail events from a source S3 bucket to a destination event data store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_startImportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudtrail_startImportCmd).Standalone()

		cloudtrail_startImportCmd.Flags().String("destinations", "", "The ARN of the destination event data store.")
		cloudtrail_startImportCmd.Flags().String("end-event-time", "", "Use with `StartEventTime` to bound a `StartImport` request, and limit imported trail events to only those events logged within a specified time period.")
		cloudtrail_startImportCmd.Flags().String("import-id", "", "The ID of the import.")
		cloudtrail_startImportCmd.Flags().String("import-source", "", "The source S3 bucket for the import.")
		cloudtrail_startImportCmd.Flags().String("start-event-time", "", "Use with `EndEventTime` to bound a `StartImport` request, and limit imported trail events to only those events logged within a specified time period.")
	})
	cloudtrailCmd.AddCommand(cloudtrail_startImportCmd)
}
