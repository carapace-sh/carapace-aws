package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_createTableCmd = &cobra.Command{
	Use:   "create-table",
	Short: "Creates a new table associated with the given namespace in a table bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_createTableCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3tables_createTableCmd).Standalone()

		s3tables_createTableCmd.Flags().String("encryption-configuration", "", "The encryption configuration to use for the table.")
		s3tables_createTableCmd.Flags().String("format", "", "The format for the table.")
		s3tables_createTableCmd.Flags().String("metadata", "", "The metadata for the table.")
		s3tables_createTableCmd.Flags().String("name", "", "The name for the table.")
		s3tables_createTableCmd.Flags().String("namespace", "", "The namespace to associated with the table.")
		s3tables_createTableCmd.Flags().String("table-bucket-arn", "", "The Amazon Resource Name (ARN) of the table bucket to create the table in.")
		s3tables_createTableCmd.Flags().String("tags", "", "A map of user-defined tags that you would like to apply to the table that you are creating.")
		s3tables_createTableCmd.MarkFlagRequired("format")
		s3tables_createTableCmd.MarkFlagRequired("name")
		s3tables_createTableCmd.MarkFlagRequired("namespace")
		s3tables_createTableCmd.MarkFlagRequired("table-bucket-arn")
	})
	s3tablesCmd.AddCommand(s3tables_createTableCmd)
}
