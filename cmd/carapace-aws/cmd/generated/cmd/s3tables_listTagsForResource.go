package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all of the tags applied to a specified Amazon S3 Tables resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_listTagsForResourceCmd).Standalone()

	s3tables_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Amazon S3 Tables resource that you want to list tags for.")
	s3tables_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	s3tablesCmd.AddCommand(s3tables_listTagsForResourceCmd)
}
