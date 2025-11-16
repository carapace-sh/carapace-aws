package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_putJobTaggingCmd = &cobra.Command{
	Use:   "put-job-tagging",
	Short: "Sets the supplied tag-set on an S3 Batch Operations job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_putJobTaggingCmd).Standalone()

	s3control_putJobTaggingCmd.Flags().String("account-id", "", "The Amazon Web Services account ID associated with the S3 Batch Operations job.")
	s3control_putJobTaggingCmd.Flags().String("job-id", "", "The ID for the S3 Batch Operations job whose tags you want to replace.")
	s3control_putJobTaggingCmd.Flags().String("tags", "", "The set of tags to associate with the S3 Batch Operations job.")
	s3control_putJobTaggingCmd.MarkFlagRequired("account-id")
	s3control_putJobTaggingCmd.MarkFlagRequired("job-id")
	s3control_putJobTaggingCmd.MarkFlagRequired("tags")
	s3controlCmd.AddCommand(s3control_putJobTaggingCmd)
}
