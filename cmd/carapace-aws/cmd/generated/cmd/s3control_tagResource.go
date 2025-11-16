package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Creates a new user-defined tag or updates an existing tag.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_tagResourceCmd).Standalone()

	s3control_tagResourceCmd.Flags().String("account-id", "", "The Amazon Web Services account ID that created the S3 resource that you're trying to add tags to or the requester's account ID.")
	s3control_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the S3 resource that you're applying tags to.")
	s3control_tagResourceCmd.Flags().String("tags", "", "The Amazon Web Services resource tags that you want to add to the specified S3 resource.")
	s3control_tagResourceCmd.MarkFlagRequired("account-id")
	s3control_tagResourceCmd.MarkFlagRequired("resource-arn")
	s3control_tagResourceCmd.MarkFlagRequired("tags")
	s3controlCmd.AddCommand(s3control_tagResourceCmd)
}
