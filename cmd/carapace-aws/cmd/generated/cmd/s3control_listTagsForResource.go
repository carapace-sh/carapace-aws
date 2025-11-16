package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "This operation allows you to list all of the tags for a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_listTagsForResourceCmd).Standalone()

	s3control_listTagsForResourceCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the resource owner.")
	s3control_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the S3 resource that you want to list tags for.")
	s3control_listTagsForResourceCmd.MarkFlagRequired("account-id")
	s3control_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	s3controlCmd.AddCommand(s3control_listTagsForResourceCmd)
}
