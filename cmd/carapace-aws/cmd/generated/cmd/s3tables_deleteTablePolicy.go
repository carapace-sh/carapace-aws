package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_deleteTablePolicyCmd = &cobra.Command{
	Use:   "delete-table-policy",
	Short: "Deletes a table policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_deleteTablePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3tables_deleteTablePolicyCmd).Standalone()

		s3tables_deleteTablePolicyCmd.Flags().String("name", "", "The table name.")
		s3tables_deleteTablePolicyCmd.Flags().String("namespace", "", "The namespace associated with the table.")
		s3tables_deleteTablePolicyCmd.Flags().String("table-bucket-arn", "", "The Amazon Resource Name (ARN) of the table bucket that contains the table.")
		s3tables_deleteTablePolicyCmd.MarkFlagRequired("name")
		s3tables_deleteTablePolicyCmd.MarkFlagRequired("namespace")
		s3tables_deleteTablePolicyCmd.MarkFlagRequired("table-bucket-arn")
	})
	s3tablesCmd.AddCommand(s3tables_deleteTablePolicyCmd)
}
