package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the specified user-defined tags from an Amazon S3 Tables resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_untagResourceCmd).Standalone()

	s3tables_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Amazon S3 Tables resource that you're removing tags from.")
	s3tables_untagResourceCmd.Flags().String("tag-keys", "", "The array of tag keys that you're removing from the S3 Tables resource.")
	s3tables_untagResourceCmd.MarkFlagRequired("resource-arn")
	s3tables_untagResourceCmd.MarkFlagRequired("tag-keys")
	s3tablesCmd.AddCommand(s3tables_untagResourceCmd)
}
