package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_deleteJobTaggingCmd = &cobra.Command{
	Use:   "delete-job-tagging",
	Short: "Removes the entire tag set from the specified S3 Batch Operations job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_deleteJobTaggingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_deleteJobTaggingCmd).Standalone()

		s3control_deleteJobTaggingCmd.Flags().String("account-id", "", "The Amazon Web Services account ID associated with the S3 Batch Operations job.")
		s3control_deleteJobTaggingCmd.Flags().String("job-id", "", "The ID for the S3 Batch Operations job whose tags you want to delete.")
		s3control_deleteJobTaggingCmd.MarkFlagRequired("account-id")
		s3control_deleteJobTaggingCmd.MarkFlagRequired("job-id")
	})
	s3controlCmd.AddCommand(s3control_deleteJobTaggingCmd)
}
