package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_getBucketIntelligentTieringConfigurationCmd = &cobra.Command{
	Use:   "get-bucket-intelligent-tiering-configuration",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_getBucketIntelligentTieringConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_getBucketIntelligentTieringConfigurationCmd).Standalone()

		s3_getBucketIntelligentTieringConfigurationCmd.Flags().String("bucket", "", "The name of the Amazon S3 bucket whose configuration you want to modify or retrieve.")
		s3_getBucketIntelligentTieringConfigurationCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_getBucketIntelligentTieringConfigurationCmd.Flags().String("id", "", "The ID used to identify the S3 Intelligent-Tiering configuration.")
		s3_getBucketIntelligentTieringConfigurationCmd.MarkFlagRequired("bucket")
		s3_getBucketIntelligentTieringConfigurationCmd.MarkFlagRequired("id")
	})
	s3Cmd.AddCommand(s3_getBucketIntelligentTieringConfigurationCmd)
}
