package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_getTableMetadataLocationCmd = &cobra.Command{
	Use:   "get-table-metadata-location",
	Short: "Gets the location of the table metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_getTableMetadataLocationCmd).Standalone()

	s3tables_getTableMetadataLocationCmd.Flags().String("name", "", "The name of the table.")
	s3tables_getTableMetadataLocationCmd.Flags().String("namespace", "", "The namespace of the table.")
	s3tables_getTableMetadataLocationCmd.Flags().String("table-bucket-arn", "", "The Amazon Resource Name (ARN) of the table bucket.")
	s3tables_getTableMetadataLocationCmd.MarkFlagRequired("name")
	s3tables_getTableMetadataLocationCmd.MarkFlagRequired("namespace")
	s3tables_getTableMetadataLocationCmd.MarkFlagRequired("table-bucket-arn")
	s3tablesCmd.AddCommand(s3tables_getTableMetadataLocationCmd)
}
