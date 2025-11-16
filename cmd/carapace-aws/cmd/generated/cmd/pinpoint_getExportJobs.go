package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getExportJobsCmd = &cobra.Command{
	Use:   "get-export-jobs",
	Short: "Retrieves information about the status and settings of all the export jobs for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getExportJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_getExportJobsCmd).Standalone()

		pinpoint_getExportJobsCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_getExportJobsCmd.Flags().String("page-size", "", "The maximum number of items to include in each page of a paginated response.")
		pinpoint_getExportJobsCmd.Flags().String("token", "", "The NextToken string that specifies which page of results to return in a paginated response.")
		pinpoint_getExportJobsCmd.MarkFlagRequired("application-id")
	})
	pinpointCmd.AddCommand(pinpoint_getExportJobsCmd)
}
