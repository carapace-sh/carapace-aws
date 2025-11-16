package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_putTableBucketPolicyCmd = &cobra.Command{
	Use:   "put-table-bucket-policy",
	Short: "Creates a new table bucket policy or replaces an existing table bucket policy for a table bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_putTableBucketPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3tables_putTableBucketPolicyCmd).Standalone()

		s3tables_putTableBucketPolicyCmd.Flags().String("resource-policy", "", "The `JSON` that defines the policy.")
		s3tables_putTableBucketPolicyCmd.Flags().String("table-bucket-arn", "", "The Amazon Resource Name (ARN) of the table bucket.")
		s3tables_putTableBucketPolicyCmd.MarkFlagRequired("resource-policy")
		s3tables_putTableBucketPolicyCmd.MarkFlagRequired("table-bucket-arn")
	})
	s3tablesCmd.AddCommand(s3tables_putTableBucketPolicyCmd)
}
