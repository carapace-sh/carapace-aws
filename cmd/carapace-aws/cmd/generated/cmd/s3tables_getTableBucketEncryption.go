package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_getTableBucketEncryptionCmd = &cobra.Command{
	Use:   "get-table-bucket-encryption",
	Short: "Gets the encryption configuration for a table bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_getTableBucketEncryptionCmd).Standalone()

	s3tables_getTableBucketEncryptionCmd.Flags().String("table-bucket-arn", "", "The Amazon Resource Name (ARN) of the table bucket.")
	s3tables_getTableBucketEncryptionCmd.MarkFlagRequired("table-bucket-arn")
	s3tablesCmd.AddCommand(s3tables_getTableBucketEncryptionCmd)
}
