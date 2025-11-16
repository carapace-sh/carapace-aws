package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_getTableCmd = &cobra.Command{
	Use:   "get-table",
	Short: "Gets details about a table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_getTableCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3tables_getTableCmd).Standalone()

		s3tables_getTableCmd.Flags().String("name", "", "The name of the table.")
		s3tables_getTableCmd.Flags().String("namespace", "", "The name of the namespace the table is associated with.")
		s3tables_getTableCmd.Flags().String("table-arn", "", "The Amazon Resource Name (ARN) of the table.")
		s3tables_getTableCmd.Flags().String("table-bucket-arn", "", "The Amazon Resource Name (ARN) of the table bucket associated with the table.")
	})
	s3tablesCmd.AddCommand(s3tables_getTableCmd)
}
