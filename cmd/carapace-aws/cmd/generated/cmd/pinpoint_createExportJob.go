package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_createExportJobCmd = &cobra.Command{
	Use:   "create-export-job",
	Short: "Creates an export job for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_createExportJobCmd).Standalone()

	pinpoint_createExportJobCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_createExportJobCmd.Flags().String("export-job-request", "", "")
	pinpoint_createExportJobCmd.MarkFlagRequired("application-id")
	pinpoint_createExportJobCmd.MarkFlagRequired("export-job-request")
	pinpointCmd.AddCommand(pinpoint_createExportJobCmd)
}
