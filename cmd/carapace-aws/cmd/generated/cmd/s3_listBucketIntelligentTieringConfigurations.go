package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_listBucketIntelligentTieringConfigurationsCmd = &cobra.Command{
	Use:   "list-bucket-intelligent-tiering-configurations",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_listBucketIntelligentTieringConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_listBucketIntelligentTieringConfigurationsCmd).Standalone()

		s3_listBucketIntelligentTieringConfigurationsCmd.Flags().String("bucket", "", "The name of the Amazon S3 bucket whose configuration you want to modify or retrieve.")
		s3_listBucketIntelligentTieringConfigurationsCmd.Flags().String("continuation-token", "", "The `ContinuationToken` that represents a placeholder from where this request should begin.")
		s3_listBucketIntelligentTieringConfigurationsCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_listBucketIntelligentTieringConfigurationsCmd.MarkFlagRequired("bucket")
	})
	s3Cmd.AddCommand(s3_listBucketIntelligentTieringConfigurationsCmd)
}
