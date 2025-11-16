package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_getTableBucketCmd = &cobra.Command{
	Use:   "get-table-bucket",
	Short: "Gets details on a table bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_getTableBucketCmd).Standalone()

	s3tables_getTableBucketCmd.Flags().String("table-bucket-arn", "", "The Amazon Resource Name (ARN) of the table bucket.")
	s3tables_getTableBucketCmd.MarkFlagRequired("table-bucket-arn")
	s3tablesCmd.AddCommand(s3tables_getTableBucketCmd)
}
