package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_deleteTableBucketPolicyCmd = &cobra.Command{
	Use:   "delete-table-bucket-policy",
	Short: "Deletes a table bucket policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_deleteTableBucketPolicyCmd).Standalone()

	s3tables_deleteTableBucketPolicyCmd.Flags().String("table-bucket-arn", "", "The Amazon Resource Name (ARN) of the table bucket.")
	s3tables_deleteTableBucketPolicyCmd.MarkFlagRequired("table-bucket-arn")
	s3tablesCmd.AddCommand(s3tables_deleteTableBucketPolicyCmd)
}
