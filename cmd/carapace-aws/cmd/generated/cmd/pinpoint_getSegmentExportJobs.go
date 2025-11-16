package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getSegmentExportJobsCmd = &cobra.Command{
	Use:   "get-segment-export-jobs",
	Short: "Retrieves information about the status and settings of the export jobs for a segment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getSegmentExportJobsCmd).Standalone()

	pinpoint_getSegmentExportJobsCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_getSegmentExportJobsCmd.Flags().String("page-size", "", "The maximum number of items to include in each page of a paginated response.")
	pinpoint_getSegmentExportJobsCmd.Flags().String("segment-id", "", "The unique identifier for the segment.")
	pinpoint_getSegmentExportJobsCmd.Flags().String("token", "", "The NextToken string that specifies which page of results to return in a paginated response.")
	pinpoint_getSegmentExportJobsCmd.MarkFlagRequired("application-id")
	pinpoint_getSegmentExportJobsCmd.MarkFlagRequired("segment-id")
	pinpointCmd.AddCommand(pinpoint_getSegmentExportJobsCmd)
}
