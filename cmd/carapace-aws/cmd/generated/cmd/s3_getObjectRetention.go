package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_getObjectRetentionCmd = &cobra.Command{
	Use:   "get-object-retention",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_getObjectRetentionCmd).Standalone()

	s3_getObjectRetentionCmd.Flags().String("bucket", "", "The bucket name containing the object whose retention settings you want to retrieve.")
	s3_getObjectRetentionCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_getObjectRetentionCmd.Flags().String("key", "", "The key name for the object whose retention settings you want to retrieve.")
	s3_getObjectRetentionCmd.Flags().String("request-payer", "", "")
	s3_getObjectRetentionCmd.Flags().String("version-id", "", "The version ID for the object whose retention settings you want to retrieve.")
	s3_getObjectRetentionCmd.MarkFlagRequired("bucket")
	s3_getObjectRetentionCmd.MarkFlagRequired("key")
	s3Cmd.AddCommand(s3_getObjectRetentionCmd)
}
