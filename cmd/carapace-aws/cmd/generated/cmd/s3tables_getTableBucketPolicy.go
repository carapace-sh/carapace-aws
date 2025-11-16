package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_getTableBucketPolicyCmd = &cobra.Command{
	Use:   "get-table-bucket-policy",
	Short: "Gets details about a table bucket policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_getTableBucketPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3tables_getTableBucketPolicyCmd).Standalone()

		s3tables_getTableBucketPolicyCmd.Flags().String("table-bucket-arn", "", "The Amazon Resource Name (ARN) of the table bucket.")
		s3tables_getTableBucketPolicyCmd.MarkFlagRequired("table-bucket-arn")
	})
	s3tablesCmd.AddCommand(s3tables_getTableBucketPolicyCmd)
}
