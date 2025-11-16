package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_deleteTableBucketEncryptionCmd = &cobra.Command{
	Use:   "delete-table-bucket-encryption",
	Short: "Deletes the encryption configuration for a table bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_deleteTableBucketEncryptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3tables_deleteTableBucketEncryptionCmd).Standalone()

		s3tables_deleteTableBucketEncryptionCmd.Flags().String("table-bucket-arn", "", "The Amazon Resource Name (ARN) of the table bucket.")
		s3tables_deleteTableBucketEncryptionCmd.MarkFlagRequired("table-bucket-arn")
	})
	s3tablesCmd.AddCommand(s3tables_deleteTableBucketEncryptionCmd)
}
