package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Applies one or more user-defined tags to an Amazon S3 Tables resource or updates existing tags.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_tagResourceCmd).Standalone()

	s3tables_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Amazon S3 Tables resource that you're applying tags to.")
	s3tables_tagResourceCmd.Flags().String("tags", "", "The user-defined tag that you want to add to the specified S3 Tables resource.")
	s3tables_tagResourceCmd.MarkFlagRequired("resource-arn")
	s3tables_tagResourceCmd.MarkFlagRequired("tags")
	s3tablesCmd.AddCommand(s3tables_tagResourceCmd)
}
