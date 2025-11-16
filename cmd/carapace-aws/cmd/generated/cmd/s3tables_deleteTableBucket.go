package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_deleteTableBucketCmd = &cobra.Command{
	Use:   "delete-table-bucket",
	Short: "Deletes a table bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_deleteTableBucketCmd).Standalone()

	s3tables_deleteTableBucketCmd.Flags().String("table-bucket-arn", "", "The Amazon Resource Name (ARN) of the table bucket.")
	s3tables_deleteTableBucketCmd.MarkFlagRequired("table-bucket-arn")
	s3tablesCmd.AddCommand(s3tables_deleteTableBucketCmd)
}
