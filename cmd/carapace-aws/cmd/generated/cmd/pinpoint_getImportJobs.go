package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getImportJobsCmd = &cobra.Command{
	Use:   "get-import-jobs",
	Short: "Retrieves information about the status and settings of all the import jobs for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getImportJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_getImportJobsCmd).Standalone()

		pinpoint_getImportJobsCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_getImportJobsCmd.Flags().String("page-size", "", "The maximum number of items to include in each page of a paginated response.")
		pinpoint_getImportJobsCmd.Flags().String("token", "", "The NextToken string that specifies which page of results to return in a paginated response.")
		pinpoint_getImportJobsCmd.MarkFlagRequired("application-id")
	})
	pinpointCmd.AddCommand(pinpoint_getImportJobsCmd)
}
