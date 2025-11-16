package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_listJobsCmd = &cobra.Command{
	Use:   "list-jobs",
	Short: "Lists current S3 Batch Operations jobs as well as the jobs that have ended within the last 90 days for the Amazon Web Services account making the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_listJobsCmd).Standalone()

	s3control_listJobsCmd.Flags().String("account-id", "", "The Amazon Web Services account ID associated with the S3 Batch Operations job.")
	s3control_listJobsCmd.Flags().String("job-statuses", "", "The `List Jobs` request returns jobs that match the statuses listed in this element.")
	s3control_listJobsCmd.Flags().String("max-results", "", "The maximum number of jobs that Amazon S3 will include in the `List Jobs` response.")
	s3control_listJobsCmd.Flags().String("next-token", "", "A pagination token to request the next page of results.")
	s3control_listJobsCmd.MarkFlagRequired("account-id")
	s3controlCmd.AddCommand(s3control_listJobsCmd)
}
