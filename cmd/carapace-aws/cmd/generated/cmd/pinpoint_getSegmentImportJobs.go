package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getSegmentImportJobsCmd = &cobra.Command{
	Use:   "get-segment-import-jobs",
	Short: "Retrieves information about the status and settings of the import jobs for a segment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getSegmentImportJobsCmd).Standalone()

	pinpoint_getSegmentImportJobsCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_getSegmentImportJobsCmd.Flags().String("page-size", "", "The maximum number of items to include in each page of a paginated response.")
	pinpoint_getSegmentImportJobsCmd.Flags().String("segment-id", "", "The unique identifier for the segment.")
	pinpoint_getSegmentImportJobsCmd.Flags().String("token", "", "The NextToken string that specifies which page of results to return in a paginated response.")
	pinpoint_getSegmentImportJobsCmd.MarkFlagRequired("application-id")
	pinpoint_getSegmentImportJobsCmd.MarkFlagRequired("segment-id")
	pinpointCmd.AddCommand(pinpoint_getSegmentImportJobsCmd)
}
