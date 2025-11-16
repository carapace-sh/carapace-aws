package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_updateJobStatusCmd = &cobra.Command{
	Use:   "update-job-status",
	Short: "Updates the status for the specified job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_updateJobStatusCmd).Standalone()

	s3control_updateJobStatusCmd.Flags().String("account-id", "", "The Amazon Web Services account ID associated with the S3 Batch Operations job.")
	s3control_updateJobStatusCmd.Flags().String("job-id", "", "The ID of the job whose status you want to update.")
	s3control_updateJobStatusCmd.Flags().String("requested-job-status", "", "The status that you want to move the specified job to.")
	s3control_updateJobStatusCmd.Flags().String("status-update-reason", "", "A description of the reason why you want to change the specified job's status.")
	s3control_updateJobStatusCmd.MarkFlagRequired("account-id")
	s3control_updateJobStatusCmd.MarkFlagRequired("job-id")
	s3control_updateJobStatusCmd.MarkFlagRequired("requested-job-status")
	s3controlCmd.AddCommand(s3control_updateJobStatusCmd)
}
