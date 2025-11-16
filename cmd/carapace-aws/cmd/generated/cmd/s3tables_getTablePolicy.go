package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_getTablePolicyCmd = &cobra.Command{
	Use:   "get-table-policy",
	Short: "Gets details about a table policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_getTablePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3tables_getTablePolicyCmd).Standalone()

		s3tables_getTablePolicyCmd.Flags().String("name", "", "The name of the table.")
		s3tables_getTablePolicyCmd.Flags().String("namespace", "", "The namespace associated with the table.")
		s3tables_getTablePolicyCmd.Flags().String("table-bucket-arn", "", "The Amazon Resource Name (ARN) of the table bucket that contains the table.")
		s3tables_getTablePolicyCmd.MarkFlagRequired("name")
		s3tables_getTablePolicyCmd.MarkFlagRequired("namespace")
		s3tables_getTablePolicyCmd.MarkFlagRequired("table-bucket-arn")
	})
	s3tablesCmd.AddCommand(s3tables_getTablePolicyCmd)
}
