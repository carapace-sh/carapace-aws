package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_renameTableCmd = &cobra.Command{
	Use:   "rename-table",
	Short: "Renames a table or a namespace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_renameTableCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3tables_renameTableCmd).Standalone()

		s3tables_renameTableCmd.Flags().String("name", "", "The current name of the table.")
		s3tables_renameTableCmd.Flags().String("namespace", "", "The namespace associated with the table.")
		s3tables_renameTableCmd.Flags().String("new-name", "", "The new name for the table.")
		s3tables_renameTableCmd.Flags().String("new-namespace-name", "", "The new name for the namespace.")
		s3tables_renameTableCmd.Flags().String("table-bucket-arn", "", "The Amazon Resource Name (ARN) of the table bucket.")
		s3tables_renameTableCmd.Flags().String("version-token", "", "The version token of the table.")
		s3tables_renameTableCmd.MarkFlagRequired("name")
		s3tables_renameTableCmd.MarkFlagRequired("namespace")
		s3tables_renameTableCmd.MarkFlagRequired("table-bucket-arn")
	})
	s3tablesCmd.AddCommand(s3tables_renameTableCmd)
}
