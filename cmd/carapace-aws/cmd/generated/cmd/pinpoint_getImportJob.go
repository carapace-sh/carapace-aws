package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getImportJobCmd = &cobra.Command{
	Use:   "get-import-job",
	Short: "Retrieves information about the status and settings of a specific import job for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getImportJobCmd).Standalone()

	pinpoint_getImportJobCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_getImportJobCmd.Flags().String("job-id", "", "The unique identifier for the job.")
	pinpoint_getImportJobCmd.MarkFlagRequired("application-id")
	pinpoint_getImportJobCmd.MarkFlagRequired("job-id")
	pinpointCmd.AddCommand(pinpoint_getImportJobCmd)
}
