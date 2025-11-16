package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_deleteNamespaceCmd = &cobra.Command{
	Use:   "delete-namespace",
	Short: "Deletes a namespace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_deleteNamespaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3tables_deleteNamespaceCmd).Standalone()

		s3tables_deleteNamespaceCmd.Flags().String("namespace", "", "The name of the namespace.")
		s3tables_deleteNamespaceCmd.Flags().String("table-bucket-arn", "", "The Amazon Resource Name (ARN) of the table bucket associated with the namespace.")
		s3tables_deleteNamespaceCmd.MarkFlagRequired("namespace")
		s3tables_deleteNamespaceCmd.MarkFlagRequired("table-bucket-arn")
	})
	s3tablesCmd.AddCommand(s3tables_deleteNamespaceCmd)
}
