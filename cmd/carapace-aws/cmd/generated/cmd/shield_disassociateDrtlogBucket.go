package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_disassociateDrtlogBucketCmd = &cobra.Command{
	Use:   "disassociate-drtlog-bucket",
	Short: "Removes the Shield Response Team's (SRT) access to the specified Amazon S3 bucket containing the logs that you shared previously.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_disassociateDrtlogBucketCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(shield_disassociateDrtlogBucketCmd).Standalone()

		shield_disassociateDrtlogBucketCmd.Flags().String("log-bucket", "", "The Amazon S3 bucket that contains the logs that you want to share.")
		shield_disassociateDrtlogBucketCmd.MarkFlagRequired("log-bucket")
	})
	shieldCmd.AddCommand(shield_disassociateDrtlogBucketCmd)
}
