package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_putBucketIntelligentTieringConfigurationCmd = &cobra.Command{
	Use:   "put-bucket-intelligent-tiering-configuration",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_putBucketIntelligentTieringConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_putBucketIntelligentTieringConfigurationCmd).Standalone()

		s3_putBucketIntelligentTieringConfigurationCmd.Flags().String("bucket", "", "The name of the Amazon S3 bucket whose configuration you want to modify or retrieve.")
		s3_putBucketIntelligentTieringConfigurationCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_putBucketIntelligentTieringConfigurationCmd.Flags().String("id", "", "The ID used to identify the S3 Intelligent-Tiering configuration.")
		s3_putBucketIntelligentTieringConfigurationCmd.Flags().String("intelligent-tiering-configuration", "", "Container for S3 Intelligent-Tiering configuration.")
		s3_putBucketIntelligentTieringConfigurationCmd.MarkFlagRequired("bucket")
		s3_putBucketIntelligentTieringConfigurationCmd.MarkFlagRequired("id")
		s3_putBucketIntelligentTieringConfigurationCmd.MarkFlagRequired("intelligent-tiering-configuration")
	})
	s3Cmd.AddCommand(s3_putBucketIntelligentTieringConfigurationCmd)
}
