package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "This operation removes the specified user-defined tags from an S3 resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_untagResourceCmd).Standalone()

		s3control_untagResourceCmd.Flags().String("account-id", "", "The Amazon Web Services account ID that owns the resource that you're trying to remove the tags from.")
		s3control_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the S3 resource that you're removing tags from.")
		s3control_untagResourceCmd.Flags().String("tag-keys", "", "The array of tag key-value pairs that you're trying to remove from of the S3 resource.")
		s3control_untagResourceCmd.MarkFlagRequired("account-id")
		s3control_untagResourceCmd.MarkFlagRequired("resource-arn")
		s3control_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	s3controlCmd.AddCommand(s3control_untagResourceCmd)
}
