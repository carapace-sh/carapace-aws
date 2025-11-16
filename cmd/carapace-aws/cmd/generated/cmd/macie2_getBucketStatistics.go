package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_getBucketStatisticsCmd = &cobra.Command{
	Use:   "get-bucket-statistics",
	Short: "Retrieves (queries) aggregated statistical data about all the S3 buckets that Amazon Macie monitors and analyzes for an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_getBucketStatisticsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_getBucketStatisticsCmd).Standalone()

		macie2_getBucketStatisticsCmd.Flags().String("account-id", "", "The unique identifier for the Amazon Web Services account.")
	})
	macie2Cmd.AddCommand(macie2_getBucketStatisticsCmd)
}
