package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_getJobTaggingCmd = &cobra.Command{
	Use:   "get-job-tagging",
	Short: "Returns the tags on an S3 Batch Operations job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_getJobTaggingCmd).Standalone()

	s3control_getJobTaggingCmd.Flags().String("account-id", "", "The Amazon Web Services account ID associated with the S3 Batch Operations job.")
	s3control_getJobTaggingCmd.Flags().String("job-id", "", "The ID for the S3 Batch Operations job whose tags you want to retrieve.")
	s3control_getJobTaggingCmd.MarkFlagRequired("account-id")
	s3control_getJobTaggingCmd.MarkFlagRequired("job-id")
	s3controlCmd.AddCommand(s3control_getJobTaggingCmd)
}
