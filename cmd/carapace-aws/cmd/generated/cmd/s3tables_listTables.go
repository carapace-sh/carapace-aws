package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_listTablesCmd = &cobra.Command{
	Use:   "list-tables",
	Short: "List tables in the given table bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_listTablesCmd).Standalone()

	s3tables_listTablesCmd.Flags().String("continuation-token", "", "`ContinuationToken` indicates to Amazon S3 that the list is being continued on this bucket with a token.")
	s3tables_listTablesCmd.Flags().String("max-tables", "", "The maximum number of tables to return.")
	s3tables_listTablesCmd.Flags().String("namespace", "", "The namespace of the tables.")
	s3tables_listTablesCmd.Flags().String("prefix", "", "The prefix of the tables.")
	s3tables_listTablesCmd.Flags().String("table-bucket-arn", "", "The Amazon resource Name (ARN) of the table bucket.")
	s3tables_listTablesCmd.MarkFlagRequired("table-bucket-arn")
	s3tablesCmd.AddCommand(s3tables_listTablesCmd)
}
