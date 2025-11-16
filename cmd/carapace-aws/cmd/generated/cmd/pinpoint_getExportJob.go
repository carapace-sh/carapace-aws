package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getExportJobCmd = &cobra.Command{
	Use:   "get-export-job",
	Short: "Retrieves information about the status and settings of a specific export job for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getExportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_getExportJobCmd).Standalone()

		pinpoint_getExportJobCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_getExportJobCmd.Flags().String("job-id", "", "The unique identifier for the job.")
		pinpoint_getExportJobCmd.MarkFlagRequired("application-id")
		pinpoint_getExportJobCmd.MarkFlagRequired("job-id")
	})
	pinpointCmd.AddCommand(pinpoint_getExportJobCmd)
}
