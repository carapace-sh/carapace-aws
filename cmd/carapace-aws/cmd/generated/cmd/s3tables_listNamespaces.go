package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_listNamespacesCmd = &cobra.Command{
	Use:   "list-namespaces",
	Short: "Lists the namespaces within a table bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_listNamespacesCmd).Standalone()

	s3tables_listNamespacesCmd.Flags().String("continuation-token", "", "`ContinuationToken` indicates to Amazon S3 that the list is being continued on this bucket with a token.")
	s3tables_listNamespacesCmd.Flags().String("max-namespaces", "", "The maximum number of namespaces to return in the list.")
	s3tables_listNamespacesCmd.Flags().String("prefix", "", "The prefix of the namespaces.")
	s3tables_listNamespacesCmd.Flags().String("table-bucket-arn", "", "The Amazon Resource Name (ARN) of the table bucket.")
	s3tables_listNamespacesCmd.MarkFlagRequired("table-bucket-arn")
	s3tablesCmd.AddCommand(s3tables_listNamespacesCmd)
}
