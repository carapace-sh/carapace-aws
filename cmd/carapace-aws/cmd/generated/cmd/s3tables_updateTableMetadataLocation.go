package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_updateTableMetadataLocationCmd = &cobra.Command{
	Use:   "update-table-metadata-location",
	Short: "Updates the metadata location for a table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_updateTableMetadataLocationCmd).Standalone()

	s3tables_updateTableMetadataLocationCmd.Flags().String("metadata-location", "", "The new metadata location for the table.")
	s3tables_updateTableMetadataLocationCmd.Flags().String("name", "", "The name of the table.")
	s3tables_updateTableMetadataLocationCmd.Flags().String("namespace", "", "The namespace of the table.")
	s3tables_updateTableMetadataLocationCmd.Flags().String("table-bucket-arn", "", "The Amazon Resource Name (ARN) of the table bucket.")
	s3tables_updateTableMetadataLocationCmd.Flags().String("version-token", "", "The version token of the table.")
	s3tables_updateTableMetadataLocationCmd.MarkFlagRequired("metadata-location")
	s3tables_updateTableMetadataLocationCmd.MarkFlagRequired("name")
	s3tables_updateTableMetadataLocationCmd.MarkFlagRequired("namespace")
	s3tables_updateTableMetadataLocationCmd.MarkFlagRequired("table-bucket-arn")
	s3tables_updateTableMetadataLocationCmd.MarkFlagRequired("version-token")
	s3tablesCmd.AddCommand(s3tables_updateTableMetadataLocationCmd)
}
