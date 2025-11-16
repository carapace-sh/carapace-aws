package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifybackend_listS3BucketsCmd = &cobra.Command{
	Use:   "list-s3-buckets",
	Short: "The list of S3 buckets in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifybackend_listS3BucketsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplifybackend_listS3BucketsCmd).Standalone()

		amplifybackend_listS3BucketsCmd.Flags().String("next-token", "", "Reserved for future use.")
	})
	amplifybackendCmd.AddCommand(amplifybackend_listS3BucketsCmd)
}
