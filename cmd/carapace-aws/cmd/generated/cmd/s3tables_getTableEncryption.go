package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_getTableEncryptionCmd = &cobra.Command{
	Use:   "get-table-encryption",
	Short: "Gets the encryption configuration for a table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_getTableEncryptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3tables_getTableEncryptionCmd).Standalone()

		s3tables_getTableEncryptionCmd.Flags().String("name", "", "The name of the table.")
		s3tables_getTableEncryptionCmd.Flags().String("namespace", "", "The namespace associated with the table.")
		s3tables_getTableEncryptionCmd.Flags().String("table-bucket-arn", "", "The Amazon Resource Name (ARN) of the table bucket containing the table.")
		s3tables_getTableEncryptionCmd.MarkFlagRequired("name")
		s3tables_getTableEncryptionCmd.MarkFlagRequired("namespace")
		s3tables_getTableEncryptionCmd.MarkFlagRequired("table-bucket-arn")
	})
	s3tablesCmd.AddCommand(s3tables_getTableEncryptionCmd)
}
