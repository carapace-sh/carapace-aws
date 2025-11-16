package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_listTableBucketsCmd = &cobra.Command{
	Use:   "list-table-buckets",
	Short: "Lists table buckets for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_listTableBucketsCmd).Standalone()

	s3tables_listTableBucketsCmd.Flags().String("continuation-token", "", "`ContinuationToken` indicates to Amazon S3 that the list is being continued on this bucket with a token.")
	s3tables_listTableBucketsCmd.Flags().String("max-buckets", "", "The maximum number of table buckets to return in the list.")
	s3tables_listTableBucketsCmd.Flags().String("prefix", "", "The prefix of the table buckets.")
	s3tables_listTableBucketsCmd.Flags().String("type", "", "The type of table buckets to filter by in the list.")
	s3tablesCmd.AddCommand(s3tables_listTableBucketsCmd)
}
