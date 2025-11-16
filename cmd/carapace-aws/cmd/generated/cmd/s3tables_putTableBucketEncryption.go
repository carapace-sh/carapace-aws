package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_putTableBucketEncryptionCmd = &cobra.Command{
	Use:   "put-table-bucket-encryption",
	Short: "Sets the encryption configuration for a table bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_putTableBucketEncryptionCmd).Standalone()

	s3tables_putTableBucketEncryptionCmd.Flags().String("encryption-configuration", "", "The encryption configuration to apply to the table bucket.")
	s3tables_putTableBucketEncryptionCmd.Flags().String("table-bucket-arn", "", "The Amazon Resource Name (ARN) of the table bucket.")
	s3tables_putTableBucketEncryptionCmd.MarkFlagRequired("encryption-configuration")
	s3tables_putTableBucketEncryptionCmd.MarkFlagRequired("table-bucket-arn")
	s3tablesCmd.AddCommand(s3tables_putTableBucketEncryptionCmd)
}
