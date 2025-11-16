package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_associateDrtlogBucketCmd = &cobra.Command{
	Use:   "associate-drtlog-bucket",
	Short: "Authorizes the Shield Response Team (SRT) to access the specified Amazon S3 bucket containing log data such as Application Load Balancer access logs, CloudFront logs, or logs from third party sources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_associateDrtlogBucketCmd).Standalone()

	shield_associateDrtlogBucketCmd.Flags().String("log-bucket", "", "The Amazon S3 bucket that contains the logs that you want to share.")
	shield_associateDrtlogBucketCmd.MarkFlagRequired("log-bucket")
	shieldCmd.AddCommand(shield_associateDrtlogBucketCmd)
}
