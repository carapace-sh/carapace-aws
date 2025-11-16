package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_deleteTableCmd = &cobra.Command{
	Use:   "delete-table",
	Short: "Deletes a table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_deleteTableCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3tables_deleteTableCmd).Standalone()

		s3tables_deleteTableCmd.Flags().String("name", "", "The name of the table.")
		s3tables_deleteTableCmd.Flags().String("namespace", "", "The namespace associated with the table.")
		s3tables_deleteTableCmd.Flags().String("table-bucket-arn", "", "The Amazon Resource Name (ARN) of the table bucket that contains the table.")
		s3tables_deleteTableCmd.Flags().String("version-token", "", "The version token of the table.")
		s3tables_deleteTableCmd.MarkFlagRequired("name")
		s3tables_deleteTableCmd.MarkFlagRequired("namespace")
		s3tables_deleteTableCmd.MarkFlagRequired("table-bucket-arn")
	})
	s3tablesCmd.AddCommand(s3tables_deleteTableCmd)
}
