package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_updateJobPriorityCmd = &cobra.Command{
	Use:   "update-job-priority",
	Short: "Updates an existing S3 Batch Operations job's priority.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_updateJobPriorityCmd).Standalone()

	s3control_updateJobPriorityCmd.Flags().String("account-id", "", "The Amazon Web Services account ID associated with the S3 Batch Operations job.")
	s3control_updateJobPriorityCmd.Flags().String("job-id", "", "The ID for the job whose priority you want to update.")
	s3control_updateJobPriorityCmd.Flags().String("priority", "", "The priority you want to assign to this job.")
	s3control_updateJobPriorityCmd.MarkFlagRequired("account-id")
	s3control_updateJobPriorityCmd.MarkFlagRequired("job-id")
	s3control_updateJobPriorityCmd.MarkFlagRequired("priority")
	s3controlCmd.AddCommand(s3control_updateJobPriorityCmd)
}
