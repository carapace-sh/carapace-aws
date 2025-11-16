package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_putTablePolicyCmd = &cobra.Command{
	Use:   "put-table-policy",
	Short: "Creates a new table policy or replaces an existing table policy for a table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_putTablePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3tables_putTablePolicyCmd).Standalone()

		s3tables_putTablePolicyCmd.Flags().String("name", "", "The name of the table.")
		s3tables_putTablePolicyCmd.Flags().String("namespace", "", "The namespace associated with the table.")
		s3tables_putTablePolicyCmd.Flags().String("resource-policy", "", "The `JSON` that defines the policy.")
		s3tables_putTablePolicyCmd.Flags().String("table-bucket-arn", "", "The Amazon Resource Name (ARN) of the table bucket that contains the table.")
		s3tables_putTablePolicyCmd.MarkFlagRequired("name")
		s3tables_putTablePolicyCmd.MarkFlagRequired("namespace")
		s3tables_putTablePolicyCmd.MarkFlagRequired("resource-policy")
		s3tables_putTablePolicyCmd.MarkFlagRequired("table-bucket-arn")
	})
	s3tablesCmd.AddCommand(s3tables_putTablePolicyCmd)
}
